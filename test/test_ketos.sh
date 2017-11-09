#!/bin/bash

mkdir /taste && cd /taste

xcb init -I alpine:3.6 .
xcb pull 3.6
touch taste
xcb commit --tag 3.6
