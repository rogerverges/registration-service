FROM golang:1.14

RUN go get "github.com/go-sql-driver/mysql"

RUN mkdir -p /dockerFiles/post

ADD . /dockerFiles/post

WORKDIR  /dockerFiles/post 

RUN go build -o post .
 
CMD ["/dockerFiles/post/post"]

