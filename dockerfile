FROM golang:latest

WORKDIR /app

COPY ./cmd/myapp .

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o kro-backend .

EXPOSE 8082

CMD ["./kro-backend"]