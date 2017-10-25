FROM centos:7 AS builder

RUN yum install -y golang git build-essential make
RUN curl -fSL -o /usr/bin/dep \
	https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && \
	chmod +x /usr/bin/dep

ENV GOPATH=/go
ENV PATH=${PATH}:${GOPATH}/bin

COPY ./Gopkg.lock ./Gopkg.toml Makefile /go/src/github.com/setekhid/ketos/
RUN cd /go/src/github.com/setekhid/ketos && \
	(make vendor || echo "thanks to the Chinese meeting")

ARG VERSION=0.1.0

COPY . /go/src/github.com/setekhid/ketos
RUN cd /go/src/github.com/setekhid/ketos && make install

FROM centos:7
COPY --from=builder /usr/local/bin /usr/local/bin
COPY --from=builder /usr/local/lib /usr/local/lib
