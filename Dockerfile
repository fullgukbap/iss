FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main ./cmd/main/main.go

FROM alpine:latest AS final

WORKDIR /app

COPY --from=builder /app/bin/main ./

# Docker 8080 노출만 시킨다. 
EXPOSE 8080

ENTRYPOINT [ "./main" ]
