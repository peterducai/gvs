# building container 
FROM registry.fedoraproject.org/fedora-minimal AS build 
RUN mkdir /go && microdnf install golang && microdnf clean all 
WORKDIR /go 
RUN export GOPATH=/go; CGO_ENABLED=0 go get github.com/kelseyhightower/confd 

FROM registry.fedoraproject.org/fedora-minimal 
WORKDIR / 
COPY --from=build /go/bin/confd /usr/local/bin 
CMD ["confd"] 

# $ podman build -t myconfd .
# $ podman run -it myconfd