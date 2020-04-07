jwt
===

Dependencies
------------

    $ go get github.com/square/go-jose
    $ go get github.com/square/go-jose/jose-util


Building
--------

    $ make


Crafting a signed JWT
---------------------

    $ TOKEN=`./jwt keys/jwk-sig-FafStFaO5aapFjOjHhz9cWifF5pr17Ymi5dskSi6QP0=-priv.json`
    $ curl -v -H "Authorization: bearer $TOKEN" http://localhost:10000/core/v5/wikipedia/en/foo/bar/baz


Contents of `keys/`
-------------------

| File | Description |
| ---- | ---- |
| jwk-sig-FafStFaO5aapFjOjHhz9cWifF5pr17Ymi5dskSi6QP0=-priv.json | RSA256 JWK for signing (private). Generated w/ `jose-util generate-key --use sig --alg RS256` |
| jwk-sig-FafStFaO5aapFjOjHhz9cWifF5pr17Ymi5dskSi6QP0=-pub.json | RSA256 JWK for signing (public). Generated w/ `jose-util generate-key --use sig --alg RS256` |
| jwks.json | JSON web key set (hand-crafted from the generated private/public JWKs) |
