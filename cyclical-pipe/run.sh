#!/bin/sh

cleanup () {
    rm -rf bin
}

mkdir bin
go build -o bin/child github.com/jasonkeene/playground/cyclical-pipe/child
trap cleanup EXIT

CHILD=bin/child go run github.com/jasonkeene/playground/cyclical-pipe/parent
