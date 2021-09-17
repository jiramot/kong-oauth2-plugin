.PHONY: build

build:
	go build -o bin/open-api cmd/main.go

docker:
	docker build -t ghcr.io/jiramot/kong-oauth2-plugin:latest . --no-cache

push:
	docker push ghcr.io/jiramot/kong-oauth2-plugin
