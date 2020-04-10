#!/bin/bash
rm -rf gvs
go build -o gvs gvs.go
go install .
rm -rf gvs
sleep 1
gvs
