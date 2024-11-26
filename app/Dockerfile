FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-logging-app ./app/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/go-logging-app .

EXPOSE 8080

CMD ["./go-logging-app"]