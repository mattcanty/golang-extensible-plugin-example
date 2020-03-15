package main

// PluginConfig which configures this plugin
var PluginConfig config

// Config for this plugin
type config struct {
	Secret string
}

// ReadChanges reads changes this plugin will make
func ReadChanges() {
	println("plugin-2 ReadChanges().")

	println("  I know the secret")
}

// WriteChanges writes changes
func WriteChanges() {
	println("plugin-2 WriteChanges().")

	println(" ", PluginConfig.Secret)
}
