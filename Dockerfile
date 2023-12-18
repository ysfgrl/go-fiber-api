FROM golang:1.18
WORKDIR /go/go-fiber-api
COPY . .
RUN go build
CMD ["./go-fiber-api"]