package backoff

import (
	"log"
	"math/rand"
	"time"
)

// Policy implements a backoff policy, randomizing its delays.
type Policy struct {
	Intervals []int
	LogPrefix string
}

// Sleep calculates a time duration of the n'th wait cycle in a
// backoff policy. It then sleeps for that duration.
func (p Policy) Sleep(n int) {
	if n >= len(p.Intervals) {
		n = len(p.Intervals) - 1
	}

	duration := time.Duration(p.jitter(p.Intervals[n])) * time.Millisecond

	if duration != 0 {
		log.Printf("%s Backing off for %dms (Attempt #%d)", p.LogPrefix, duration, n)
	}

	time.Sleep(duration)
}

// jitter returns a random integer uniformly distributed in the range
// [0.5 * millis .. 1.5 * millis]
func (p Policy) jitter(millis int) int {
	if millis == 0 {
		return 0
	}

	return millis/2 + rand.Intn(millis)
}
