FROM golang:1.23.0 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM golang:1.23.0

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["/app/main"]
