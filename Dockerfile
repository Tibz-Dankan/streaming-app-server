FROM golang:1.21-alpine

WORKDIR /app/server

COPY . .

RUN go build -o ./bin/myapp ./

EXPOSE 8080

CMD ["./bin/myapp"]
