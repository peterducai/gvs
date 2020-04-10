#!/bin/bash
rm -rf gvs
rm -rf ~/go/bin/gvs
go build -o gvs gvs.go
go install .
rm -rf gvs
sleep 1
gvs
