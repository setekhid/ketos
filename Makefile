GOCMD=go
PRONAME=xcb

build:
	$(GOCMD) build -o $(PRONAME) github.com/setekhid/ketos/cmd/xcb
