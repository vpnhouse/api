IMAGE_NAME := oapi-codegen-image
CONTAINER_NAME := oapi-codegen-container
DOCKERFILE_DIR := ./oapi-codegen

.PHONY: build
build:
	docker build -t $(IMAGE_NAME) $(DOCKERFILE_DIR)

.PHONY: gen
gen:
	docker run --rm --name $(CONTAINER_NAME) -v $(PWD):/app -w /app/oapi-codegen $(IMAGE_NAME) ./update.sh

.PHONY: all
all: build gen