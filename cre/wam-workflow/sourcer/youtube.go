package sourcer

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

const youtubeBaseURL = "https://www.googleapis.com/youtube/v3"

type YouTubeConfig struct {
	APIKey  string
	Keyword string
}

type searchResponse struct {
	PageInfo struct {
		TotalResults int `json:"totalResults"`
	} `json:"pageInfo"`
	Items []struct {
		Id struct {
			VideoId string `json:"videoId"`
		} `json:"id"`
	} `json:"items"`
}

type videoStatsResponse struct {
	Items []struct {
		Id         string `json:"id"`
		Statistics struct {
			ViewCount    string `json:"viewCount"`
			LikeCount    string `json:"likeCount"`
			CommentCount string `json:"commentCount"`
		} `json:"statistics"`
	} `json:"items"`
}

func fetchYouTubeData(config *YouTubeConfig, logger *slog.Logger, sendRequester *http.SendRequester) (float64, error) {
	publishedAfter := time.Now().Add(-24 * time.Hour).UTC().Format(time.RFC3339)

	searchParams := url.Values{}
	searchParams.Set("part", "snippet")
	searchParams.Set("type", "video")
	searchParams.Set("order", "relevance")
	searchParams.Set("maxResults", "50")
	searchParams.Set("publishedAfter", publishedAfter)
	searchParams.Set("q", config.Keyword)
	searchParams.Set("key", config.APIKey)

	searchURL := fmt.Sprintf("%s/search?%s", youtubeBaseURL, searchParams.Encode())

	searchResp, err := sendRequester.SendRequest(&http.Request{
		Url:    searchURL,
		Method: "GET",
	}).Await()
	if err != nil {
		return 0, fmt.Errorf("YouTube search request failed: %w", err)
	}
	if searchResp.StatusCode != 200 {
		return 0, fmt.Errorf("YouTube search error: status %d, body: %s", searchResp.StatusCode, string(searchResp.Body))
	}

	var searchData searchResponse
	if err := json.Unmarshal(searchResp.Body, &searchData); err != nil {
		return 0, fmt.Errorf("failed to parse YouTube search response: %w", err)
	}

	totalResults := searchData.PageInfo.TotalResults

	videoIds := make([]string, 0, len(searchData.Items))
	for _, item := range searchData.Items {
		if item.Id.VideoId != "" {
			videoIds = append(videoIds, item.Id.VideoId)
		}
	}

	if len(videoIds) == 0 {
		logger.Info("No videos found for keyword", "keyword", config.Keyword)
		return 0, nil
	}

	statsParams := url.Values{}
	statsParams.Set("part", "statistics")
	statsParams.Set("id", strings.Join(videoIds, ","))
	statsParams.Set("key", config.APIKey)

	statsURL := fmt.Sprintf("%s/videos?%s", youtubeBaseURL, statsParams.Encode())

	statsResp, err := sendRequester.SendRequest(&http.Request{
		Url:    statsURL,
		Method: "GET",
	}).Await()
	if err != nil {
		return 0, fmt.Errorf("YouTube videos stats request failed: %w", err)
	}
	if statsResp.StatusCode != 200 {
		return 0, fmt.Errorf("YouTube videos error: status %d, body: %s", statsResp.StatusCode, string(statsResp.Body))
	}

	var statsData videoStatsResponse
	if err := json.Unmarshal(statsResp.Body, &statsData); err != nil {
		return 0, fmt.Errorf("failed to parse YouTube videos response: %w", err)
	}

	var totalViews, totalLikes, totalComments int64
	for _, item := range statsData.Items {
		views, _ := strconv.ParseInt(item.Statistics.ViewCount, 10, 64)
		likes, _ := strconv.ParseInt(item.Statistics.LikeCount, 10, 64)
		comments, _ := strconv.ParseInt(item.Statistics.CommentCount, 10, 64)
		totalViews += views
		totalLikes += likes
		totalComments += comments
	}

	score :=
		0.50*math.Log(float64(totalResults)+1) +
			0.30*math.Log(float64(totalComments)+1) +
			0.15*math.Log(float64(totalLikes)+1) +
			0.05*math.Log(float64(totalViews)+1)

	logger.Info("YouTube score computed",
		"keyword", config.Keyword,
		"totalResults", totalResults,
		"videos", len(videoIds),
		"totalViews", totalViews,
		"totalLikes", totalLikes,
		"totalComments", totalComments,
		"score", score,
	)

	return score, nil
}

func GetYouTubeScore(ytConfig *YouTubeConfig, runtime cre.Runtime) cre.Promise[float64] {
	client := &http.Client{}
	return http.SendRequest(ytConfig, runtime, client,
		fetchYouTubeData,
		cre.ConsensusMedianAggregation[float64](),
	)
}
