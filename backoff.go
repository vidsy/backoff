package backoff

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	// Policy implements a backoff policy, randomizing its delays.
	Policy struct {
		Intervals         []int
		LogMessageHandler func(message string)
	}
)

// Perform loops through each interval, calls the provided anonymous function,
// sleeps if it fails, and returns a boolean of the state.
func (p Policy) Perform(f func() (bool, error)) (bool, error) {
	var errors Errors

	for i := 0; i < len(p.Intervals); i++ {
		ok, err := f()
		if err != nil {
			errors = append(errors, err)
		}

		if !ok {
			p.sleep(i)
			continue
		}

		return true, nil
	}

	if len(errors) > 0 {
		return false, errors
	}

	return false, nil
}

func (p Policy) jitter(millis int) int {
	if millis == 0 {
		return 0
	}

	// 50 to 150%
	// x = 100; (100/2 + rand(0..100); e.g: 50 + 72 = 122
	return millis/2 + rand.Intn(millis)
}

func (p Policy) sleep(n int) {
	if n >= len(p.Intervals) {
		n = len(p.Intervals) - 1
	}

	duration := time.Duration(p.jitter(p.Intervals[n])) * time.Millisecond

	if duration != 0 && p.LogMessageHandler != nil {
		p.LogMessageHandler(
			fmt.Sprintf(
				"Backing off for %.2fs (attempt %d/%d)", duration.Seconds(), n+1, len(p.Intervals),
			),
		)
	}

	time.Sleep(duration)
}
