package backoff

import (
	"fmt"
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

func TestPerform(t *testing.T) {
	intervals := []int{1000, 1000}

	bp := Policy{
		Intervals: intervals,
		LogPrefix: "",
	}

	anon := func() bool {
		fmt.Println("Connecting...")
		return false
	}

	start := time.Now()
	_ = bp.Perform(anon)
	end := time.Since(start).Seconds()

	if end < 1 || end > 3 {
		t.Error(".Perform() did not sleep for the correct duration.")
	}
}
