package main

import (
	"encoding/json"

	arachne_plugin_scaffold "github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/scaffold"
	"github.com/wapc/wapc-guest-tinygo"
)

/*
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
You should really never edit this file. The agent expects plugins to operate in a very specific way, and this code handles that for you.
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
Again.. don't edit this file.
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/

type ArachnePluginRegister struct {
	funcs   []interface{}
	funcMap wapc.Functions
}

// Global plugin instance

var arachne_plugin ArachnePluginRegister

// Attaches the requires function scaffolding to an ArachnePlugin struct - these are MANDATORY for the Arachne agent to accept the plugin as valid!
func ImplementRequiredFunctions(arachne_plugin *ArachnePluginRegister) {
	arachne_plugin.funcMap["DescribePlugin"] = DescribePlugin
	arachne_plugin.funcMap["GetFunctions"] = GetFunctions
}

// describePlugin will return serialized information about the plugin
func DescribePlugin(payload []byte) ([]byte, error) {

	funcNames := make([]string, 0, len(arachne_plugin.funcMap))
	for name := range arachne_plugin.funcMap {
		funcNames = append(funcNames, name)
	}

	// Create a new instance of the PluginInfo message
	info := &arachne_plugin_scaffold.PluginInfo{
		PluginName:        plugin_name,
		DeveloperIdentity: dev_identity,
		PluginUrl:         plugin_url,
		PluginVersion:     pluginVersion,
		PluginDescription: description,
		PluginFunctions:   funcNames,
	}

	// Marshal the message to JSON bytes
	data, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	// Return the serialized bytes
	return data, nil
}

// GetFunctions returns a slice of strings containing the function names in the ArachnePlugin funcMap
func GetFunctions(payload []byte) ([]byte, error) {
	// Get the function names from the funcMap
	funcNames := make([]string, 0, len(arachne_plugin.funcMap))
	for name := range arachne_plugin.funcMap {
		funcNames = append(funcNames, name)
	}

	// Marshal the function names to JSON bytes
	data, err := json.Marshal(funcNames)
	if err != nil {
		return nil, err
	}

	// Return the serialized bytes
	return data, nil
}
