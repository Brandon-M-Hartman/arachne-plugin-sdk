package plugin_funcs

import (
	"context"
	"os"

	"log/slog"

	arachne_logging "github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/logging"
	arachne_plugin_scaffold "github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/scaffold"
	origWazero "github.com/tetratelabs/wazero"
	"github.com/wapc/wapc-go"
	"github.com/wapc/wapc-go/engines/wazero"
)

// Invoke and return data from a validated plugin
func InvokePluginByInstance(instance *wapc.Instance, input string, function string) ([]byte, error) {

	// Invoke the function in the wasm module with the input as argument
	result, err := (*instance).Invoke(context.Background(), function, []byte(input))

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
func LoadPlugin(plugin_name string, plugin_path string, host_function wapc.HostCallHandler) (context.Context, wapc.Module, wapc.Instance, error) {

	// Create a new context
	ctx := context.Background()

	// Read the wasm plugin file
	guest, err := os.ReadFile(plugin_path)

	// Check for errors in reading the file
	if err != nil {
		return nil, nil, nil, err
	}

	// Initialize the wasm engine
	engine := wazero.Engine()

	// Create a new module with the wasm engine, host function, wasm file and module configuration, logging to the global logger.
	module, err := CreateModule(&engine, &guest, ctx, host_function)

	// Check for errors in instantiating the module
	if err != nil {
		slog.Error("Error occurred while instantiating plugin module.", "type", "PLUGIN", "plugin_name", plugin_name, "error", err.Error())
		return nil, nil, nil, err
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
		return nil, nil, nil, err
	}

	return ctx, *module, instance, nil

}

func CreateModule(engine *wapc.Engine, guest *[]byte, ctx context.Context, HostFunction wapc.HostCallHandler) (*wapc.Module, error) {

	module, err := (*engine).New(ctx, HostFunction, *guest, &wapc.ModuleConfig{
		Logger: wapc.PrintlnLogger,
		Stdout: &arachne_logging.SlogWriter{Logger: slog.Default()},
		Stderr: os.Stderr,
	})

	return &module, err
}
