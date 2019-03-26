BUILD_ID ?= 1
BUILD_SHA1 = $(shell git rev-parse --short=7 --verify HEAD)
BUILD_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
MAJOR ?= 1
MINOR ?= 0
ifeq ($(BUILD_BRANCH),master)
	PATCH = $(BUILD_ID)
else
	PATCH = $(BUILD_ID)-$(BUILD_SHA1)
endif

IMAGE_TAG := $(MAJOR).$(MINOR).$(PATCH)

ci: push
.PHONY: ci

clean:
	IMAGE_TAG=$(IMAGE_TAG) \
	BUILD_SHA1=$(BUILD_SHA1) \
    docker-compose --project-name app down || true
.PHONY: clean

build: clean
	IMAGE_TAG=$(IMAGE_TAG) \
	BUILD_SHA1=$(BUILD_SHA1) \
	docker-compose --project-name app up --detach
	E=$$(docker wait sut-$(IMAGE_TAG)) && exit $$E

	@echo successfully built docker image
.PHONY: build

push: build
	@docker login --username=${DOCKER_USERNAME} --password=${DOCKER_PASSWORD}
	IMAGE_TAG=$(IMAGE_TAG) \
	BUILD_SHA1=$(BUILD_SHA1) \
	docker-compose --project-name app push app
.PHONY: push