FROM golang:1.21-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go build -o house ./main.go

CMD ["./house"]