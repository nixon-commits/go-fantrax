package models

import "testing"

func TestCalculateWinProbability(t *testing.T) {
	cases := []struct {
		name                                   string
		homeFpts, awayFpts, homeProj, awayProj float64
		homeTimeLeft, awayTimeLeft             int
		wantHome, wantAway                     int
	}{
		{
			name:     "no projections returns coin flip",
			homeProj: 0, awayProj: 0,
			homeFpts: 100, awayFpts: 50,
			homeTimeLeft: 3, awayTimeLeft: 3,
			wantHome: 50, wantAway: 50,
		},
		{
			name:     "all games done — home wins",
			homeFpts: 320, awayFpts: 280,
			homeProj: 320, awayProj: 280,
			homeTimeLeft: 0, awayTimeLeft: 0,
			wantHome: 100, wantAway: 0,
		},
		{
			name:     "all games done — away wins",
			homeFpts: 100, awayFpts: 110,
			homeProj: 100, awayProj: 110,
			homeTimeLeft: 0, awayTimeLeft: 0,
			wantHome: 0, wantAway: 100,
		},
		{
			// Pre-game with equal projections, equal current scores → 50/50
			// would otherwise be nudged off — but pure 0.5 stays at 0.5.
			name:     "equal projections, no points yet — exactly 50/50",
			homeProj: 500, awayProj: 500,
			homeFpts: 0, awayFpts: 0,
			homeTimeLeft: 5, awayTimeLeft: 5,
			wantHome: 50, wantAway: 50,
		},
		{
			// Late game with home projected MUCH higher AND ahead:
			// projection ratio (not score gap) is what drives sharpening,
			// so this case pins WP at 100/0.
			name:     "late game home heavily favored",
			homeProj: 600, awayProj: 400,
			homeFpts: 580, awayFpts: 200,
			homeTimeLeft: 1, awayTimeLeft: 1,
			wantHome: 100, wantAway: 0,
		},
		{
			// Equal projections, big in-game lead: the 0.05 actualWeight
			// barely budges Vh, so this *should not* round to 100. Locks
			// in current behavior so refactors don't accidentally let the
			// in-game score dominate the projection.
			name:     "equal projections, big in-game lead — modest swing",
			homeProj: 500, awayProj: 500,
			homeFpts: 480, awayFpts: 300,
			homeTimeLeft: 1, awayTimeLeft: 1,
			wantHome: 56, wantAway: 44,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			h, a := CalculateWinProbability(c.homeFpts, c.awayFpts, c.homeProj, c.awayProj, c.homeTimeLeft, c.awayTimeLeft)
			if h != c.wantHome || a != c.wantAway {
				t.Errorf("got (%d, %d), want (%d, %d)", h, a, c.wantHome, c.wantAway)
			}
			if h+a != 100 {
				t.Errorf("home+away should sum to 100, got %d", h+a)
			}
		})
	}
}
