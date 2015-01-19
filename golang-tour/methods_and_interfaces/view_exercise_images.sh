#!/bin/sh
export GOPATH=$(pwd)
go get
go run 16_exercise_images.go |
    tr ':' ' ' |
    awk '{print $2}' |
    base64 -D > exercise_images.png
open exercise_images.png
