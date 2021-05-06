FROM golang
MAINTAINER  idevlab
WORKDIR /go/src/
COPY . .
EXPOSE 80
CMD ["/bin/bash", "go build ."]