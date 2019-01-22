# image名称
fibonacci
给定一个整数N，返回长度为N的fibonacci数列

# image包含的环境变量
* bash变量
* fibonacci

# 使用示例
1. 写一个要处理的脚本命名为fibonacci_6.sh，内容为
```
fibonacci -fibNum 6 && \
echo  Still_water_run_deep
```
2. 配置fibonacci_6.sh投递到k8s需要的资源`fibonacci_6.sh.ini`
```
[templetes]
pod_name = fibonacci6
templete_name = fibonacci
env = idc_physical
module = xxx

[requests]
memory = 1
cpu = 1

[limits]
memory = 2
cpu = 2

[container]
image = registry.cn-beijing.aliyuncs.com/annoroad/fibonacci:v0.0.2
args = /home/zanyuan/bin/example/fibonacci_6.sh
[volumeMounts]
home = /cluster_home	store	/home
```
3. 使用`gomonitor submit -p /home/zanyuan/bin/example/fibonacci_6.sh.ini`命令投递到k8s，投递成功会得到podname，例如`fibonacci6-vc5h60v`
4. 使用`kubectl get pod podname`会得到此pod的状态显示,`kubectl get pod podname -o yaml`会得到详细的投递配置信息
5. `samtools_example.sh.o.podname`和`samtools_example.sh.e.podname`分别为标准输出流和标准错误流输出
6. 忘记了podname可通过.o或.e的后缀来确定

