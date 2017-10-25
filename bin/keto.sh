#!/bin/sh

export KETOS_ROOTPATH_LAYERS=/:/_ketos
export LD_PRELOAD=/usr/local/lib/libketos-hookroot.so

"$@"
