FROM centos:7
RUN yum install -y golang build-essential make && \
	mkdir -p /go/src

ADD . /go/src/github.com/setekhid/ketos
ADD ./bin/keto.sh /

RUN export GOPATH=/go && \
	( \
		cd $GOPATH/src/github.com/setekhid/ketos/libcfs && \
		go build -buildmode=c-shared -o /usr/local/lib/libketos-chroot.so *.go \
	) && \
	( \
		cd $GOPATH/src/github.com/setekhid/ketos && \
		make build && make install \
	)
