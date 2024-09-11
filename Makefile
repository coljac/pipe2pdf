# Makefile

APP_NAME := pipe2pdf
VERSION := 1.0.0

# Define the output directories
OUTPUT_DIR := dist
MACOS_DIR := $(OUTPUT_DIR)/macos
LINUX_AMD64_DIR := $(OUTPUT_DIR)/linux_amd64
LINUX_ARM64_DIR := $(OUTPUT_DIR)/linux_arm64
WINDOWS_DIR := $(OUTPUT_DIR)/windows

# Define the build commands
build_macos_amd64:
	mkdir -p $(MACOS_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(MACOS_DIR)/$(APP_NAME)
	tar -czvf $(MACOS_DIR)/$(APP_NAME)_$(VERSION)_macos_amd64.tar.gz -C $(MACOS_DIR) $(APP_NAME)

build_macos_arm64:
	mkdir -p $(MACOS_DIR)
	GOOS=darwin GOARCH=arm64 go build -o $(MACOS_DIR)/$(APP_NAME)
	tar -czvf $(MACOS_DIR)/$(APP_NAME)_$(VERSION)_macos_arm64.tar.gz -C $(MACOS_DIR) $(APP_NAME)

build_linux_amd64:
	mkdir -p $(LINUX_AMD64_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(LINUX_AMD64_DIR)/$(APP_NAME)
	tar -czvf $(LINUX_AMD64_DIR)/$(APP_NAME)_$(VERSION)_linux_amd64.tar.gz -C $(LINUX_AMD64_DIR) $(APP_NAME)

build_windows:
	mkdir -p $(WINDOWS_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_DIR)/$(APP_NAME).exe
	zip -j $(WINDOWS_DIR)/$(APP_NAME)_$(VERSION)_windows_amd64.zip $(WINDOWS_DIR)/$(APP_NAME).exe

# Define the default target
all: build_macos_amd64 build_macos_arm64 build_linux_amd64 build_linux_arm64 build_windows

# Clean up the build artifacts
clean:
	rm -rf $(OUTPUT_DIR)

.PHONY: all build_macos_amd64 build_macos_arm64 build_linux_amd64 build_linux_arm64 build_windows clean
