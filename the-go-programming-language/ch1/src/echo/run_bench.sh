#!/bin/bash
go test -bench . | grep -v asdf
