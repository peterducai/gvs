#!/bin/bash

ctr1=$(buildah from "registry.fedoraproject.org/fedora-minimal:30")

# FROM registry.fedoraproject.org/fedora-minimal:30 
# RUN microdnf install nodejs && microdnf clean all

## Get all updates and install ruby
buildah run "$ctr1" -- microdnf update -y
buildah run "$ctr1" -- microdnf install -y golang
buildah run "$ctr1" -- microdnf clean all
buildah run "$ctr1" -- go get github.com/peterducai/gvs
#buildah run "$ctr1" -- mkdir -p /workdir

## Include some buildtime annotations
buildah config --annotation "com.example.build.host=$(uname -n)" "$ctr1"

## Commit this container to an image name
buildah commit "$ctr1" "${2:-$USER/gvs}"