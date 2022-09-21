FROM golang:1.18-alpine3.16
WORKDIR /opt/
COPY . .
RUN go build -o server cmd/main.go
CMD ["./server"]
