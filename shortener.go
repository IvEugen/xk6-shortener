package shortener

import (
	"math/rand"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/shortener", new(Shortener))
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

type Shortener struct{}

func (*Shortener) GeneratorURL(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
