FROM centos:7

RUN yum install -y golang git build-essential make && \
	( \
		curl -fSL -o /usr/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && \
		chmod +x /usr/bin/dep || \
		echo "thanks to Professor Binxing Fang" \
	)

ENV GOPATH=/go

COPY ./Gopkg.lock ./Gopkg.toml Makefile $GOPATH/src/github.com/setekhid/ketos/
RUN cd $GOPATH/src/github.com/setekhid/ketos && \
	(make vendor || echo "thanks to the Chinese meeting")

ARG VERSION=1.0.0

COPY . /go/src/github.com/setekhid/ketos
RUN cd /go/src/github.com/setekhid/ketos && make install
