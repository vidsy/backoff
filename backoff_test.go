package backoff_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vidsy/backoff"
)

func TestPolicy(t *testing.T) {
	t.Run("Perform", func(t *testing.T) {
		t.Run("IncorrectSleepDuration", func(t *testing.T) {
			intervals := []int{1000, 1000}

			bp := backoff.Policy{
				Intervals: intervals,
			}

			anon := func() (bool, error) {
				return false, nil
			}

			start := time.Now()
			ok, err := bp.Perform(anon)
			require.NoError(t, err)
			require.False(t, ok)

			end := time.Since(start).Seconds()

			if end < 1 || end > 3 {
				t.Error(".Perform() did not sleep for the correct duration.")
			}
		})

		t.Run("CallsLogHandlerMethod", func(t *testing.T) {
			intervals := []int{1000, 1000}
			handlerCalled := false

			bp := backoff.Policy{
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

			require.Nil(t, err)
			require.True(t, ok)
			require.True(t, handlerCalled)
		})

		t.Run("ReturnsErrors", func(t *testing.T) {
			intervals := []int{1000, 1000}

			bp := backoff.Policy{
				Intervals: intervals,
			}

			anon := func() (bool, error) {
				return false, errors.New("Function error")
			}

			ok, err := bp.Perform(anon)
			require.Error(t, err)
			require.False(t, ok)
			require.IsType(t, backoff.Errors{}, err)
			require.Len(t, err.(backoff.Errors), 2)
		})
	})
}
