# Cross Container Builder (xcbuild)

A project for Hackathon 2017 at Shanghai.

Utilities to seperate the whole `docker build` and provide subcommands to be better integrated with CI platform.

## Prototype

* Implement the functions RUN, ADD/COPY and ENV from original dockerfile commands.
* Integrate with docker registry, push and pull down docker images.

## Authors

* Ace-Tang   github.com/ace-tang <aceapril@126.com>
* Huitse Tai github.com/setekhid <geb.1989@gmail.com>

## Installation

```bash
go get github.com/setekhid/ketos/cmd/xcb
```
