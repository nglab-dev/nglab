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

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@npm install

migrate:
	@go run . migrate

build:
	@npm run build:css
	@npm run build:js
	@templ generate
	@go build -o $(BIN) .
	@echo "compiled you application with all its assets to a single binary => $(BIN)"

run:
	@go run . run

watch-templ:
	@templ generate --watch --proxy="http://localhost:8080" --open-browser=false

watch-css:
	@npm run dev:css

watch-js:
	@npm run dev:js

watch-go:
	@go run github.com/cosmtrek/air@v1.51.0 \
		--build.cmd "go build -o $(BIN) ." \
		--build.bin "$(BIN) run" \
		--build.delay "100" \
		--build.exclude_dir "node_modules" \
		--build.include_ext "go" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true

sync-assets:
	@go run github.com/cosmtrek/air@v1.51.0 \
		--build.cmd "templ generate --notify-proxy" \
		--build.bin "true" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_dir "public" \
		--build.include_ext "js,css"

watch:
	@make -j5 watch-css watch-js watch-templ watch-go sync-assets