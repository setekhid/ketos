FROM centos:7
RUN yum install -y build-essential golang git

ADD . /go/src/tastes/inject-libc

RUN export GOPATH=/go && \
	go get github.com/rainycape/dl && \
	cd /go/src/tastes/inject-libc && \
	(cd taste_exec && go build -o /exec) && \
	go build -buildmode=c-shared -o /inject-libc.so libc_fs.go && \
	gcc -o /taste taste.c

RUN /exec
RUN LD_PRELOAD=/inject-libc.so /taste
