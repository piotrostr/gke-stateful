binary_name = stateful-api
docker_image = piotrostr/$(binary_name)
project_id = piotrostr-resources
tag = $(shell git rev-parse HEAD)

.PHONY:

all: test build docker submit

test:
	go test ./...

build:
	go build -o $(binary_name) .

docker:
	docker build \
		--tag $(docker_image):$(tag) \
		--platform linux/amd64 \
		.
	docker push $(docker_image):$(tag)

submit:
	bash check.sh
	gcloud builds submit \
		--tag gcr.io/$(project_id)/$(binary_name):$(tag) .
