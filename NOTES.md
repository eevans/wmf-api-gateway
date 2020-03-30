NOTES
=====


General (routing requests)
--------------------------


Rate limiting
-------------

### What we want
Ideally, we parse the limit from an attribute of the JWT payload, and
pass that to the rate-limiter.

### What we have
Envoy has an HTTP filter for global rate limiting, it utilizes a rate
limit service with a generic gRPC interface (as defined by
[rls.proto].  The RPC is something like:

    ShouldRateLimit(RateLimitRequest) → RateLimitResponse

A `RateLimitRequest` has attributes `domain` (application
namespacing), `descriptors`, and `hits_addend` (number of requests to
count toward the limit; defaults to 1).  The `descriptors` attribute
is used to specify a list of identifiers, if *any* is over the limit,
the request should be rejected.

Obviously, the assumption that is made here is that limit
configuration is encapsulated in the rate limiter service, (and this
is in fact how the [reference implementation][ratelimiter] works).

### Ideas

##### Idea
Create a filter that validates a JWT signature, parses out a
configurable set of payload attributes, and adds them to dynamic
metadata.  Create a rate limiter filter that can retrieve a limit from
dynamic metadata, and pass it to its rate limiter service.

The new rate-limiter filter could be a (lightly) modified fork of the
existing one (or perhaps pushed upstream).  The gRPC interface could
also be a fork of the existing one (or changes proposed upstream).
Using the [reference rate limiter service][ratelimiter] as a starting
point might also be an option.


[rls.proto]: https://github.com/envoyproxy/envoy/blob/master/api/envoy/service/ratelimit/v3/rls.proto
[ratelimiter]: https://github.com/lyft/ratelimiter
