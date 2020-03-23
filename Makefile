
echoapi:
	go build echoapi.go

clean:
	rm -f echoapi

.PHONY:
	echoapi
