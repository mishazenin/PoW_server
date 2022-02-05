# Build.
FROM golang:1.16 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/main.go

# Run.
FROM alpine:3.12.0 AS launcher

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .

CMD /root/server
