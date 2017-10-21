GOCMD=go
PRONAME=xcb
IMPORTPATH="github.com/setekhid/ketos/cmd/version"
VERSION="0.0.0"


build:
	$(GOCMD) build -ldflags "-X $(IMPORTPATH).Version=$(VERSION)" -o $(PRONAME) github.com/setekhid/ketos/cmd/xcb
