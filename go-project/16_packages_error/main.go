package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	// config fields goes here ...
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.Wrap(err, "can't open config file")
	}

	defer file.Close()

	config := &Config{}

	// Parse the file here

	return config, nil
}

func main() {
	config, err := readConfig("/path/tp/fake/file.yml")

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)

		os.Exit(1)
	}

	// Normal operation
	fmt.Println(config)
}
