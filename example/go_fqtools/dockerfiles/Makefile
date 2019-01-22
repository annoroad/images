BASEDIR=$(shell echo `dirname $(abspath $(lastword $(MAKEFILE_LIST)))`)
contextDir=$(BASEDIR)/context
dockerfile=Dockerfile_go_fqtools
imagename=go_fqtools

VERSIONFILE = $(BASEDIR)/../images/VERSION
TAG=`cat $(VERSIONFILE)`

dockerHub=registry.cn-beijing.aliyuncs.com/annoroad

build:
	cd $(BASEDIR) && go build $(BASEDIR)/go_fqtools.go
	mv $(BASEDIR)/go_fqtools $(contextDir)
	sudo docker build -t $(dockerHub)/$(imagename):$(TAG) -f $(contextDir)/$(dockerfile) $(contextDir)

test:
	gomonitor submitpod $(BASEDIR)/..//go_fqtools_test.ini

push:
	sudo docker push $(dockerHub)/$(imagename):$(TAG)

pull:
	sudo docker pull $(dockerHub)/$(imagename):$(TAG)

help:
	@echo 'Makefile for $(imagename) docker image                       '
	@echo '                                                             '
	@echo 'Usage:                                                       '
	@echo '   make build         build $(imagename) docker image        '
	@echo '   make test          test the image                         '
	@echo '   make push          push the image to Docker hub registry  '
	@echo '   make pull          pull the image from Docker registry    '


