type Job struct {
	ID     string
	Name   string
	Status Status
	// add timestamps later
}

type Status string

const (
	StatusQueued    Status = "queued"
	StatusRunning   Status = "running"
	StatusSucceeded Status = "succeeded"
	StatusFailed    Status = "failed"
	StatusCanceled  Status = "canceled"
)

import "fmt"

func (j *Job) UpdateStatus(newStatus Status) error {
	cur := j.Status
	if cur == "" {
		cur = StatusQueued
	}

	if cur == newStatus {
		j.Status = newStatus
		return nil
	}

	switch cur {
	case StatusQueued:
		if newStatus == StatusRunning || newStatus == StatusCanceled {
			j.Status = newStatus
			return nil
		}
	case StatusRunning:
		if newStatus == StatusSucceeded || newStatus == StatusFailed || newStatus == StatusCanceled {
			j.Status = newStatus
			return nil
		}
	case StatusFailed:
		// allow retry if you want it
		if newStatus == StatusRunning || newStatus == StatusCanceled {
			j.Status = newStatus
			return nil
		}
	case StatusSucceeded, StatusCanceled:
		// terminal states
		// no transitions allowed
	}

	return fmt.Errorf("invalid status transition: %s -> %s", cur, newStatus)
}