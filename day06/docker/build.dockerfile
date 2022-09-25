# step 1: build executable binary
FROM golang:1.18-alpine AS builder
LABEL maintainer="Irham Sahbana <irhamsahbana@gmail.com>"
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN cd ./app && go build -o /clean-arch

# step 2: build a small image
FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder clean-arch .
COPY ./config.yaml .
EXPOSE 9099
CMD ["./clean-arch"]
# CMD ["./clean-architecture-go", "-m=migrate"]