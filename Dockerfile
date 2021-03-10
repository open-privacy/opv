######################################
# Prepare go_builder
######################################
FROM golang:1.16 as go_builder
WORKDIR /go/src/github.com/open-privacy/opv
ADD . .
RUN make build

######################################
# Copy from builder to alpine image
######################################
FROM frolvlad/alpine-glibc:alpine-3.10
RUN apk add --no-cache curl
VOLUME ["/data"]


ENV OPV_DB_DRIVER=sqlite3
ENV OPV_HOST=0.0.0.0
ENV OPV_DB_CONNECTION_STR=/data/_opv.sqlite?cache=shared&_fk=1

COPY --from=go_builder /go/src/github.com/open-privacy/opv/build/dataplane /usr/local/bin/dataplane
COPY --from=go_builder /go/src/github.com/open-privacy/opv/build/controlplane /usr/local/bin/controlplane

EXPOSE 27999
EXPOSE 28000

CMD dataplane

