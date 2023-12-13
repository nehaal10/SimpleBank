package util

import (
	"math/rand"
	"strings"
	"time"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandonInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alpha)

	for i := 0; i < n; i++ {
		c := alpha[rand.Intn((k))]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandOwner() string {
	return RandomString(6)
}

func RandomBalnce() int64 {
	return RandonInt(100, 1000)
}

func RandomCurrency() string {
	curr := []string{"EUR", "USD", "INR", "CAD"}
	n := len(curr)
	return curr[rand.Intn(n)]
}
