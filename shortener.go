package shortener

import (
	"math/rand"

	"go.k6.io/k6/js/modules"
)

const version = "v0.0.1"

func init() {
	modules.Register("k6/x/shortener", &Short{
		Version: version,
	})
}

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
