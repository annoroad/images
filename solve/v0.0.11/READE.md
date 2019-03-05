# image名称
Solve3.3
# image包含的环境变量
* python2.7.5
* perl5.16
* R3.1.2
* R包data.table
* R包igraph
* R包MASS
* R包intervals
* R包XML
# 使用示例
/opt/xxx/pipelineCL.py

参数

-T 每个节点的线程；
-j 每个节点的最大线程 ；
-jp MAXTHREADSPW Max Threads per pairwise or stage0 job [default -T arg]
-N 切割的bnx个数；
-G gap区的bed文件；检测SV使用；
-i 迭代的个数；设置0-20，0代表跳过；
-b 输入的bnx文件；
-l 输出的根目录，不存在会创建，存在会覆盖。
-t 工具软件的目录，Location of executable files (RefAligner and Assembler,required)
-B 跳的步骤，使用之前的结果。0:None, 1:ImgDetect, 2:NoiseChar/Subsample, 3:Pairwise, 4:Assembly, 5:RefineA, 6:RefineB, 7:merge0
-e 输出前缀；
-r 参考序列；.cmap格式；
-x 不组装；
-c Remove contig results(0 - keep all (default), 1 - remove intermediate files, 2 - store in sqlite, 3 - store in sqlite and remove)
