#!/bin/sh
export GOPATH=$(pwd)
go get
go run 14_excercise_slices.go |
    tr ':' ' ' |
    awk '{print $2}' |
    base64 -D > exercise_slices.png
open exercise_slices.png
