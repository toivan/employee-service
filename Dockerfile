FROM golang:1.21

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o employee-service ./cmd/app

EXPOSE 8080

CMD ["./employee-service"]

