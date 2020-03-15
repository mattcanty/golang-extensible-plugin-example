package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"plugin"

	"gopkg.in/yaml.v2"
)

type pluginsConfig struct {
	Plugins []pluginConfig
}

type pluginConfig struct {
	Name       string
	Parameters string
}

func main() {
	println("This is a test application to demonstrate loading plugins at runtime.")

	cfg := loadConfig()

	for _, plugin := range cfg.Plugins {
		runPlugin(fmt.Sprintf("%s.so", plugin.Name), plugin.Parameters)
	}
}

func runPlugin(name string, parameters string) {
	plugin, err := plugin.Open(name)

	if err != nil {
		panic(err)
	}
	cfg, err := plugin.Lookup("PluginConfig")
	if err != nil {
		panic(err)
	}
	readFunc, err := plugin.Lookup("ReadChanges")
	if err != nil {
		panic(err)
	}
	writeFunc, err := plugin.Lookup("WriteChanges")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(parameters), cfg)
	if err != nil {
		panic(err)
	}

	readFunc.(func())()
	writeFunc.(func())()
}

func loadConfig() pluginsConfig {
	contents, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to find config %s\n", err)
	}

	var cfg pluginsConfig

	err = yaml.Unmarshal(contents, &cfg)
	if err != nil {
		log.Fatalf("Failed to Unmarshal YAMl. %s\n", err)
	}

	return cfg
}
