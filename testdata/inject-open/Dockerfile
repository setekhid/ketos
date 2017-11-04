FROM centos:7
RUN yum install -y build-essential golang git

ADD . /go/src/tastes/inject-open

RUN export GOPATH=/go && \
	go get github.com/rainycape/dl && \
	cd /go/src/tastes/inject-open && \
	go build -buildmode=c-shared -o /inject-libc.so libc_fs.go

RUN echo abc > /abc && \
	LD_PRELOAD=/inject-libc.so cat /abc
