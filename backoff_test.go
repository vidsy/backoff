package backoff

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPolicy(t *testing.T) {
	t.Run(".{}", func(t *testing.T) {
		intervals := []int{0, 200, 400}
		logPrefix := "[example]"

		bp := Policy{
			Intervals: intervals,
			LogPrefix: logPrefix,
		}

		assert.NotNil(t, bp.Intervals)
		assert.NotNil(t, bp.LogPrefix)
	})

	t.Run(".Perform", func(t *testing.T) {
		t.Run("IncorrectSleepDuration", func(t *testing.T) {
			intervals := []int{1000, 1000}

			bp := Policy{
				Intervals: intervals,
				LogPrefix: "",
			}

			anon := func() (bool, error) {
				fmt.Println("Connecting...")
				return false, nil
			}

			start := time.Now()
			bp.Perform(anon)
			end := time.Since(start).Seconds()

			if end < 1 || end > 3 {
				t.Error(".Perform() did not sleep for the correct duration.")
			}
		})

		t.Run("ReturnsError", func(t *testing.T) {
			intervals := []int{1000, 1000}

			bp := Policy{
				Intervals: intervals,
				LogPrefix: "",
			}

			anon := func() (bool, error) {
				return false, errors.New("Function error")
			}

			ok, err := bp.Perform(anon)

			assert.Error(t, err)
			assert.False(t, ok)
		})
	})
}
