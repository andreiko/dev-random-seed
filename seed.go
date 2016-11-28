package dev_random_seed

import (
	"fmt"
	"os"
)

const Filename = "/dev/random"

func GetSeed() (int64, error) {
	f, err := os.OpenFile(Filename, os.O_RDONLY, 0)
	if err != nil {
		return 0, err
	}

	var bytes [8]byte
	n, err := f.Read(bytes[:])

	if err != nil {
		return 0, err
	}

	if n < len(bytes) {
		return 0, fmt.Errorf("%d byte(s) instead of %d were read from %s", n, len(bytes), Filename)
	}

	var result int64
	for i, v := range bytes {
		result += int64(v) << uint(8*i)
	}

	return result, nil
}
