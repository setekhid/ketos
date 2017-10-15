FROM centos:7
RUN yum install -y build-essential golang

ADD . /go/src/tastes/shared-library

RUN cd /go/src/tastes/shared-library && \
	go build -buildmode=c-shared -o /nautilus.so nautilus.go && \
	gcc -o /wale wale.c /nautilus.so
