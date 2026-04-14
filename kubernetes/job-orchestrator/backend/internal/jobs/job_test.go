package jobs

import (
	"testing"
)

func TestJob(t *testing.T) {
	job := &Job{ID: "1", Name: "Test Job"}

	// Test initial status
	if job.Status != "" {
		t.Errorf("expected initial status to be empty, got %s", job.Status)
	}

	// Test valid transitions
	if err := job.UpdateStatus(StatusRunning); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if job.Status != StatusRunning {
		t.Errorf("expected status to be %s, got %s", StatusRunning, job.Status)
	}

	if err := job.UpdateStatus(StatusSucceeded); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if job.Status != StatusSucceeded {
		t.Errorf("expected status to be %s, got %s", StatusSucceeded, job.Status)
	}

	// Test invalid transition
	if err := job.UpdateStatus(StatusRunning); err == nil {
		t.Errorf("expected error for invalid transition, got nil")
	}

	// Test transition from failed to running (retry)
	job.Status = StatusFailed
	if err := job.UpdateStatus(StatusRunning); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if job.Status != StatusRunning {
		t.Errorf("expected status to be %s, got %s", StatusRunning, job.Status)
	}
}
