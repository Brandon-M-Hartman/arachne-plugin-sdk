JSONFILE = ./build.json

ifeq ($(OS),Windows_NT)
	SHELL = cmd
	OP = build_on_windows
	PLUGIN_NAME=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).PluginName")
	PLUGIN_DEV_IDENTITY=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).PluginDevIdentity")
	PLUGIN_URL=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).PluginUrl")
	PLUGIN_VERSION=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).PluginVersion")
	PLUGIN_DESCRIPTION=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).PluginDescription")
	PLUGIN_BUILDTIME=$(shell powershell -Command "Get-Date -Format o")
else
	SHELL = /bin/bash
	OP = build_on_linux
	PLUGIN_NAME=$(shell jq -r '.PluginName' $(JSONFILE))
	PLUGIN_DEV_IDENTITY=$(shell jq -r '.PluginDevIdentity' $(JSONFILE))
	PLUGIN_URL=$(shell jq -r '.PluginUrl' $(JSONFILE))
	PLUGIN_VERSION=$(shell jq -r '.PluginVersion' $(JSONFILE))
	PLUGIN_DESCRIPTION=$(shell jq -r '.PluginDescription' $(JSONFILE))
	PLUGIN_BUILDTIME=$(shell date)	
endif

build: build-binary

update-mods:
	@echo "Updating plugin go.mod..."
	@go mod tidy
	@go get -u ./...

build-binary:
	@echo "----------"
	@echo "Building Plugin with TinyGo"
	@echo "Build OS: $(OS)"
	@echo "Build OP: $(OP)"
	@echo "----------"
	@echo "Starting WASM build..."
	@tinygo build -o ../../bin/wasm/arachne-plugin.wasm -scheduler=none --no-debug \
	-ldflags="-X 'github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/plugin_config.PluginName=$(PLUGIN_NAME)' \
			-X 'github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/plugin_config.PluginDevIdentity=$(PLUGIN_DEV_IDENTITY)' \
			-X 'github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/plugin_config.PluginUrl=$(PLUGIN_URL)' \
			-X 'github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/plugin_config.PluginVersion=$(PLUGIN_VERSION)' \
			-X 'github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/plugin_config.PluginDescription=$(PLUGIN_DESCRIPTION)' \
			-X 'github.com/getarachne/arachne-plugin-sdk/pkg/shared/arachne-plugin/plugin_config.PluginBuildTime=$(PLUGIN_BUILDTIME)'" \
	-target=wasi ./
	@echo "WASM build finished!"