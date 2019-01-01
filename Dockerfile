FROM golang:latest

ADD . /go/src/github.com/CelesteComet/celeste-web-server
WORKDIR /go/src/github.com/CelesteComet/celeste-web-server
RUN go get 
RUN go install
ENTRYPOINT /go/bin/celeste-web-server
EXPOSE 8080

