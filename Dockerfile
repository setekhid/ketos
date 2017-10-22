FROM centos:7 AS builder

RUN yum install -y golang git build-essential which

ENV GOPATH=/go
ENV PATH=${PATH}:${GOPATH}/bin

RUN mkdir -p /go/{src,bin,pkg} && \
	curl https://glide.sh/get | sh && \
	mkdir -p ~/.glide && \
	glide mirror set https://golang.org/x/mobile \
		https://github.com/golang/mobile --vcs git && \
	glide mirror set https://golang.org/x/crypto \
		https://github.com/golang/crypto --vcs git && \
	glide mirror set https://golang.org/x/net \
		https://github.com/golang/net --vcs git && \
	glide mirror set https://golang.org/x/tools \
		https://github.com/golang/tools --vcs git && \
	glide mirror set https://golang.org/x/text \
		https://github.com/golang/text --vcs git && \
	glide mirror set https://golang.org/x/image \
		https://github.com/golang/image --vcs git && \
	glide mirror set https://golang.org/x/sys \
		https://github.com/golang/sys --vcs git

COPY ./glide.lock ./glide.yaml /go/src/github.com/setekhid/ketos/
RUN cd /go/src/github.com/setekhid/ketos && \
	glide install && mkdir -p /opt/ketos

ARG VERSION=0.1.0

COPY . /go/src/github.com/setekhid/ketos
RUN cd /go/src/github.com/setekhid/ketos/libcfs && \
	go build -buildmode=c-shared -o /opt/ketos/libketos-hookroot.so . && \
	cd cmd/xcb && \
	go build \
		-ldflags "-X github.com/setekhid/ketos/cmd/version.Version=${VERSION}" \
		-o /opt/ketos/xcb \
		.

FROM centos:7
COPY --from=builder /opt/ketos /opt/ketos
