<h1 align="center">backoff-policy</h1>

<p align="center">
  <b>Go</b> package for implementing a generic backoff policy.
</p>



### Usage

```go
package main

import(
	"log"
	"time"

	"github/vidsy/backoff-policy"
)

func main() {
	bp := backoff.Policy{
		[]int{0, 500, 1000, 2000, 4000, 8000},
	},

	for i := 0; i < len(bp.Intervals); i++ {
		duration := bp.Duration(i)
		time.Sleep(duration)

		if i != 0 {
			log.Printf("Backing off for %dms... (Attempt #%d)", duration. i)
		}

		err := doSomething()
		if err != nil {
			if i == len(bp.Intervals) {
				log.Fatal(err)
			}
			continue
		} else {
			break
		}
	}

	log.Println("Success!")
}
```

### Notes

[MIT License (MIT)](https://opensource.org/licenses/MIT)
