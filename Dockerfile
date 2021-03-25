######################################
# Prepare go_builder
######################################
FROM golang:1.16 as go_builder
WORKDIR /data/opv
ADD . .
RUN make build

######################################
# Copy from builder to alpine image
######################################
FROM frolvlad/alpine-glibc:alpine-3.10
RUN apk add --no-cache curl
VOLUME ["/data"]

ENV OPV_HOST=0.0.0.0
ENV OPV_DB_DRIVER=sqlite3
ENV OPV_DB_CONNECTION_STR=/data/_opv.sqlite?cache=shared&_fk=1
ENV OPV_PROXY_PLANE_HTTP_CONFIG=/opv-proxyplane-http.example.json

COPY --from=go_builder /data/opv/build/dataplane    /usr/local/bin/dataplane
COPY --from=go_builder /data/opv/build/controlplane /usr/local/bin/controlplane
COPY --from=go_builder /data/opv/build/proxyplane   /usr/local/bin/proxyplane

# Default example of the configuration json for the proxy plane
# Can be overridden by OPV_PROXY_PLANE_HTTP_CONFIG
COPY --from=go_builder /data/opv/cmd/proxyplane/opv-proxyplane-http.example.json /opv-proxyplane-http.example.json

EXPOSE 27999
EXPOSE 28000
EXPOSE 28001

CMD dataplane
