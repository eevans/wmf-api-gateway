FROM envoyproxy/envoy-dev:b49d117cde32518565ae74a070fd0a63304eadaf
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml --service-cluster proxy
COPY jwks.json /etc/envoy/jwks.json
