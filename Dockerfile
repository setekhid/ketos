FROM centos:7
RUN yum install -y golang build-essential && \
	mkdir -p /go/src

ADD . /go/src/github.com/setekhid/ketos

RUN export GOPATH=/go && \
	cd $GOPATH/src/github.com/setekhid/ketos/libcfs && \
	go build -buildmode=c-shared -o /libketos-chroot.so *.go

VOLUME /data
CMD ["cp", "/libketos-chroot.so", "/data/"]
