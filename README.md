wmf-api-gateway
===============

A Docker Compose configuration for testing/development of an API
gateway with rate-limiting.

Prerequisites
-------------

You will need Docker Engine, and Docker compose installed.

### Rate limit service

You will need a Docker image for the ratelimit service:

    $ git submodule init  # First time
    $ git submodule update --remote  # To pull new changes
    $ cd ratelimit
    $ docker build --tag ratelimit .
    ...
    $ cd -

### Echoapi

You will need a Docker image for the echoapi service:

    $ make -C echoapi clean docker_image
    
### JWT generator

You will need to build the `jwt` binary (requires a Golang build
environment):

    $ go get github.com/square/go-jose
    $ make -C jwt
    

Running
-------

    $ docker-compose <up|down> [--build]

*NOTE: If you edit `envoy.yaml`, use `--build` on the next `up` to
re-create the envoy Docker image and copy the file in.*

In another terminal:

    $ curl -v http://localhost:10000/foo/bar/baz  # In other terminal
    $ curl -v http://localhost:10000/core/v5/wikipedia/en/foo/bar/baz
    $ WEB_TOKEN=`jwt/jwt jwt/keys/jwk-sig-FafStFaO5aapFjOjHhz9cWifF5pr17Ymi5dskSi6QP0=-priv.json`
    $ curl -v -H "Authorization: Bearer $WEB_TOKEN" http://localhost:10000/core/v5/wikipedia/en/foo/bar/baz
    $ curl -v -H "Authorization: Bearer incorrect_JWT" http://localhost:10000/core/v5/wikipedia/en/foo/bar/baz  # 401
