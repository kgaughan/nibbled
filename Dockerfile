FROM golang:1.16 AS builder

ENV GOPATH /go
ENV APPPATH /repo
COPY . /repo
RUN cd /repo && CGO_ENABLED=0 go build -tags netgo -trimpath -ldflags '-s -w' -o nibbled .

FROM alpine:latest
COPY --from=builder /repo/nibbled /nibbled
ENTRYPOINT ["/nibbled"]
