<h1 align="center">backoff</h1>

<p align="center">
  <b>Go</b> package for implementing a generic backoff policy.
</p>



### Usage

```go
package main

import (
	"errors"
	"log"
	"time"

	"github/vidsy/backoff-policy"
)

func main() {
	bp := backoff.Policy{
		Intervals: []int{0, 500, 1000, 2000, 4000, 8000},
		LogPrefix: "[example]"
	}

	err := connect()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Success!")
}

func connect(bp backoff.Policy) error {
	for i := 0; i < len(bp.Intervals); i++ {
		err := doSomething()
		if err != nil {
			log.Printf("Error: %s", err.Error())
			bp.Sleep(i)
			continue
		}

		return nil
	}

	return errors.New("Unable to connect after backoff")
}
```

### Tests

```
$ go test -v
```

### Notes

[MIT License (MIT)](https://opensource.org/licenses/MIT)
