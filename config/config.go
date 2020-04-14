package config

import (
	"encoding/json"
	"os"
	"strings"
)

type Config struct {
	DiceName string
	Faces    []string
}

func FromEnv() (Config, error) {
	var faces []string
	if err := json.NewDecoder(strings.NewReader(os.Getenv("FACES"))).Decode(&faces); err != nil {
		return Config{}, err
	}
	return Config{
		DiceName: os.Getenv("DICE_NAME"),
		Faces:    faces,
	}, nil
}
