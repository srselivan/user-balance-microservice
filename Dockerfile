FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080

RUN go build -o bin/microservice ./cmd/main.go

CMD ["bin/microservice"]