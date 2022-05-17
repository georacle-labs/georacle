# syntax=docker/dockerfile:1

FROM golang:1.18.2-alpine
RUN apk add --no-cache ca-certificates gcc git linux-headers musl-dev
WORKDIR /georacle

# build executable
COPY . ./
RUN go mod download && go mod verify
RUN go install ./cmd/...

EXPOSE 3141
ENTRYPOINT [ "georacle" ]
