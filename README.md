# dev-random-seed

Usage example:
```go
package main

import (
	"math/rand"
	"github.com/andreiko/dev-random-seed"
)

func main() {
	seed, err := dev_random_seed.GetSeed()
	if err != nil {
		panic(err.Error())
	}

	rand.Seed(seed)

}
```