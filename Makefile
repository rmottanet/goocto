# Makefile for GoOcto

# Variáveis
BUILD_DIR := ./
CMD_DIR := ./cli/cmd
MAIN_FILE := $(CMD_DIR)/main.go
OUTPUT_NAME := goocto

# Comandos
.PHONY: all build clean

all: download-deps build

build:
	go build -o $(BUILD_DIR)/$(OUTPUT_NAME) $(MAIN_FILE)

clean:
	rm -f $(BUILD_DIR)/$(OUTPUT_NAME)

download-deps:
	go mod download
