package main

import (
	"fmt"

	"github.com/smartcontractkit/cre-sdk-go/cre"

	"cre-workflow/wam-workflow/sourcer"
)

type SourceResult struct {
	Score float64
	Ok    bool
}

type AttentionResult struct {
	YouTube      SourceResult
	Twitter      SourceResult
	GoogleTrends SourceResult
	RawIndex     float64
}

// computeRawIndex re-weights available sources proportionally.
// Base weights: G=0.30, Y=0.35, X=0.35 (Y and X each get half of 0.70).
// If a source is unavailable, its weight is redistributed among the rest.
func computeRawIndex(r *AttentionResult) float64 {
	type entry struct {
		weight float64
		score  float64
		ok     bool
	}

	entries := []entry{
		{0.30, r.GoogleTrends.Score, r.GoogleTrends.Ok},
		{0.35, r.YouTube.Score, r.YouTube.Ok},
		{0.35, r.Twitter.Score, r.Twitter.Ok},
	}

	var totalWeight float64
	for _, e := range entries {
		if e.ok {
			totalWeight += e.weight
		}
	}

	if totalWeight == 0 {
		return 0
	}

	var raw float64
	for _, e := range entries {
		if e.ok {
			raw += (e.weight / totalWeight) * e.score
		}
	}

	return 10 * raw
}

func GetRawIndex(keyword string, runtime cre.Runtime, googleAPIKey, twitterAPIKey, serpAPIKey string) (*AttentionResult, error) {
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

	logger := runtime.Logger()

	var yt, tw, gt SourceResult

	ytScore, err := ytPromise.Await()
	if err != nil {
		logger.Warn("YouTube fetch failed, excluding from index", "error", err)
	} else {
		yt = SourceResult{Score: ytScore, Ok: true}
	}

	twScore, err := twPromise.Await()
	if err != nil {
		logger.Warn("Twitter fetch failed, excluding from index", "error", err)
	} else {
		tw = SourceResult{Score: twScore, Ok: true}
	}

	gtScore, err := gtPromise.Await()
	if err != nil {
		logger.Warn("Google Trends fetch failed, excluding from index", "error", err)
	} else {
		gt = SourceResult{Score: gtScore, Ok: true}
	}

	if !yt.Ok && !tw.Ok && !gt.Ok {
		return nil, fmt.Errorf("all data sources failed")
	}

	r := &AttentionResult{
		YouTube:      yt,
		Twitter:      tw,
		GoogleTrends: gt,
	}
	r.RawIndex = computeRawIndex(r)

	return r, nil
}
