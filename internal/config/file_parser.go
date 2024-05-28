package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

func parseYAML(r io.Reader, str interface{}) error {
	return yaml.NewDecoder(r).Decode(str)
}

func parseJSON(r io.Reader, str interface{}) error {
	return json.NewDecoder(r).Decode(str)
}

func parseENV(r io.Reader, _ interface{}) error {
	vars, err := godotenv.Parse(r)
	if err != nil {
		return err
	}

	for env, val := range vars {
		if err = os.Setenv(env, val); err != nil {
			return fmt.Errorf("Error load .env file: %w", err)
		}
	}

	return nil
}
