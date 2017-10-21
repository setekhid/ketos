GOCMD=go
PRONAME=xcb
IMPORTPATH="github.com/setekhid/ketos/cmd/version"
VERSION="0.0.0"
COMMIT_NO := $(shell git rev-parse HEAD 2> /dev/null || true)
COMMIT := $(if $(shell git status --porcelain --untracked-files=no),${COMMIT_NO}-dirty,${COMMIT_NO})


build:
	$(GOCMD) build -ldflags "-X $(IMPORTPATH).Version=$(VERSION) -X $(IMPORTPATH).Commit=$(COMMIT)" -o $(PRONAME) github.com/setekhid/ketos/cmd/xcb
