// randgen
package gorest

import (
	"math/rand"
	"time"
)

var alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

// generates a random string of fixed size
func RandGen(size int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)
}
