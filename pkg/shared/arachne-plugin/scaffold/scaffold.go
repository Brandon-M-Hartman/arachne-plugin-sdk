package arachne_plugin_scaffold

import (
	"context"

	"github.com/wapc/wapc-go"
)

type ArachnePlugin struct {
	Context    *context.Context
	Module     *wapc.Module
	Instance   *wapc.Instance
	PluginInfo *PluginInfo
}

type PluginInfo struct {
	PluginName        string   `json:"plugin_name" validate:"required,alphanumunicode,max=40,min=1"`
	DeveloperIdentity string   `json:"developer_identity" validate:"required,alphanumunicode,max=40,min=1"`
	PluginUrl         string   `json:"plugin_url" validate:"url,max=40,min=6"`
	PluginVersion     string   `json:"plugin_version" validate:"required,semver,max=14,min=5"`
	PluginDescription string   `json:"plugin_description" validate:"required,alphanumunicode,max=256,min=1"`
	PluginFunctions   []string `json:"plugin_functions" validate:"required,alphanum,max=256,min=1"`
}

// Host function to handle calls from the wasm module. No real functionality here, just for example purposes.
func BasicHostFunc(_ context.Context, binding, namespace, operation string, payload []byte) ([]byte, error) {
	// Route the payload to any custom functionality accordingly.
	// You can even route to other plugins!
	switch namespace {
	case "demo":

		switch operation {

		case "MESSAGE":
			// If the operation is "MESSAGE", return a predefined string.
			reply := "MESSAGE_REPLY"
			return []byte(reply), nil
		}
	}
	// If no operation matches, return a default string
	return []byte("MESSAGE_DEFAULT_REPLY"), nil
}
