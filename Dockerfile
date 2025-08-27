FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o ./build/release/renamer ./cmd/renamer


FROM alpine
RUN apk add --no-cache tzdata alpine-conf && \
    setup-timezone -z Asia/Shanghai && \
    apk del alpine-conf && \
    rm -rf /var/cache/apk/* && \
    mkdir -p /app

WORKDIR /app
COPY --from=builder /app/build/release/renamer /app/renamer
ENTRYPOINT ["/app/renamer"]
