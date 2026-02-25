package sourcer

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
	"net/url"
	"strings"
	"time"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

const twitterBaseURL = "https://api.twitterapi.io/twitter/tweet/advanced_search"

const intervalsPerDay = 96
const maxPages = 5

type TwitterConfig struct {
	APIKey  string
	Keyword string
}

type tweetSearchResponse struct {
	Tweets []struct {
		LikeCount    int `json:"likeCount"`
		ReplyCount   int `json:"replyCount"`
		RetweetCount int `json:"retweetCount"`
	} `json:"tweets"`
	HasNextPage bool   `json:"has_next_page"`
	NextCursor  string `json:"next_cursor"`
}

func fetchTwitterData(config *TwitterConfig, logger *slog.Logger, sendRequester *http.SendRequester) (float64, error) {
	since := time.Now().Add(-24 * time.Hour).UTC()
	sinceStr := since.Format("2006-01-02_15:04:05") + "_UTC"
	query := fmt.Sprintf("%s since:%s", config.Keyword, sinceStr)

	var allLikes, allReplies, allRetweets int64
	var totalPosts int64
	cursor := ""

	for page := 0; page < maxPages; page++ {
		params := url.Values{}
		params.Set("query", query)
		params.Set("queryType", "Top")
		if cursor != "" {
			params.Set("cursor", cursor)
		}

		reqURL := fmt.Sprintf("%s?%s", twitterBaseURL, params.Encode())

		resp, err := sendRequester.SendRequest(&http.Request{
			Url:    reqURL,
			Method: "GET",
			Headers: map[string]string{
				"X-API-Key": config.APIKey,
			},
		}).Await()
		if err != nil {
			return 0, fmt.Errorf("Twitter search request failed (page %d): %w", page, err)
		}
		if resp.StatusCode != 200 {
			return 0, fmt.Errorf("TwitterAPI.io error: status %d, body: %s", resp.StatusCode, string(resp.Body))
		}

		var data tweetSearchResponse
		if err := json.Unmarshal(resp.Body, &data); err != nil {
			return 0, fmt.Errorf("failed to parse Twitter search response: %w", err)
		}

		for _, t := range data.Tweets {
			allLikes += int64(t.LikeCount)
			allReplies += int64(t.ReplyCount)
			allRetweets += int64(t.RetweetCount)
		}
		totalPosts += int64(len(data.Tweets))

		if !data.HasNextPage || strings.TrimSpace(data.NextCursor) == "" {
			break
		}
		cursor = data.NextCursor
	}

	if totalPosts == 0 {
		logger.Info("No tweets found for keyword", "keyword", config.Keyword)
		return 0, nil
	}

	normPosts := float64(totalPosts) / float64(intervalsPerDay)
	normReplies := float64(allReplies) / float64(intervalsPerDay)
	normRetweets := float64(allRetweets) / float64(intervalsPerDay)
	normLikes := float64(allLikes) / float64(intervalsPerDay)

	score :=
		0.50*math.Log(normPosts+1) +
			0.30*math.Log(normReplies+1) +
			0.15*math.Log(normRetweets+1) +
			0.05*math.Log(normLikes+1)

	logger.Info("Twitter score computed",
		"keyword", config.Keyword,
		"totalPosts", totalPosts,
		"totalLikes", allLikes,
		"totalReplies", allReplies,
		"totalRetweets", allRetweets,
		"score", score,
	)

	return score, nil
}

func GetTwitterScore(twConfig *TwitterConfig, runtime cre.Runtime) cre.Promise[float64] {
	client := &http.Client{}
	return http.SendRequest(twConfig, runtime, client,
		fetchTwitterData,
		cre.ConsensusMedianAggregation[float64](),
	)
}
