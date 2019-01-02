##################################
FROM golang

# Build Executable Binary

ADD . /go/src/github.com/CelesteComet/celeste-web-server
WORKDIR /go/src/github.com/CelesteComet/celeste-web-server

# Fetch Dependencies
RUN go get 

# Build Binary
RUN go build .

# Run server when container is run

CMD ./celeste-web-server

# Expose port 8080 of container

EXPOSE 8080 

