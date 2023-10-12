package main

import (
	"github.com/getarachne/arachne-plugin-sdk/pkg/plugin_requirements"
	wapc "github.com/wapc/wapc-guest-tinygo"
)

// Keep main.go pretty simple so you can follow the code of your plugin easier.
func main() {

	// Add functions to the slice. Whenever we write a new function, we need to add it here.
	plugin_requirements.Arachne_Plugin_Register.Funcs = append(plugin_requirements.Arachne_Plugin_Register.Funcs, Hello)
	plugin_requirements.Arachne_Plugin_Register.Funcs = append(plugin_requirements.Arachne_Plugin_Register.Funcs, Goodbye)

	// Only put non-mandatory functions here, the mandatory ones are implemented in ImplementRequiredFunctions, called below.
	plugin_requirements.Arachne_Plugin_Register.FuncMap = wapc.Functions{
		"Hello":   Hello,
		"Goodbye": Goodbye,
	}

	// This line is mandatory, and should go AFTER you've declared your functions above. Without your plugin implementing the scaffolding functions, Arachne won't accept it and it won't be allowed to run.
	plugin_requirements.ImplementRequiredFunctions(&plugin_requirements.Arachne_Plugin_Register)

	wapc.RegisterFunctions(plugin_requirements.Arachne_Plugin_Register.FuncMap)
}
