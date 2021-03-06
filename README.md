<h1 align="center">backoff</h1>

<p align="center">
  <img src="https://circleci.com/gh/vidsy/backoff/tree/master.svg?style=shield">
</p>


<p align="center">
  <b>Go</b> package for implementing a generic backoff policy.
</p>

### Usage

```go
package main

import (
  "log"

  "github/vidsy/backoff"
)

func main() {
  bp := backoff.Policy{
    Intervals: []int{0, 500, 1000, 2000, 4000, 8000},
    LogMessageHandler: func(message string) {
      log.Printf("Log message from backoff.Policy: %s", message)
    }
  }

  anon := func() (bool, error) {
    attemptConnection()
  }

  ok, err := bp.Perform(anon)
  if err != nil {
    log.Fatalf("Error in anon function: %s", err)
  }

  if !ok {
    log.Fatal("Failed to connect...")
  }

  log.Println("Success!")
}

func attemptConnection() bool {
  // Do something here...
  return false
}
```

### Tests

```
$ go test -v
```

### Notes

[MIT License (MIT)](https://opensource.org/licenses/MIT)

