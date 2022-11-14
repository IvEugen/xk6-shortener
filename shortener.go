package shortener

import (
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

type Short struct {
	Version string
}

func (*Short) GeneratorURL(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
