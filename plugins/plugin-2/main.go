package main

import (
	"log"

	"gopkg.in/yaml.v2"
)

var cfg Config

func init() {
	println("plugin-2 init().")
}

// Config for this plugin
type Config struct {
	Secret string
}

// Load loads the config
func Load(parameters string) {
	println("plugin-1 Load().")

	err := yaml.Unmarshal([]byte(parameters), &cfg)
	if err != nil {
		log.Fatalf("Failed to Unmarshal YAMl. %s\n", err)
	}
}

// ReadChanges reads changes this plugin will make
func ReadChanges() {
	println("plugin-2 ReadChanges().")

	println("  I know the secret")
}

// WriteChanges writes changes
func WriteChanges() {
	println("plugin-2 WriteChanges().")

	println(" ", cfg.Secret)
}
