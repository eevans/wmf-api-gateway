

echoapi:
	go build echoapi.go

clean:
	rm -f echoapi

docker_image:
	docker build --tag echoapi .

.PHONY:
	clean docker_image echoapi
