FROM centos:7
ENV GOPATH=/go
RUN mkdir -p $GOPATH && \
	yum install -y build-essential golang git libseccomp-devel

RUN yum install -y vim && \
	go get github.com/lizrice/strace-from-scratch
