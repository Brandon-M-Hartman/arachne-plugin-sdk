package plugin_funcs

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"os"
	"reflect"
	"strings"

	"log/slog"

	arachne_logging "github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/logging"
	arachne_plugin_scaffold "github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/scaffold"
	origWazero "github.com/tetratelabs/wazero"
	"github.com/wapc/wapc-go"
	"github.com/wapc/wapc-go/engines/wazero"
)

// Invoke and return data from a validated plugin
func InvokePluginByInstance(instance *wapc.Instance, input []byte, function string) ([]byte, error) {

	// Invoke the function in the wasm module with the input as argument
	result, err := (*instance).Invoke(context.Background(), function, input)

	// Check for errors in invoking the function
	if err != nil {
		return nil, err
	}

	return result, nil

}

func UnloadPlugin(a *arachne_plugin_scaffold.ArachnePlugin) {
	err := (*a.Instance).Close(*a.Context)
	if err != nil {
		slog.Error("Error occurred while closing plugin instance.", "type", "PLUGIN", "plugin_name", a.PluginInfo.PluginName, "error", err.Error())
		return
	}

	err = (*a.Module).Close(*a.Context)
	if err != nil {
		slog.Error("Error occurred while closing plugin module.", "type", "PLUGIN", "plugin_name", a.PluginInfo.PluginName, "error", err.Error())
		return
	}
}

// Load an individual plugin
func LoadPlugin(plugin_name string, plugin_path string, host_function wapc.HostCallHandler) (*arachne_plugin_scaffold.ArachnePlugin, error) {

	// Create a new context
	ctx := context.Background()

	// Read the wasm plugin file
	guest, err := os.ReadFile(plugin_path)

	// Check for errors in reading the file
	if err != nil {
		return &arachne_plugin_scaffold.ArachnePlugin{}, err
	}

	// Initialize the wasm engine
	engine := wazero.Engine()

	// Create a new module with the wasm engine, host function, wasm file and module configuration, logging to the global logger.
	module, err := CreateModule(&engine, &guest, ctx, host_function)

	// Check for errors in instantiating the module
	if err != nil {
		slog.Error("Error occurred while instantiating plugin module.", "type", "PLUGIN", "plugin_name", plugin_name, "error", err.Error())
		return &arachne_plugin_scaffold.ArachnePlugin{}, err
	}

	fsvar := os.DirFS("./expose_to_plugin")

	RegConfig := origWazero.NewModuleConfig().
		WithFSConfig(origWazero.NewFSConfig().WithFSMount(fsvar, "/"))

	(*module).(*wazero.Module).WithConfig(func(config origWazero.ModuleConfig) origWazero.ModuleConfig {
		return RegConfig
	})

	// Instantiate the instance
	instance, err := (*module).Instantiate(ctx)

	// Check for errors in instantiating the instance
	if err != nil {
		slog.Error("Error occurred while instantiating plugin instance.", "type", "PLUGIN", "plugin_name", plugin_name, "error", err.Error())
		return &arachne_plugin_scaffold.ArachnePlugin{}, err
	}

	plugin := arachne_plugin_scaffold.ArachnePlugin{
		Context:  &ctx,
		Module:   module,
		Instance: &instance,
	}

	val_err := ValidatePlugin(&plugin)
	if val_err != nil {
		slog.Error("Plugin validation failed.", "type", "PLUGIN", "plugin_name", plugin_name, "plugin_path", plugin_path, "error", val_err.Error())
	}

	return &plugin, nil
}

func CreateModule(engine *wapc.Engine, guest *[]byte, ctx context.Context, HostFunction wapc.HostCallHandler) (*wapc.Module, error) {

	module, err := (*engine).New(ctx, HostFunction, *guest, &wapc.ModuleConfig{
		Logger: wapc.PrintlnLogger,
		Stdout: &arachne_logging.SlogWriter{Logger: slog.Default()},
		Stderr: os.Stderr,
	})

	return &module, err
}

// Run plugin validation
func ValidatePlugin(p *arachne_plugin_scaffold.ArachnePlugin) error {

	slog.Debug("Validating a plugin...", "type", "PLUGIN")

	instance := p.Instance

	// Invoke the function in the wasm module with the input as argument
	resp, err := InvokePluginByInstance(instance, []byte(""), "DescribePlugin")

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

	// Unmarshal the JSON response into a map to convert it into a Go data structure
	var data map[string]interface{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return err
	}

	// Iterate over the keys in the map and check if each key corresponds to a field in the PluginInfo struct
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
	slog.Debug("Developer Identity: " + info.PluginDevIdentity)
	slog.Debug("Plugin URL: " + info.PluginUrl)
	slog.Debug("Plugin Version: " + info.PluginVersion)
	slog.Debug("Plugin Description: " + info.PluginDescription)
	for i, function := range info.PluginFunctions {
		// Sanitize the function name
		info.PluginFunctions[i] = html.EscapeString(function)
		slog.Debug("Available Plugin Function: " + info.PluginFunctions[i])
	}

	p.PluginInfo = info

	// Log the result of the function invocation
	slog.Debug("Plugin passed validation.", "type", "PLUGIN", "plugin", p.PluginInfo.PluginName)

	return nil
}
