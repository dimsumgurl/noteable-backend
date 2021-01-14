FROM golang:1.15

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o bin/noteable cmd/noteable/main.go 

# Run the binary program produced by `go install`
CMD ["./bin/noteable"]