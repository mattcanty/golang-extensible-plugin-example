package main

// PluginConfig which configures this plugin
var PluginConfig config

// Config for this plugin
type config struct {
	NestedStuff nestedItem `yaml:"nestedStuff"`
}

type nestedItem struct {
	Name string
}

// ReadChanges reads changes this plugin will make
func ReadChanges() {
	println("plugin-3 ReadChanges().")

	println("  I know about complex nested objects")

}

// WriteChanges writes changes
func WriteChanges() {
	println("plugin-3 WriteChanges().")

	println(" ", PluginConfig.NestedStuff.Name)

}
