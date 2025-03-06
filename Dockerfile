FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o todo-list

FROM golang:1.23

WORKDIR /app

COPY --from=builder /app/todo-list .
COPY .env .

EXPOSE 8080
CMD ["./todo-list"]
