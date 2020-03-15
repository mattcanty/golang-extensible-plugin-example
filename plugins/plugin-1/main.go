package main

// PluginConfig which configures this plugin
var PluginConfig config

type config struct {
	Message string
}

// ReadChanges reads changes this plugin will make
func ReadChanges() {
	println("plugin-1 ReadChanges().")

	println("  Will print message:", PluginConfig.Message)
}

// WriteChanges writes changes
func WriteChanges() {
	println("plugin-1 WriteChanges().")

	println(" ", PluginConfig.Message)
}
