# Build
FROM golang:1.21.11-alpine AS builder

COPY . /src
WORKDIR /src

RUN apk --no-cache --no-progress add build-base git bash

RUN make build

# Run
FROM alpine:latest

ARG TIMEZONE
ENV TIMEZONE=${TIMEZONE:-"Asia/Shanghai"}

RUN apk update \
    && apk --no-cache add \
        bash \
        ca-certificates \
        curl \
        tzdata \
    && ln -sf /usr/share/zoneinfo/${TIMEZONE} /etc/localtime \
    && echo "${TIMEZONE}" > /etc/timezone

WORKDIR /app

COPY --from=builder /src/bin/nglab ./nglab
COPY --from=builder /src/configs ./configs
COPY --from=builder /src/scripts/entrypoint.sh ./entrypoint.sh

RUN chmod 755 ./nglab
RUN chmod 755 ./entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]