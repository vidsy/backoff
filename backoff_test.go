package backoff

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	intervals := []int{0, 200, 400}
	logPrefix := "[example]"

	bp := Policy{
		Intervals: intervals,
		LogPrefix: logPrefix,
	}

	if bp.Intervals == nil {
		t.Error("Intervals not accessible.")
	}

	if bp.LogPrefix == "" {
		t.Error("LogPrefix not accessible.")
	}
}

func TestSleep(t *testing.T) {
	intervals := []int{1000}

	bp := Policy{
		Intervals: intervals,
		LogPrefix: "",
	}

	start := time.Now()
	bp.Sleep(0)
	end := time.Since(start).Seconds()

	if end < 0.5 || end > 1.5 {
		t.Error(".Sleep() did not sleep for the correct duration.")
	}
}
