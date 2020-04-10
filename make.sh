#!/bin/bash

go build -o gvs gvs.go
go install .
rm -rf gvs
sleep 1
gvs
