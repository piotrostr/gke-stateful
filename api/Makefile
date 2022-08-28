image_tag = piotrostr/stateful-api
binary_name = stateful-api

.PHONY: all

all: test build docker

test:
	go test ./...

build:
	go build -o $(stateful_api) .

docker:
	docker build \
		--tag $(image_tag) \
		--platform linux/amd64 .
	docker push $(image_tag)
