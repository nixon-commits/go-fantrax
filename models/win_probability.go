package models

import "math"

// CalculateWinProbability returns the home/away win probability percentages
// (each in [0, 100], summing to 100) for a head-to-head matchup.
//
// This is a direct Go port of the algorithm Fantrax runs client-side in its
// live-scoring Angular bundle (chunk-5QCCTEMV.js): it does not appear in any
// API response, so calling code must reconstruct it from getLiveScoringStats
// fields. Use this when you want the same numbers the Fantrax UI displays.
//
// Inputs:
//
//	homeFpts, awayFpts      — points scored so far this period
//	                          (use ACTIVE.TotalFpts2.Total, falling back
//	                          to ACTIVE.TotalFpts.Total)
//	homeProj, awayProj      — projected period totals
//	                          (mean of ACTIVE.CalculatedProjectedTotalsMap)
//	homeTimeLeft, awayTimeLeft — games remaining for each team this period
//	                          (ACTIVE.PlayerGameInfo last element, or 0)
//
// Edge cases mirror the JS exactly: zero projections → 50/50, both teams
// out of games → winner takes 100%, and the result is nudged off 50/50
// when the underlying value is close but not exactly tied.
func CalculateWinProbability(homeFpts, awayFpts, homeProj, awayProj float64, homeTimeLeft, awayTimeLeft int) (homePct, awayPct int) {
	if homeProj == 0 && awayProj == 0 {
		return 50, 50
	}
	if homeTimeLeft == 0 && awayTimeLeft == 0 {
		if homeFpts > awayFpts {
			return 100, 0
		}
		return 0, 100
	}

	const (
		pow          = 4.0
		slack        = 0.08
		actualWeight = 0.05
	)

	playedFrac := (homeFpts + awayFpts) / (homeProj + awayProj)
	s := 1 - playedFrac + slack
	if s > 1 {
		s = 1
	}

	dh := homeProj + homeFpts*actualWeight
	da := awayProj + awayFpts*actualWeight
	vh := dh / (dh + da)
	va := 1 - vh

	exp := (1 / s) * pow
	kh := math.Pow(vh, exp)
	ka := math.Pow(va, exp)
	p := kh / (kh + ka)

	if p >= 0.495 && p <= 0.505 && p != 0.5 {
		if p > 0.5 {
			p = 0.51
		} else {
			p = 0.49
		}
	}

	hp := int(math.Round(p * 100))
	return hp, 100 - hp
}
