# 项目名称
APP_NAME := online-shopping

# 源代码目录
SRC_DIR := ./internal

# 编译后输出的二进制文件
BIN := $(APP_NAME)

# Go 编译器
GO := go

# 编译选项
GOFLAGS :=

# 目标：编译
.PHONY: all
all: build

# 编译项目
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	$(GO) build $(GOFLAGS) -o $(BIN) ./main.go

# 运行项目
.PHONY: run
run: build
	@echo "Running $(APP_NAME)..."
	./$(BIN)

# 清理编译生成的文件
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(BIN)

## up	- Boots up docker dev env
up:
	@echo "$(OK_COLOR)==> Booting up docker environment$(NO_COLOR)"
	docker-compose up -d

stop:
	@printf "$(OK_COLOR)==> Stopping docker containers$(NO_COLOR)\n"
	@docker-compose down
