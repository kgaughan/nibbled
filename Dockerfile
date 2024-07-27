FROM alpine:latest
COPY nibbled .
ENTRYPOINT ["/nibbled"]
