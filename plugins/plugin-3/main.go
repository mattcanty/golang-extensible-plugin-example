package main

import (
	"log"

	"gopkg.in/yaml.v2"
)

var cfg Config

func init() {
	println("plugin-3 init().")
}

// Config for this plugin
type Config struct {
	NestedStuff []nestedItem `yaml:"nestedStuff"`
}

type nestedItem struct {
	Name string
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
	println("plugin-3 ReadChanges().")
}

// WriteChanges writes changes
func WriteChanges() {
	println("plugin-3 WriteChanges().")
}
