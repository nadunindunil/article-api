FROM golang:1.19
WORKDIR /go/src/article-api
COPY . .
RUN go build -o bin/server main.go
CMD ["./bin/server"]