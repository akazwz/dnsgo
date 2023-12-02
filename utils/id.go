package utils

import (
	"github.com/matoous/go-nanoid/v2"
)

var ID = &id{}

type id struct{}

const Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func (i *id) Generate(size int) (string, error) {
	id, err := gonanoid.Generate(Alphabet, size)
	if err != nil {
		return "", err
	}
	return id, nil
}
