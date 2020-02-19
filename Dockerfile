FROM golang

WORKDIR /go/src/github.com/earthrockey/CICD-Golang
COPY . .

RUN go build -o main .
EXPOSE 8888
CMD ["./main"]