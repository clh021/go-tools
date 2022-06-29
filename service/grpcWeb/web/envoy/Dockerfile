# docker build -t tmp-envoy .
FROM envoyproxy/envoy:v1.9.0
COPY envoy.docker.yaml /etc/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy.yaml
