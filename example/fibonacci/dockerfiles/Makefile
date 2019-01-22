BASEDIR=$(shell echo `dirname $(abspath $(lastword $(MAKEFILE_LIST)))`)
contextDir=$(BASEDIR)/context
dockerfile=Dockerfile_fibonacci
imagename=fibonacci
VERSIONFILE = $(BASEDIR)/../images/VERSION
TAG=`cat $(VERSIONFILE)`

#dockerHub=registry.cn-qingdao.aliyuncs.com/seqyuan
#dockerHub=registry-vpc.cn-qingdao.aliyuncs.com/seqyuan

dockerHub=registry.cn-beijing.aliyuncs.com/annoroad

build:
	cd $(BASEDIR) && go build $(BASEDIR)/fibonacci.go
	mv $(BASEDIR)/fibonacci $(contextDir)
	sudo docker build -t $(dockerHub)/$(imagename):$(TAG) -f $(contextDir)/$(dockerfile) $(contextDir)

test:
	sudo docker run $(dockerHub)/$(imagename):$(TAG) fibonacci -fibNum 10

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


