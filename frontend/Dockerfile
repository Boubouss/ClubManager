FROM golang:1.25-alpine AS builder

WORKDIR /app/src

COPY frontend . 

COPY shared /shared

RUN go mod download

RUN go build -o /clubmanager /app/src/cmd/main.go

FROM alpine:latest 

COPY --from=builder /clubmanager .

EXPOSE 8080

ENTRYPOINT ["/clubmanager"]
