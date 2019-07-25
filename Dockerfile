FROM golang:1.11-alpine3.10 as builder

RUN apk update && apk add --no-cache \
    coreutils \
    git \
    make \
    openssh-client

WORKDIR /go/src/github.com/streaming-user

COPY . .
COPY ./schemas /schemas

RUN make build

FROM alpine:3.10

RUN apk update && apk add --no-cache \
    ca-certificates

COPY --from=builder /go/src/github.com/streaming-user/bin/linux_amd64/strm-user /usr/bin
COPY --from=builder /go/src/github.com/streaming-user/schemas /schemas/
CMD ["/usr/bin/strm-user"]
EXPOSE 8080