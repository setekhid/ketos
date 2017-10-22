#!/bin/bash

export KETOS_CHROOT_WD=FALSE
export KETOS_CHROOT_ROOT=/_ketos 
export LD_PRELOAD=/usr/local/lib/libketos-chroot.so

"$@"
