# Build
FROM golang:1.21.11-alpine AS builder

COPY . /src
WORKDIR /src

RUN apk --no-cache --no-progress add build-base git bash

RUN make build

# Run
FROM alpine:latest

WORKDIR /app

COPY --from=builder /src/bin/nglab ./app/nglab

EXPOSE 8080

ENTRYPOINT ["./nglab"]