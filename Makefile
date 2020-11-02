DOCKER_IMAGE_VER=0.3

DOCKER_CONTAINER=devplugin:${DOCKER_IMAGE_VER}

.PHONY: build
build:
	CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' ./cmd/dp

.PHONY: docker-build
docker-build:
	docker build -t ${DOCKER_CONTAINER} .

.PHONY: docker-push
docker-push:
	docker tag ${DOCKER_CONTAINER} docker.io/dmitsh/${DOCKER_CONTAINER} && docker push docker.io/dmitsh/${DOCKER_CONTAINER}
