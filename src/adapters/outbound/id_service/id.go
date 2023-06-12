package id_service

import (
	"errors"
	"math/rand"
)

var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Create(size int) (error, string) {
	if size < 6 {
		return errors.New("size must be bigger than 6"), ""
	}

	str := make([]rune, size)

	for i := range str {
		str[i] = runes[rand.Intn(len(runes))]
	}

	return nil, string(str)
}
