# Build stage
FROM golang:1.20-alpine3.16 AS builder
ENV CGO_ENABLED=0

WORKDIR /app
COPY . .

RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY config.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration
EXPOSE 8080

CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
