package main

import (
	"log/slog"

	arachne_logging "github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/logging"
	"github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/plugin_funcs"
	arachne_plugin_scaffold "github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/scaffold"
	arachne_plugin_validation "github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/validation"
)

func main() {

	arachne_logging.InitLogger(config.AddSource, config.LogLevel, config.LogFilePath, 512, 2, 30, true)

	ctx, module, instance, err := plugin_funcs.LoadPlugin(config.plugin_name, config.plugin_path, arachne_plugin_scaffold.BasicHostFunc)

	if err != nil {
		slog.Error("Plugin loading failed.", "type", "PLUGIN", "plugin_name", config.plugin_name, "plugin_path", config.plugin_path, "error", err.Error())
	}

	plugin := arachne_plugin_scaffold.ArachnePlugin{
		Context:  &ctx,
		Module:   &module,
		Instance: &instance,
	}

	val_err := arachne_plugin_validation.ValidatePlugin(&plugin)
	if val_err != nil {
		slog.Error("Plugin validation failed.", "type", "PLUGIN", "plugin_name", config.plugin_name, "plugin_path", config.plugin_path, "error", val_err.Error())
	}

	plugin_funcs.UnloadPlugin(&plugin)
}
