# imagetools
提供管理k8s image dockerfile的工具

# 使用方法
程序的command如下：

```
Commands:

  add     Build and push an image to registry.cn-beijing.aliyuncs.com/annoroad
  list    Images list
  build   rebuild and push all images
  init    init the imagetools
  delete  Delete a image version

Arguments:

  -h  --help  Print help information
```

## 查询image仓库内都有哪些image可用
> 使用`imagetools list`命令，默认展示所有image的latest版本，一共包含3列信息

* image image名称
* version 相应的image版本
* latest 0 表示非latest版本，1 表示latest版本

```
imagetools list
image            	version   	latest
go_fqtools       	v0.0.2    	1
fibonacci        	v0.0.3    	1
```

* 默认只展示latest version，加上-a参数展示所有的image version
* 想展示某个image的所有version，可在-i参数之后加上相应的image名称

```
usage: imagetools list [-a|--showall] [-i|--imagename "<value>"] [-h|--help]

                  Images list

Arguments:

  -a  --showall    show all version image or not
  -i  --imagename  image name
  -h  --help       Print help information
```

## 向工具提交新的image dockerfile
> 使用`imagetools add -d /abspath/imagename`命令

`-d`参数是一个绝对路径，例如某个image名称为fibonacci，这个绝对路径的目录结构如下：

```
fibonacci/
├── dockerfiles
│   ├── context
│   │   ├── Dockerfile_fibonacci
│   │   └── fibonacci
│   ├── fibonacci.go
│   └── Makefile
└── images
    ├── AUTHORS
    ├── CONTRIBUTORS
    ├── READE.md
    └── VERSION
```

### image目录
* AUTHORS 包含image dockerfile的作者信息
`Yuan Zan <zanyuan@genome.cn>`
* CONTRIBUTORS 包含image dockerfile的除作者外其他版本贡献者信息
```
v0.0.2 Fan Xuning <xuningfan@genome.cn>
v0.0.3 Han Shaohuai <shaohuaihan@genome.cn>
```
* VERSION `image的版本号tag`，以三段式命名，不能包含其他内容，`非常重要`，示例:`v0.0.2`
* READE.md 包含此image使用说明，请务必包含以下内容
### READE.md规范

以下是一个READE.md的示例
``` 
# image名称
samtools-0.1.19
# image包含的环境变量
* bash变量
* samtools
# 使用示例
1. 写一个要处理的脚本命名为samtools_example.sh，内容为`samtools view -h test.bam`
2. 配置samtools_example.sh投递到k8s需要的资源`samtools_example.sh.ini`
3. 使用`gomonitor submit -p /abspath/samtools_example.sh.ini`命令投递到k8s，具体samtools_example.sh.ini里的配置形式参见`gomonitor submit -h`，投递成功会得到podname
4. 使用`kubectl get pod podname`会得到此pod的状态显示,`kubectl get pod podname -o yaml`会得到详细的投递配置信息
5. `samtools_example.sh.o.podname`和`samtools_example.sh.e.podname`分别为标准输出流和标准错误流输出
6. 忘记了podname可通过.o或.e的后缀来确定
```
### dockerfiles 目录
>dockerfile要按照`Dockerfile_`+`软件名称`+`软件版本`的形式命名，并且放在dockerfiles/context目录下

### Dockerfile的编写
1. 基础镜像请使用`registry.cn-beijing.aliyuncs.com/annoroad/annogene-base:v0.1`，这个镜像是修改过的alpine，定制了权限控制等，而且基于alpine对缩减镜像大小帮助很大
2. 要求镜像一定要具有bash普通的命令，主要包括：make、wget、zip、bash（这些一般在`/bin`能找到其链接，所以ENV`/bin`是必要的）
3. 能够直接使用软件变量（所以ENV软件的安装路径是必要的）
4. 尽量减少RUN的层数，RUN的层数越多，越不利于镜像的加载

以下是一个编写示例
```
FROM registry.cn-beijing.aliyuncs.com/annoroad/annogene-base:v0.1

MAINTAINER Zan Yuan <seqyuan@gmail.com>
ENV LANG=en_US.UTF-8

COPY ./fibonacci /opt/
WORKDIR /opt
RUN apk add bash make wget bzip2 ca-certificates && \
	 chmod +x fibonacci


ENV PATH /opt:$PATH:/bin
```
### dockerfiles/Makefile编写
以下是一个dockerfiles/Makefile示例

```
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
```

需要改动的有个地方
1. dockerfile
2. imagename
3. build模块 (如果涉及image的build过程的不止有Dockerfile)
4. test模块

# k8s常用的命令
1. `kubectl get pods --show-all`显示所有投递到k8s的pod，类似于qsat
2. `kubectl delete pod podname`从k8s删除一个pod
3. `sudo docker run -it $(dockerHub)/$(imagename):$(TAG) /bin/bash` 用于在docker测试节点，检查image推送前build是否有问题

