SHELL := /bin/bash

define USAGE
Message API

Commands:
	help				Shows this.
	start				Builds and starts the application in a docker container on port 8080
	stop 				Stops the application
endef

help:
	@echo "${USAGE}"

start:
	docker build -t message_api -f ./build/dev/Dockerfile .
	docker run -d --rm --name message_api -p 8080:8080 message_api:latest

stop:
	docker stop message_api