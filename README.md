wmf-api-gateway
===============

A Docker Compose configuration for testing/development of an API
gateway with rate-limiting.

Prerequisites
-------------

You will need Docker Engine, and Docker compose installed.

You will need a Docker image for the ratelimit service:

    $ git submodule init  # First time
    $ git submodule update --remote  # To pull new changes
    $ cd ratelimit
    $ docker build --tag ratelimit .
    ...
    $ cd -

You will need a Docker image for Hydra:

    $ git submodule init  # First time
    $ git submodule update --remote  # To pull new changes
    $ cd hydra
    $ docker-compose -f quickstart.yml  -f quickstart-postgres.yml   -f quickstart-jwt.yml  up --build
    ...
 In another Terminal for Hydra:

    $ docker-compose -f quickstart.yml exec hydra\
      hydra clients create \
      --endpoint http://127.0.0.1:4445/ \
      --id my-client \
      --secret secret \
      -g client_credentials
You will need a Docker image for the echoapi service:

    $ make -C echoapi clean docker_image
    

Running
-------

    $ docker-compose <up|down> [--build]

*NOTE: If you edit `envoy.yaml`, use `--build` on the next `up` to
re-create the envoy Docker image and copy the file in.*

In another terminal:

    $ curl -v http://localhost:10000/foo/bar/baz  # In other terminal
    $ curl -v http://localhost:10000/core/v5/wikipedia/en/foo/bar/baz
    $ curl -v -H "Authorization: Bearer <JWT>" http://localhost:10000/core/v5/wikipedia/en/foo/bar/baz
    $ curl -v -H "Authorization: Bearer incorrect_JWT" http://localhost:10000/core/v5/wikipedia/en/foo/bar/baz  # 401

To get JWT (only copy access_token from output):

    $ curl -s -k -X POST -H "Content-Type: application/x-www-form-urlencoded" \
      -d grant_type=client_credentials -u 'my-client:secret' http://localhost:4444/oauth2/token

To get JWTK for jwks.json:

    $ curl -s -k -X POST -H "Content-Type: application/x-www-form-urlencoded" \
      -d grant_type=client_credentials -u 'my-client:secret' http://localhost:4444/oauth2/token
