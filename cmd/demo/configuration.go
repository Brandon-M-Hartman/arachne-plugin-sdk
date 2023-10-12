package main

type Config struct {
	plugin_name string
	plugin_path string
	LogLevel    string
	LogFilePath string
	AddSource   bool
}

var config Config

func init() {
	config = Config{}
	config.plugin_name = "arachne-plugin"
	config.plugin_path = "../wasm/arachne-plugin.wasm"
	config.LogLevel = "Debug"
	config.LogFilePath = "./arachne-plugin.log"

	// Set to true to have every log line print the source of the call - spammy, but can be useful.
	config.AddSource = false
}
