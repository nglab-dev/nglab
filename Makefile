ifeq ($(OS),Windows_NT)
    DETECTED_OS = Windows
else
    DETECTED_OS = $(shell uname -s)
endif

ifeq ($(DETECTED_OS),Windows)
	BIN=.\\bin\\nglab.exe
else
	BIN=./bin/nglab
endif

swag:
	@go run github.com/swaggo/swag/cmd/swag@latest init --parseDependency

build:swag
	@go build -o $(BIN) .

run:
	@go run .

watch:
	@go run github.com/cosmtrek/air@v1.51.0 \
		--build.cmd "go build -o $(BIN) ." \
		--build.bin "$(BIN)" \
		--build.delay "100" \
		--build.exclude_dir "node_modules,*.db" \
		--build.include_ext "go" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true
