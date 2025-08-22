build_dir = ./build
binary_name = main

.PHONY: help
help:
	@echo "Usage":
	@sed -n 's/^##//p' ${MAKEFILE_LIST} |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

## dev: Enter Docker Go development
.PHONY: dev
dev:
	docker run --rm -it -v .:/app -p 8080:8080 --name go-tutorial -w /app  golang:1.22

## build: Build the application
.PHONY: build
build:
	@mkdir -p ${build_dir} && chown 1000:1000 ${build_dir}
	GOARCH=amd64 GOOS=linux go build -o ${build_dir}/${binary_name} -tags prod main.go

## clean: Clean-up the build binaries
.PHONY: clean
clean: confirm
	@echo "Cleaning up..."
	@rm -rf ${build_dir}

## run: Run the server
.PHONY: run
run:
	@go run -tags prod main.go
