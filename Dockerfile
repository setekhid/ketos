FROM centos:7
RUN yum install -y golang build-essential make && \
	mkdir -p /go/src

ADD . /go/src/github.com/setekhid/ketos

RUN export GOPATH=/go && \
	( \
		cd $GOPATH/src/github.com/setekhid/ketos/libcfs && \
		go build -buildmode=c-shared -o /libketos-chroot.so *.go \
	) && \
	( \
		cd $GOPATH/src/github.com/setekhid/ketos && \
		make \
	)
