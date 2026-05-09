package random

import (
	"math/rand"
	"time"
)

func NewRandomString(size int) string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	char := []rune("QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm0123456789")

	b := make([]rune, size)
	for i := range b {
		b[i] = char[rand.Intn(len(char))]
	}
	return string(b)

}
