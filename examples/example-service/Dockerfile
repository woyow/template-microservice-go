FROM golang:1.19.1-alpine as builder
RUN apk add --update gcc musl-dev

WORKDIR /app
COPY go.* /app/
COPY /internal/ /app/internal
COPY /config/ /app/config
RUN go mod tidy
RUN go mod download

COPY /cmd /app/cmd
RUN CGO_ENABLED=0 go build ./cmd/example-service/main.go

#FROM alpine:3.15
FROM scratch
COPY --from=builder /app/main /app/
WORKDIR /app

COPY .env .env
COPY /db /app/db
COPY /config /app/config

CMD ["./main"]