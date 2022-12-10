dir = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
.PHONY: help

help: ## help command for available tasks
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help


build: ## build the container
	docker build -t pizza .

build-nc: ## build the container w/o a cache
	docker build --no-cache -t pizza .

run: ## run the container with default parameters
	docker run --net=host -v $(dir)/src/config/:/config pizza

up: build run ## build the container and boot

image:
	docker tag pizza docker.pkg.github.com/adityaxdiwakar/pizza/pizza:latest
	
push-image:
	docker push docker.pkg.github.com/adityaxdiwakar/pizza/pizza:latest
