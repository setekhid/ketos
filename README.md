> **Notice**:
>
> For now, everything should be worked theoretically, but some dependencies may broken caused by GFW and Chinese meeting. Continuous developing is not possible. Before that, I will complete refactoring my graduation design first.

# Cross Container Builder (Ketos)

[![GoDoc](https://godoc.org/github.com/setekhid/ketos?status.svg)](https://godoc.org/github.com/setekhid/ketos) [![Go Report Card](https://goreportcard.com/badge/github.com/setekhid/ketos)](https://goreportcard.com/report/github.com/setekhid/ketos) [![Build Status](https://travis-ci.org/setekhid/ketos.svg?branch=master)](https://travis-ci.org/setekhid/ketos)

A tasting project for Go Hackathon 2017 Shanghai.

This project aims to help you building a docker image in CI platform better.

## Hackathon Team (alpha order)

```
Ace-Tang   (github.com/ace-tang) <aceapril@126.com>
Huitse Tai (github.com/setekhid) <geb.1989@gmail.com>
```

## Building & Usage

```bash
docker build -t setekhid/ketos:latest .
docker run -it --rm  setekhid/ketos:latest xcb help
```
