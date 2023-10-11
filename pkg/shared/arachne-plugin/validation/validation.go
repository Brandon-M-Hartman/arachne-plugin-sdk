package arachne_plugin_validation

import (
	"encoding/json"
	"fmt"
	"html"
	"log/slog"
	"reflect"
	"strings"

	"github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/plugin_funcs"
	arachne_plugin_scaffold "github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/scaffold"
)

// Run plugin validation
func ValidatePlugin(p *arachne_plugin_scaffold.ArachnePlugin) error {

	slog.Debug("Validating a plugin.", "type", "PLUGIN")

	instance := p.Instance

	// Invoke the function in the wasm module with the input as argument
	resp, err := plugin_funcs.InvokePluginByInstance(instance, "", "DescribePlugin")

	// Create a new instance of the PluginInfo message
	info := &arachne_plugin_scaffold.PluginInfo{}

	// Check for errors in invoking the function
	if err != nil {
		return err
	}

	// Validate that the response is a valid JSON
	if !json.Valid(resp) {
		return fmt.Errorf("invalid JSON received from the plugin")
	}

	// Unmarshal the data into a map first to check for unexpected fields
	var data map[string]interface{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return err
	}

	// Check for unexpected fields in the JSON data
	for key := range data {
		structType := reflect.TypeOf(info).Elem()
		_, ok := structType.FieldByNameFunc(func(s string) bool {
			field, _ := structType.FieldByName(s)
			return strings.EqualFold(field.Tag.Get("json"), key)
		})
		if !ok {
			return fmt.Errorf("unexpected field %s in JSON data", key)
		}
	}

	// Now unmarshal the data into the PluginInfo struct
	err = json.Unmarshal(resp, info)
	if err != nil {
		return err
	}

	// Print the plugin description
	slog.Debug("Plugin Name: " + info.PluginName)
	slog.Debug("Developer Identity: " + info.DeveloperIdentity)
	slog.Debug("Plugin URL: " + info.PluginUrl)
	slog.Debug("Plugin Version: " + info.PluginVersion)
	slog.Debug("Plugin Description: " + info.PluginDescription)
	for _, function := range info.PluginFunctions {
		sanitizedFunction := html.EscapeString(function)
		slog.Debug("Available Plugin Function: " + sanitizedFunction)
	}

	p.PluginInfo = info

	// Log the result of the function invocation
	slog.Debug("Plugin passed validation.", "type", "PLUGIN", "plugin", p.PluginInfo.PluginName)

	return nil
}
