FROM golang:latest

WORKDIR /app

COPY ./cmd/myapp .

copy go.mod go.sum .

COPY ./internal ./internal

RUN go build -o kro-backend .

EXPOSE 8060

CMD ["./kro-backend"]
