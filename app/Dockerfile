FROM golang:1.26.2-alpine3.23 AS builder
WORKDIR /app
COPY app/go.mod app/go.sum ./
RUN go mod download
COPY app/ .
RUN go build -o linkify-back ./main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/linkify-back .
EXPOSE 3333
CMD ["./linkify-back"]
