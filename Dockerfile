FROM envoyproxy/envoy-dev:93dd9459b380bdb7c4a61d8e8c14fdb88580669c
COPY envoy.yaml /etc/envoy/envoy.yaml
COPY jwks.json /etc/envoy/jwks.json
