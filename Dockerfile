FROM debian:buster

USER root

ENV DEBIAN_FRONTEND noninteractive
ENV LC_ALL C.UTF-8
ENV LANG C.UTF-8

RUN set -ex; \
    apt-get update; \
    apt-get install -y build-essential \
                       golang-go

RUN groupadd -r echoapi && useradd -r -g echoapi echoapi

WORKDIR /app

COPY echoapi.go Makefile ./

RUN make

EXPOSE 8080

USER echoapi

CMD ["./echoapi"]
