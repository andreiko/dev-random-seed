# dev-random-seed

Usage example:
```go
package main

import (
	"fmt"
	"github.com/andreiko/dev-random-seed"
)

func main() {
	r, err := dev_random_seed.GetRand()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(r.Intn(1000))
}
```