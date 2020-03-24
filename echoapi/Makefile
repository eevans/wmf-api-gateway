

echoapi:
	go build echoapi.go

clean:
	rm -f echoapi

docker_image:
	docker build --tag echoapi .

docker_run:
	docker run --rm --name echoapi echoapi

.PHONY:
	clean docker_image docker_run echoapi
