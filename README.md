wmf-api-gateway
===============

A Docker Compose configuration for testing/development of an API
gateway with rate-limiting.

Prerequisites
-------------

You will need Docker Engine, and Docker compose installed.

You will need a Docker image for the ratelimit service:

    $ cd ratelimit
    $ docker build --tag ratelimit .
    ...
    $ cd -
    
You will need a Docker image for the echoapi service:

    $ make -C echoapi clean docker_image
    

Running the environment

    $ docker-compose up [--build]
    $ curl -v http://localhost:10000/foo/bar/baz  # In other terminal


