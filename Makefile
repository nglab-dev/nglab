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
	@go install github.com/cosmtrek/air@v1.51.0

# run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
# Default url: http://localhost:8080
templ:
	@templ generate --watch --proxy="http://localhost:8080" --open-browser=false

# run air to detect any go file changes to re-build and re-run the server.
server:
	@air --build.cmd "go build -o $(BIN) ." \
	--build.bin "$(BIN) run" \
	--build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

# run tailwindcss to generate the styles.css bundle in watch mode.
watch-assets:
	npx tailwindcss -i web/assets/app.css -o ./public/assets/styles.css --watch   

# run esbuild to generate the index.js bundle in watch mode.
watch-esbuild:
	npx esbuild web/assets/index.js --bundle --outdir=public/assets --watch

# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.
sync_assets:
	@air \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "public" \
	--build.include_ext "js,css"

# start the application in development
dev:
	@make -j5 templ server watch-assets watch-esbuild sync_assets

# build the application for production. This will compile your app
# to a single binary with all its assets embedded.
build:
	@npx tailwindcss -i web/assets/app.css -o ./public/assets/styles.css
	@npx esbuild web/assets/index.js --bundle --outdir=public/assets
	@go build -o $(BIN) .
	@echo "compiled you application with all its assets to a single binary => $(BIN)"
