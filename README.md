wmf-api-gateway
===============

echoapi
-------

A REST API that simply echos back request data as JSON


To build a Docker image:

    $ make docker_image

To run `echoapi` in the foreground:

    $ docker run --rm --name echoapi echoapi

