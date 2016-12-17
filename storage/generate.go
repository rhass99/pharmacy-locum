package storage

import (
	//"fmt"
	"math/rand"
	//"regexp"
	"crypto/sha1"
	"io"
	"time"
)

// validName matches a valid name string.
//var validName = regexp.MustCompile(`^[a-zA-Z0-9\-]+$`)

// randId returns a string of random letters.

func RandId(l int) string {
	n := make([]byte, l)
	for i := range n {
		n[i] = 'a' + byte(rand.Intn(26))
	}
	return string(n)
}

func WeakPasswordHash(password string) []byte {
	hash := sha1.New()
	io.WriteString(hash, password)
	return hash.Sum(nil)
}

func init() {
	// Seed number generator with the current time.
	rand.Seed(time.Now().UnixNano())
}
