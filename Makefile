GOCMD=go
DOCKERCMD=docker

ifndef ($(GOPATH))
	GOPATH = $(HOME)/go
endif

all: build docker-build

build:
	$(GOCMD)  get -d
	$(GOCMD)  build -o bin/ovs-exporter .

docker-build:
	$(DOCKERCMD) build -t ovs-exporter:0.27 .
