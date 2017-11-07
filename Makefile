KETOS_GO_EXEC ?= go
KETOS_DEP_EXEC ?= dep
VERSION ?= $(shell git describe --tags)

VERSION_PKG = github.com/setekhid/ketos/cmd/version
XCB_PKG = github.com/setekhid/ketos/cmd/xcb
CHR_PKG = github.com/setekhid/ketos/cmd/chr
HOOKROOT_PKG = github.com/setekhid/ketos/libcfs


build: xcb chr libketos-hookroot.so

test:
	go test -short -race $(shell go list ./... | grep -v /vendor/)

install: build test
	install -d ${DESTDIR}/usr/local/bin/
	install -m 755 ./xcb ${DESTDIR}/usr/local/bin/xcb
	install -m 755 ./bin/keto.sh ${DESTDIR}/usr/local/bin/keto.sh
	install -d ${DESTDIR}/usr/local/lib/
	install -m 755 \
		./libketos-hookroot.so ${DESTDIR}/usr/local/lib/libketos-hookroot.so

xcb: $(shell find . -name '*.go')
	${KETOS_GO_EXEC} build \
		-ldflags "-X ${VERSION_PKG}.Version=${VERSION}" \
		-o xcb ${XCB_PKG}

chr: $(shell find . -name '*.go')
	${KETOS_GO_EXEC} build \
		-o chr ${CHR_PKG}

libketos-hookroot.so: $(shell find . -name '*.go')
	${KETOS_GO_EXEC} build \
		-ldflags "-X ${HOOKROOT_PKG}.ketos_libcfs_version=${VERSION}" \
		-buildmode=c-shared \
		-o libketos-hookroot.so ${HOOKROOT_PKG}

clean:
	rm -rfv xcb chr libketos-hookroot.so

vendor: Gopkg.lock Gopkg.toml
	${KETOS_DEP_EXEC} ensure -vendor-only

.PHONY: build test install clean
