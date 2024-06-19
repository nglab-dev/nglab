# Build
FROM golang:1.21.11-alpine AS builder

COPY . /src
WORKDIR /src

RUN apk --no-cache --no-progress add build-base git bash nodejs npm

RUN make install && make build

# Run
FROM alpine:latest

COPY --from=builder /src/bin/nglab /app/nglab
COPY --from=builder /src/configs /app/configs
COPY --from=builder /src/scripts/entrypoint.sh /app/entrypoint.sh

WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["./entrypoint.sh"]