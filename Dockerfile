FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o todo-list ./cmd/todo

FROM alpine as runner

COPY --from=builder /app/todo-list .
COPY config/config.yaml ./config/config.yaml

CMD ["./todo-list"]