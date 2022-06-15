#!/bin/bash
docker run -it --name envoy --rm -v "$(pwd)/envoy.yaml":/etc/envoy.yaml -p 18080:18080 envoyproxy/envoy:v1.9.0 \
/usr/local/bin/envoy -c /etc/envoy.yaml