# Dockerfile
FROM golang:1.23-alpine

# Dependências básicas
RUN apk add --no-cache gcc musl-dev git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY .env ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/api

# Porta default (usada no .env)
EXPOSE 8080

CMD ["./server"]
