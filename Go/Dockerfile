FROM golang:1.13

MAINTAINER crase

WORKDIR /go/src/app

#port being used
EXPOSE 8080

COPY . .

RUN go get -u github.com/go-sql-driver/mysql
RUN go build main.go

CMD ["./main"]