FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY . .

RUN go mod tidy
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
