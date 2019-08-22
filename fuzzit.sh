#!/bin/bash
set -xe

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <fuzz-type>"
    exit 1
fi

NAME=gojay
TYPE=$1

function fuzz {
    TARGET=$NAME-$1
    FUNCTION=Fuzz$2
    go-fuzz-build -libfuzzer -func $FUNCTION -o fuzzer.a .
    clang -fsanitize=fuzzer fuzzer.a -o fuzzer
    ./fuzzit create job --type $TYPE $TARGET fuzzer
}

# Setup
export GO111MODULE="off"
go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build
dep ensure -v
wget -q -O fuzzit https://github.com/fuzzitdev/fuzzit/releases/download/v2.4.29/fuzzit_Linux_x86_64
chmod a+x fuzzit

# Fuzz
fuzz unmarshal Unmarshal
fuzz decode Decode
fuzz stream Stream
