FROM golang:latest

WORKDIR /go/src/github.com/jackale/simple-webapp

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql

CMD [ "go" , "run", "main.go" ]