package config

import (
	"os"
	"strings"
)

type Config struct {
	DiceName string
	Faces    []string
}

func FromEnv() (Config, error) {
	return Config{
		DiceName: os.Getenv("DICE_NAME"),
		Faces:    strings.Split(os.Getenv("DICE_FACES"), ","),
	}, nil
}
