KETOS_GO_EXEC ?= go
KETOS_DEP_EXEC ?= dep
VERSION ?= $(shell git describe --tags)

VERSION_PKG = github.com/setekhid/ketos/cmd/version
XCB_PKG = github.com/setekhid/ketos/cmd/xcb
HOOKROOT_PKG = github.com/setekhid/ketos/libcfs


build: xcb libketos-hookroot.so

test:
	go test -v -short ./...

install: build
	install -d ${DESTDIR}/usr/local/bin/
	install -m 755 ./xcb ${DESTDIR}/usr/local/bin/xcb
	install -d ${DESTDIR}/usr/local/lib/
	install -m 755 \
		./libketos-hookroot.so ${DESTDIR}/usr/local/lib/libketos-hookroot.so

xcb: $(shell find . -name '*.go')
	${KETOS_GO_EXEC} build \
		-ldflags "-X ${VERSION_PKG}.Version=${VERSION}" \
		-o xcb ${XCB_PKG}

libketos-hookroot.so: $(shell find . -name '*.go')
	${KETOS_GO_EXEC} build \
		-ldflags "-X ${HOOKROOT_PKG}.Version=${VERSION}" \
		-buildmode=c-shared \
		-o libketos-hookroot.so ${HOOKROOT_PKG}

clean:
	rm -rfv xcb libketos-hookroot.so

vendor: Gopkg.lock Gopkg.toml
	${KETOS_DEP_EXEC} ensure -vendor-only
