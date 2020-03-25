wmf-api-gateway
===============

A Docker Compose configuration for testing/development of an API
gateway with rate-limiting.

You will need Docker Engine and Compose, then...

    $ make -C echoapi clean docker_image
    $ docker-compose up [--build]
    $ curl -v http://localhost:10000/foo/bar/baz  # In other terminal
