JSONFILE = ./build.json

ifeq ($(OS),Windows_NT)
	SHELL = cmd
	OP = build_on_windows
	PLUGIN_NAME=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).PluginName")
	DEVELOPER_IDENTITY=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).DeveloperIdentity")
	PLUGIN_URL=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).PluginUrl")
	PLUGIN_VERSION=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).PluginVersion")
	PLUGIN_DESCRIPTION=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).PluginDescription")
	PLUGIN_BUILDTIME=$(shell powershell -Command "Get-Date -Format o")
else
	SHELL = /bin/bash
	OP = build_on_linux
	PLUGIN_NAME=$(shell jq -r '.PluginName' $(JSONFILE))
	DEVELOPER_IDENTITY=$(shell jq -r '.DeveloperIdentity' $(JSONFILE))
	PLUGIN_URL=$(shell jq -r '.PluginUrl' $(JSONFILE))
	PLUGIN_VERSION=$(shell jq -r '.PluginVersion' $(JSONFILE))
	PLUGIN_DESCRIPTION=$(shell jq -r '.PluginDescription' $(JSONFILE))
	PLUGIN_BUILDTIME=$(shell date)	
endif

build: build-binary

build-binary:
	@echo "----------"
	@echo "Building Plugin with TinyGo"
	@echo "Build OS: $(OS)"
	@echo "Build OP: $(OP)"
	@echo "----------"
	@echo "Starting WASM build..."
	@tinygo build -o ../../bin/wasm/arachne-plugin.wasm -scheduler=none --no-debug \
	-ldflags="-X 'main.PluginName=$(PLUGIN_NAME)' \
			-X 'main.PluginDevIdentity=$(DEVELOPER_IDENTITY)' \
			-X 'main.PluginUrl=$(PLUGIN_URL)' \
			-X 'main.PluginVersion=$(PLUGIN_VERSION)' \
			-X 'main.PluginDescription=$(PLUGIN_DESCRIPTION)' \
			-X 'main.PluginBuildTime=$(PLUGIN_BUILDTIME)'" \
	-target=wasi ./
	@echo "WASM build finished!"