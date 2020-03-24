wmf-api-gateway
===============

You will need Docker Engine and Compose, then...

    $ make -C echoapi clean docker_image
    $ docker-compose up
    $ curl -v http://localhost:10000/foo/bar/baz  # In other terminal
