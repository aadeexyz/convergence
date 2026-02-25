package main

import (
	"fmt"

	"github.com/smartcontractkit/cre-sdk-go/cre"

	"cre-workflow/wam-workflow/sourcer"
)

type Scores struct {
	YouTube      float64
	Twitter      float64
	GoogleTrends float64
	Aggregate    float64
}

func GetAggregate(keyword string, runtime cre.Runtime, googleAPIKey, twitterAPIKey, serpAPIKey string) (*Scores, error) {
	ytPromise := sourcer.GetYouTubeScore(&sourcer.YouTubeConfig{
		APIKey:  googleAPIKey,
		Keyword: keyword,
	}, runtime)

	twPromise := sourcer.GetTwitterScore(&sourcer.TwitterConfig{
		APIKey:  twitterAPIKey,
		Keyword: keyword,
	}, runtime)

	gtPromise := sourcer.GetGoogleTrendsScore(&sourcer.GoogleTrendsConfig{
		APIKey:  serpAPIKey,
		Keyword: keyword,
	}, runtime)

	ytScore, err := ytPromise.Await()
	if err != nil {
		return nil, fmt.Errorf("failed to compute YouTube score: %w", err)
	}

	twScore, err := twPromise.Await()
	if err != nil {
		return nil, fmt.Errorf("failed to compute Twitter score: %w", err)
	}

	gtScore, err := gtPromise.Await()
	if err != nil {
		return nil, fmt.Errorf("failed to compute Google Trends score: %w", err)
	}

	aggregate := (ytScore + twScore + gtScore) / 3.0

	return &Scores{
		YouTube:      ytScore,
		Twitter:      twScore,
		GoogleTrends: gtScore,
		Aggregate:    aggregate,
	}, nil
}
