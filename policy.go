package backoff

import (
	"math/rand"
	"time"
)

// Policy implements a backoff policy, randomizing its delays.
type Policy struct {
	Intervals []int
}

// Duration returns the time duration of the n'th wait cycle in a
// backoff policy. This is b.Intervals[n], randomized to avoid thundering
// herds.
func (p Policy) Duration(n int) time.Duration {
	if n >= len(p.Intervals) {
		n = len(p.Intervals) - 1
	}

	return time.Duration(p.jitter(p.Intervals[n])) * time.Millisecond
}

// jitter returns a random integer uniformly distributed in the range
// [0.5 * millis .. 1.5 * millis]
func (p Policy) jitter(millis int) int {
	if millis == 0 {
		return 0
	}

	return millis/2 + rand.Intn(millis)
}
