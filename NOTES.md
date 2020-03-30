ENVOY NOTES
===========

See: [T248543: Evaluate Envoy proxy for API gateway (and rate-limiter)][phab]


General (routing requests)
--------------------------

The routes we require will generally take the form
`api.wm.o/{something}/v{version}/{project}/{lang}/{path}` to
`{lang}.{project}.org/{...}`; We need to parse language and project
from the source URL, and use it to construct a destination hostname.
Envoy has no in-built mechanism for dealing with this.

Our test configuration uses the Lua HTTP filter to parse the URL,
string format the destination hostname, and inject a new HTTP header,
`x-internal-host`.  During routing, the `auto_host_rewrite_header` is
used to substitute the destination hostname with the value of
`x-internal-host` (see snipet below).

```yaml
- name: envoy.filters.network.http_connection_manager
  typed_config:
    "@type": type.googleapis.com/envoy.config.filter.network.http_connection_manager.v2.HttpConnectionManager
    stat_prefix: ingress_http
    route_config:
      name: local_route
      virtual_hosts:
      - name: local_service
        domains: ["*"]
        routes:
        - match:
            safe_regex:
              google_re2: {}
              regex: "^/core/v\\d{1}/wikipedia/en/.*$"
          route:
            regex_rewrite:
              pattern:
                google_re2: {}
                regex: ".*\/([^\/]+)\/wikipedia\/([^\/]+)(\/.*)$"
              substitution: /w/rest.php/\1\3
            auto_host_rewrite_header: "x-internal-host"
            cluster: service_echoapi
        - match:
            prefix: "/"
          route:
            host_rewrite: www.google.com
            cluster: service_echoapi
    http_filters:
    - name: envoy.filters.http.lua
      typed_config:
        "@type": type.googleapis.com/envoy.config.filter.http.lua.v2.Lua
        inline_code: |
          function envoy_on_request(request_handle)
            local path = request_handle:headers():get(":path")
            project, lang = string.match(path, "^/%a+/v%d/(%a+)/(%a+)/.*$")

            request_handle:headers():add("x-internal-host", lang .. "." ..project .. ".org")
          end
    - name: envoy.filters.http.router
```

This works, and despite seeming hacky, seems to be endorsed upstream.

It is (as of yet) unclear what the performance impact of Lua scripting
like this would be.

Any *fix* would entail coding changes (forking filters if the changes
could not be pushed upstream).


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


[phab]: https://phabricator.wikimedia.org/T248543
[rls.proto]: https://github.com/envoyproxy/envoy/blob/master/api/envoy/service/ratelimit/v3/rls.proto
[ratelimiter]: https://github.com/lyft/ratelimiter
