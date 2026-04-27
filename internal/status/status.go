// Package status defines the possible states of a pathfinding
// attempt-in-progress and of a completed pathfinding attempt
package status

type Status uint8

//go:generate stringer -type=Status
const (
	Failure Status = iota
	Success
	Continue
)

type RunResult uint8

//go:generate stringer -type=RunResult
const (
	FailedRun RunResult = iota
	SuccessfulRun
)
