package generator

import "math/rand"

var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
var length = 10

func Rand_generate() (string, error) {
	result := make([]byte, length)

	for i := range result {
		val := rand.Intn(len(alphabet))
		result[i] = alphabet[val]
	}

	return string(result), nil
}
