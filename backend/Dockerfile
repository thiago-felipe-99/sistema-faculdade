FROM golang:1.18-alpine

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go get -d -v ./...
RUN go install -v ./...

COPY . .
RUN mkdir -p logs/data

RUN go build -o main .

CMD ["./main"]
