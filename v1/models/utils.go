package models

import (
	"math/rand"
	"time"
)

// RandomStringGenerator - structure of random string generator with saved base and random engine
type RandomStringGenerator struct {
	pool 	string
	gen		*rand.Rand
}

// NewRandomStringGenerator - returns new instance of RandomStringGenerator with given base
func NewRandomStringGenerator(base string) *RandomStringGenerator {
	return &RandomStringGenerator{
		base,
		rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// Generate - generates random string from registered base with given length
func (rsg *RandomStringGenerator) Generate(length int) (randomString string) {
	if length < 1 {
		return
	}
	randomBytes := make([]byte, length)
	for i, _ := range randomBytes {
		randomBytes[i] = rsg.pool[rsg.gen.Intn(len(rsg.pool))]
	}
	randomString = string(randomBytes)
	return
}
