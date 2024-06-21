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

GO_ENV=CGO_ENABLED=0 GO111MODULE=on
GO=$(GO_ENV) go

migrate:
	@$(GO) run . migrate

swag:
	@$(GO) run github.com/swaggo/swag/cmd/swag@latest init

build:swag
	@$(GO) build -o $(BIN) .

run:
	@$(GO) run . run

watch:
	@$(GO) run github.com/cosmtrek/air@v1.51.0 \
		--build.cmd "$(GO) build -o $(BIN) ." \
		--build.bin "$(BIN) run" \
		--build.delay "100" \
		--build.exclude_dir "node_modules" \
		--build.include_ext "go" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true
