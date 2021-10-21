package main

import (
	"math/rand"
)

func notSoRandom() bool {
	// rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(2) == 1
}
