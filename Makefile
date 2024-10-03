# Change these variables as necessary.
main_package_path = ./main.go
binary_name = croppy

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

## build: build the application
.PHONY: build
build:
	go build -o=./bin/${binary_name} ${main_package_path}
