##################################
FROM golang

# Build Executable Binary

ADD . /go/src/github.com/CelesteComet/celeste-web-server
WORKDIR /go/src/github.com/CelesteComet/celeste-web-server

# Fetch Dependencies
RUN go get 

# Build Binary
RUN go build .

CMD ./celeste-web-server

EXPOSE 8080 

