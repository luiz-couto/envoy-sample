FROM lyft/envoy:latest

COPY /edge-envoy/envoy.json /etc/envoy.json
CMD /usr/local/bin/envoy -c /etc/envoy.json