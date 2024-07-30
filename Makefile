APP_NAME=bin/nglab

ifeq ($(OS),Windows_NT)
    DETECTED_OS = Windows
else
    DETECTED_OS = $(shell uname -s)
endif

ifeq ($(DETECTED_OS),Windows)
	EXT=.exe
endif

swag:
	@go run github.com/swaggo/swag/cmd/swag@latest init --parseDependency

build:swag
	@go build -o $(APP_NAME)$(EXT) .

run:
	@go run .

dev:
	@go run github.com/cosmtrek/air@v1.51.0 \
		--build.cmd "go build -o $(APP_NAME)$(EXT) ." \
		--build.bin "$(APP_NAME)$(EXT)" \
		--build.delay "100" \
		--build.exclude_dir "node_modules,*.db" \
		--build.include_ext "go,html" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true
