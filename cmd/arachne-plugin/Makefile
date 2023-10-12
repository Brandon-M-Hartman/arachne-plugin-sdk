JSONFILE = ./build.json

# Check if the OS is Windows
ifeq ($(OS),Windows_NT)
	SHELL := cmd
	OP = build_on_windows
	# Extract plugin details from JSON file using PowerShell
	PLUGIN_NAME=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).PluginName")
	PLUGIN_DEV_IDENTITY=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).PluginDevIdentity")
	PLUGIN_URL=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).PluginUrl")
	PLUGIN_VERSION=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).PluginVersion")
	PLUGIN_DESCRIPTION=$(shell powershell -Command "(Get-Content -Raw -Path $(JSONFILE) | ConvertFrom-Json).PluginDescription")
	# Get the current date and time
	PLUGIN_BUILDTIME=$(shell powershell -Command "Get-Date -Format o")
else
	SHELL := /bin/bash
	OP = build_on_linux

	# Check if jq is installed
ifeq ($(shell command -v jq 2> /dev/null),)
	# Extract plugin details from JSON file using sed, which isn't great but it's an okay fallback method since our JSON is very simple here. This probably breaks if we do anything weird in the fields.
	JSON_METHOD=awk
	PLUGIN_NAME=$(shell sed -n 's/.*"PluginName":\s*"\([^"]*\)".*/\1/p' $(JSONFILE))
	PLUGIN_DEV_IDENTITY=$(shell sed -n 's/.*"PluginDevIdentity":\s*"\([^"]*\)".*/\1/p' $(JSONFILE))
	PLUGIN_URL=$(shell sed -n 's/.*"PluginUrl":\s*"\([^"]*\)".*/\1/p' $(JSONFILE))
	PLUGIN_VERSION=$(shell sed -n 's/.*"PluginVersion":\s*"\([^"]*\)".*/\1/p' $(JSONFILE))
	PLUGIN_DESCRIPTION=$(shell sed -n 's/.*"PluginDescription":\s*"\([^"]*\)".*/\1/p' $(JSONFILE))

else
	# Extract plugin details from JSON file using jq
	JSON_METHOD=jq
	PLUGIN_NAME=$(shell jq -r '.PluginName' $(JSONFILE))
	PLUGIN_DEV_IDENTITY=$(shell jq -r '.PluginDevIdentity' $(JSONFILE))
	PLUGIN_URL=$(shell jq -r '.PluginUrl' $(JSONFILE))
	PLUGIN_VERSION=$(shell jq -r '.PluginVersion' $(JSONFILE))
	PLUGIN_DESCRIPTION=$(shell jq -r '.PluginDescription' $(JSONFILE))
endif

	# Get the current date and time
	PLUGIN_BUILDTIME=$(shell date)



endif

# Build the binary
build: build-binary

# Update the go.mod file
update-mods:
	@echo "Updating plugin go.mod..."
	@go mod tidy
	@go get -u ./...

# Build the binary
build-binary:
	# Build the WASM binary with TinyGo
	
	@echo "----------"
	@echo "Building Plugin with TinyGo"
	@echo "Build OS: $(OS)"
	@echo "Build OP: $(OP)"
	@echo "JSON Extraction Method: $(JSON_METHOD)"
	@echo "PLUGIN_NAME: $(PLUGIN_NAME)"
	@echo "PLUGIN_DEV_IDENTITY: $(PLUGIN_DEV_IDENTITY)"
	@echo "PLUGIN_URL: $(PLUGIN_URL)"
	@echo "PLUGIN_VERSION: $(PLUGIN_VERSION)"
	@echo "PLUGIN_DESCRIPTION: $(PLUGIN_DESCRIPTION)"	
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