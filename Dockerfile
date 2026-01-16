FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/movies.json .
CMD ["./app"]
