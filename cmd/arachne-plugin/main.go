package main

import (
	wapc "github.com/wapc/wapc-guest-tinygo"
)

// Keep main.go pretty simple so you can follow the code of your plugin easier.
func main() {

	arachne_plugin = ArachnePluginRegister{}

	// Add functions to the slice. Whenever we write a new function, we need to add it here.
	arachne_plugin.funcs = append(arachne_plugin.funcs, Hello)
	arachne_plugin.funcs = append(arachne_plugin.funcs, Goodbye)

	// Only put non-mandatory functions here, the mandatory ones are implemented in ImplementRequiredFunctions, called below.
	arachne_plugin.funcMap = wapc.Functions{
		"Hello":   Hello,
		"Goodbye": Goodbye,
	}

	// This line is mandatory, and should go AFTER you've declared your functions above. Without your plugin implementing the scaffolding functions, Arachne won't accept it and it won't be allowed to run.
	ImplementRequiredFunctions(&arachne_plugin)

	wapc.RegisterFunctions(arachne_plugin.funcMap)
}
