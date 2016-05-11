package backoff

import (
	"log"
	"math/rand"
	"time"
)

type (
	// Policy implements a backoff policy, randomizing its delays.
	Policy struct {
		Intervals []int
		LogPrefix string
	}
)

// Perform loops through each interval, calls the provided anonymous function,
// sleeps if it fails, and returns a boolean of the state.
func (p Policy) Perform(f func() bool) bool {
	for i := 0; i < len(p.Intervals); i++ {
		ok := f()
		if !ok {
			p.sleep(i)
			continue
		}
		return true
	}
	return false
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

	if duration != 0 {
		log.Printf("%s Backing off for %.2fs (Attempt #%d)", p.LogPrefix, duration.Seconds(), n)
	}

	time.Sleep(duration)
}
