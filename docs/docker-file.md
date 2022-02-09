# Dockerfile

## FROM-指定基础镜像
指定基础镜像，并且Dockerfile中第一条指令必须是FROM指令，且在同一个Dockerfile中创建多个镜像时，可以使用多个FROM指令。

语法格式如下：

```Dockerfile
FROM <image> 
FROM <image>:<tag> 
```

其中`tag`是可选项，如果没有选择，那么默认值为latest。

如果不以任何镜像为基础，那么写法为：FROM scratch。


## RUN-运行指定的命令
运行指定的命令。

包含两种语法格式，如下所示：

### shell格式：就像在命令行中输入的Shell脚本命令一样。
RUN <command> 

### exec格式：就像是函数调用的格式。
RUN ["executable", "param1", "param2"] 
第一种后边直接跟shell命令。

在linux操作系统上默认为/bin/sh -c
在windows操作系统上默认为cmd /S /C
第二种是类似于函数调用。可将executable理解成为可执行文件，后面就是两个参数。

样例：
```Dockerfile
RUN /bin/bash -c 'source $HOME/.bashrc; echo $HOME'

RUN ["/bin/bash", "-c", "echo hello"] 
```
注意：
多行命令不要写多个RUN，原因是Dockerfile中每一个指令都会建立一层。多少个RUN就构建了多少层镜像，会造成镜像的臃肿、多层，不仅仅增加了构件部署的时间，还容易出错。RUN书写时的换行符是\


# WORKDIR

WORKDIR指令为Dockerfile中的任何RUN、CMD、ENTRYPOINT、COPY和ADD指令设置工作目录。如果WORKDIR不存在，它将被创建，即使它没有在任何后续Dockerfile指令中使用。

语法 :

WORKDIR dirpath
WORKDIR指令可以在Dockerfile中多次使用。如果提供了一个相对路径，它将相对于前一个WORKDIR指令的路径。例如:

```Dockerfile
# 这个Dockerfile中最后一个pwd命令的输出将是/a/b/c。
WORKDIR /a
WORKDIR b
WORKDIR c
RUN pwd
```

WORKDIR指令可以解析之前使用ENV设置的环境变量。只能使用在Dockerfile中显式设置的环境变量。例如:
```Dockerfile
# 这个Dockerfile中最后一个pwd命令的输出是/path/$DIRNAME
ENV DIRPATH=/path
WORKDIR $DIRPATH/$DIRNAME
RUN pwd
```