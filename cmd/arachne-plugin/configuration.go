package main

import (
	"encoding/json"
)

const dev_identity = "Brandon Hartman"

const plugin_url = "https://github.com/getarachne/arachne-plugin-sdk"

// Should be a string, not a number.
const pluginVersion = "1.0"

const description = "A demo plugin for Arachne, to help developers understand how they can make their own plugins."

type PluginInfo struct {
	DeveloperIdentity string `json:"developer_identity"`
	PluginUrl         string `json:"plugin_url"`
	PluginVersion     string `json:"plugin_version"`
	PluginDescription string `json:"plugin_description"`
}

// describePlugin will return serialized information about the plugin
func DescribePlugin(payload []byte) ([]byte, error) {
	// Create a new instance of the PluginInfo message
	info := &PluginInfo{
		DeveloperIdentity: dev_identity,
		PluginUrl:         plugin_url,
		PluginVersion:     pluginVersion,
		PluginDescription: description,
	}

	// Marshal the message to JSON bytes
	data, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	// Return the serialized bytes
	return data, nil
}
