#FROM debian:wheezy
FROM ubuntu:14.04
MAINTAINER Tomokazu Kiyohara "tomokazu.kiyohara@gmail.com"

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH $PATH:$GOPATH/bin:$GOROOT/bin
ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update
RUN apt-get install -y git
RUN apt-get install -y mercurial
RUN apt-get install -y wget

WORKDIR /tmp
RUN wget --quiet https://godeb.s3.amazonaws.com/godeb-amd64.tar.gz
RUN tar xvf godeb-amd64.tar.gz
RUN ./godeb install 1.3.3
RUN rm -rf godeb-amd64.tar.gz godeb


RUN go get git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git

RUN mkdir -p /opt/mqtt-client
WORKDIR /opt/mqtt-client

ADD pub.go /opt/mqtt-client/pub.go
RUN go build /opt/mqtt-client/pub.go

ADD sub.go /opt/mqtt-client/sub.go
RUN go build /opt/mqtt-client/sub.go

CMD [ "/opt/mqtt-client/sub" ]
