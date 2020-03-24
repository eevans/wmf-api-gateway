
docker_image:
	docker build -t envoy:test .

docker_run:
	docker run --rm -d --name envoy -p 9901:9901 -p 10000:10000 envoy:test

.PHONY:
	docker_image docker_run
