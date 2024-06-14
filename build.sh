#!/bin/bash -ex
tinygo build -target wasi -o count-vowels.wasm .
