package backoff

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPolicy(t *testing.T) {
	t.Run(".{}", func(t *testing.T) {
		intervals := []int{0, 200, 400}

		bp := Policy{
			Intervals: intervals,
		}

		assert.NotNil(t, bp.Intervals)
	})

	t.Run(".Perform", func(t *testing.T) {
		t.Run("IncorrectSleepDuration", func(t *testing.T) {
			intervals := []int{1000, 1000}

			bp := Policy{
				Intervals: intervals,
			}

			anon := func() (bool, error) {
				return false, nil
			}

			start := time.Now()
			bp.Perform(anon)
			end := time.Since(start).Seconds()

			if end < 1 || end > 3 {
				t.Error(".Perform() did not sleep for the correct duration.")
			}
		})

		t.Run("CallsLogHandlerMethod", func(t *testing.T) {
			intervals := []int{1000, 1000}
			handlerCalled := false

			bp := Policy{
				Intervals: intervals,
				LogMessageHandler: func(message string) {
					handlerCalled = true
				},
			}

			attemptCount := 0
			anon := func() (bool, error) {
				if attemptCount == 0 {
					attemptCount++
					return false, nil
				}

				return true, nil
			}

			ok, err := bp.Perform(anon)

			assert.Nil(t, err)
			assert.True(t, ok)
			assert.True(t, handlerCalled)
		})

		t.Run("ReturnsError", func(t *testing.T) {
			intervals := []int{1000, 1000}

			bp := Policy{
				Intervals: intervals,
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
