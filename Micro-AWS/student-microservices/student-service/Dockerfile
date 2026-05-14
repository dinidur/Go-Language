FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o student-service ./student-service/main.go

EXPOSE 9090

CMD ["./student-service"]