FROM golang

WORKDIR /go/src/github.com/earthrockey/CICD-go-angular/backend
COPY ./main .

EXPOSE 8080
CMD ["./main"]