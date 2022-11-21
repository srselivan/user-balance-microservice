FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

EXPOSE 8000

RUN go build -o bin/microservice ./cmd/main.go

CMD ["bin/microservice"]