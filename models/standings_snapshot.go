package models

// StandingsSnapshotStatus is the response of the getStandingsSnapshotStatus
// fxpa method. A snapshot becomes "ready" once Fantrax has closed out the
// most recent scoring period (week) and recomputed standings — useful as a
// trigger for weekly recap jobs.
type StandingsSnapshotStatus struct {
	Ready bool `json:"ready"`
}

// StandingsSnapshotStatusResponse is the full fxpa envelope.
type StandingsSnapshotStatusResponse struct {
	Responses []struct {
		Data StandingsSnapshotStatus `json:"data"`
	} `json:"responses"`
}
