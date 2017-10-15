FROM centos:7
RUN yum install -y build-essential golang git

ADD . /go/src/tastes/inject-libc

RUN export GOPATH=/go && \
	go get github.com/rainycape/dl && \
	cd /go/src/tastes/inject-libc && \
	go build -buildmode=c-shared -o /inject-libc.so libc_fs.go && \
	gcc -o /taste taste.c

RUN LD_PRELOAD=/inject-libc.so /taste
