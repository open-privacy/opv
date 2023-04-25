######################################
# Prepare go_builder
######################################
FROM golang:1.20 as go_builder
WORKDIR /data/opv
ADD . .
RUN make build

######################################
# Copy from builder to alpine image
######################################
FROM frolvlad/alpine-glibc:alpine-3.10
RUN apk add --no-cache curl
VOLUME ["/data"]

COPY --from=go_builder /data/opv/build/dataplane    /usr/local/bin/dataplane
COPY --from=go_builder /data/opv/build/controlplane /usr/local/bin/controlplane
COPY --from=go_builder /data/opv/build/proxyplane   /usr/local/bin/proxyplane

EXPOSE 27999
EXPOSE 28000
EXPOSE 28001

CMD dataplane
