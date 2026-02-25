package sourcer

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
	"net/url"
	"strings"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

const serpAPIBaseURL = "https://serpapi.com/search"

var controlTerms = []string{
	"ankle sprain",
	"wrist pain",
	"broken bone",
	"blurry vision",
}

type GoogleTrendsConfig struct {
	APIKey  string
	Keyword string
}

type trendsResponse struct {
	InterestOverTime struct {
		Averages []struct {
			Query string `json:"query"`
			Value int    `json:"value"`
		} `json:"averages"`
	} `json:"interest_over_time"`
}

func fetchGoogleTrendsData(config *GoogleTrendsConfig, logger *slog.Logger, sendRequester *http.SendRequester) (float64, error) {
	q := strings.Join(append([]string{config.Keyword}, controlTerms...), ",")

	params := url.Values{}
	params.Set("engine", "google_trends")
	params.Set("date", "now 1-d")
	params.Set("q", q)
	params.Set("api_key", config.APIKey)

	reqURL := fmt.Sprintf("%s?%s", serpAPIBaseURL, params.Encode())

	resp, err := sendRequester.SendRequest(&http.Request{
		Url:    reqURL,
		Method: "GET",
	}).Await()
	if err != nil {
		return 0, fmt.Errorf("SerpAPI request failed: %w", err)
	}
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("SerpAPI error: status %d, body: %s", resp.StatusCode, string(resp.Body))
	}

	var data trendsResponse
	if err := json.Unmarshal(resp.Body, &data); err != nil {
		return 0, fmt.Errorf("failed to parse SerpAPI response: %w", err)
	}

	averages := data.InterestOverTime.Averages

	var keywordAvg float64
	controlAvgs := make([]float64, 0, len(controlTerms))

	for _, a := range averages {
		if a.Query == config.Keyword {
			keywordAvg = float64(a.Value)
		}
		for _, ct := range controlTerms {
			if a.Query == ct {
				controlAvgs = append(controlAvgs, float64(a.Value))
				break
			}
		}
	}

	var baselineSum float64
	for _, v := range controlAvgs {
		baselineSum += v
	}
	baseline := 0.0
	if len(controlAvgs) > 0 {
		baseline = baselineSum / float64(len(controlAvgs))
	}

	standardized := keywordAvg / (baseline + 1)

	score := math.Log(standardized + 1)

	logger.Info("Google Trends score computed",
		"keyword", config.Keyword,
		"keywordAvg", keywordAvg,
		"baseline", baseline,
		"standardized", standardized,
		"score", score,
	)

	return score, nil
}

func GetGoogleTrendsScore(gtConfig *GoogleTrendsConfig, runtime cre.Runtime) cre.Promise[float64] {
	client := &http.Client{}
	return http.SendRequest(gtConfig, runtime, client,
		fetchGoogleTrendsData,
		cre.ConsensusMedianAggregation[float64](),
	)
}
