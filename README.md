wmf-api-gateway
===============

envoy
-----

    $ docker pull envoyproxy/envoy-dev:93dd9459b380bdb7c4a61d8e8c14fdb88580669c
    $ docker run --rm -d -p 10000:10000 envoyproxy/envoy-dev:93dd9459b380bdb7c4a61d8e8c14fdb88580669c
    $ curl -v localhost:10000


echoapi
-------

A REST API that simply echos back request data as JSON


To build a Docker image:

    $ cd echoapi/
    $ make docker_image

To run `echoapi` in the foreground:

    $ docker run --rm --name echoapi echoapi

