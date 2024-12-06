.PHONY: build
build:
	docker build --tag bfs-image .
	-rm Dockerfile

.PHONY: push-dockerhub
push-dockerhub:
	docker tag bfs-image mensur/bfs-image
	docker push mensur/bfs-image

	docker tag bfs-image mensur/bfs-image:abc-123
	docker push mensur/bfs-image:abc-123

.PHONY: push-github-packages
push-github-packages:
	docker tag bfs-image ghcr.io/mensurui/bfs-image
	docker push ghcr.io/mensurui/bfs-image

	docker tag bfs-image ghcr.io/mensurui/bfs-image:abc-123
	docker push ghcr.io/mensurui/bfs-image:abc-123
