SHELL := /bin/bash

# 项目与输出
PROJECT := cos-auto-upload
BIN_DIR := bin
BIN := $(BIN_DIR)/$(PROJECT)

# 可覆盖的编译参数
OS   ?= $(shell go env GOOS)
ARCH ?= $(shell go env GOARCH)
CGO_ENABLED ?= 0

.PHONY: all build build-all fmt vet clean help

# 默认一键编译：格式化 + 静态检查 + 构建
# 默认编译所有常见平台
all: build-all

build: fmt vet
	@mkdir -p $(BIN_DIR)
	@echo "Building $(BIN) for $(OS)/$(ARCH)..."
	@GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=$(CGO_ENABLED) \
		go build -o $(BIN) .
	@echo "OK -> $(BIN)"

# 多平台编译默认矩阵
# 包含: linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64
PLATFORMS ?= linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64

build-all: fmt vet
	@mkdir -p $(BIN_DIR)
	@echo "Building for platforms: $(PLATFORMS)"
	@set -e; \
	for plat in $(PLATFORMS); do \
	  os=$${plat%/*}; arch=$${plat#*/}; \
	  suffix="$$os-$$arch"; \
	  bin_name="$(PROJECT)-$$suffix"; \
	  if [ "$$os" = "windows" ]; then bin_name="$$bin_name.exe"; fi; \
	  out="$(BIN_DIR)/$$bin_name"; \
	  echo "Building $$out ..."; \
	  GOOS="$$os" GOARCH="$$arch" CGO_ENABLED=$(CGO_ENABLED) \
	    go build -o "$$out" .; \
	 done; \
	 echo "All done -> $(BIN_DIR)"

fmt:
	@go fmt ./...

vet:
	@go vet ./...

clean:
	@rm -rf $(BIN_DIR)
	@echo "Cleaned $(BIN_DIR)"

help:
	@echo "Targets:"
	@echo "  make           一键编译所有平台 (build-all)"
	@echo "  make build     仅构建，依赖 fmt、vet"
	@echo "  make fmt       代码格式化"
	@echo "  make vet       静态检查"
	@echo "  make clean     清理构建产物"
	@echo "Variables (可覆盖): OS=$(OS) ARCH=$(ARCH) CGO_ENABLED=$(CGO_ENABLED) PLATFORMS=$(PLATFORMS)"
	@echo "示例: make build OS=linux ARCH=amd64"
