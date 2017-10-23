FROM centos:7 AS builder

RUN yum install -y golang git build-essential which wget && \
	wget -O /usr/bin/dep \
	https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && \
	chmod +x /usr/bin/dep

ENV GOPATH=/go
ENV PATH=${PATH}:${GOPATH}/bin

COPY ./Gopkg.lock ./Gopkg.toml /go/src/github.com/setekhid/ketos/
RUN cd /go/src/github.com/setekhid/ketos && \
	(dep ensure -vendor-only || true) && mkdir -p /opt/ketos

ARG VERSION=0.1.0

COPY . /go/src/github.com/setekhid/ketos
RUN cd /go/src/github.com/setekhid/ketos && \
	go build \
		-buildmode=c-shared -o /opt/ketos/libketos-hookroot.so \
		github.com/setekhid/ketos/libcfs && \
	go build \
		-ldflags "-X github.com/setekhid/ketos/cmd/version.Version=${VERSION}" \
		-o /opt/ketos/xcb \
		github.com/setekhid/ketos/cmd/xcb

FROM centos:7
COPY --from=builder /opt/ketos /opt/ketos
