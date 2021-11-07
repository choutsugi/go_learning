## 一、Docker安装

### 环境准备

腾讯云主机实例系统：Ubuntu20.04（腾讯云默认不支持root登录）

新建用户：

```bash
root@ubuntu:~# adduser okarin		# 添加用户：自动创建用户目录，使用默认shell
```

设置新建用户为sudoer：

```bash
root@ubuntu:~# chmod u+w /etc/sudoers
root@ubuntu:~# vim /etc/sudoers			# okarin  ALL=(ALL:ALL) ALL
root@ubuntu:~# chmod u-w /etc/sudoers
```

修改主机名：

```bash
root@ubuntu:~# vim /etc/hostname	# 修改主机名
root@ubuntu:~# reboot				# 重启生效
```

### Docker部署

配置阿里源：

```bash
okarin@LAB:~$ sudo curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
okarin@LAB:~$ sudo add-apt-repository "deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
```

安装Docker：

```bash
okarin@LAB:~$ sudo apt-get install docker-ce docker-ce-cli containerd.io -y
```

验证：

```bash
okarin@LAB:~$ docker -v
```

网卡：安装docker后，系统增加`docker0`网卡，地址为`172.17.0.1`

### Docker加速器

参考文档：

- http://guide.daocloud.io/dcs/daocloud-9153151.html
- https://www.daocloud.io/mirror#accelerator-doc

快速配置（Ubuntu）：

```bash
okarin@LAB:~$ curl -sSL https://get.daocloud.io/daotools/set_mirror.sh | sh -s http://f1361db2.m.daocloud.io
okarin@LAB:~$ sudo systemctl restart docker.service	# 重启Docker生效
```

### Docker服务

docker服务相关命令：

```bash
okarin@LAB:~$ sudo systemctl start docker	# 开启docker服务
okarin@LAB:~$ sudo systemctl status docker	# 查看docker状态
okarin@LAB:~$ sudo systemctl restart docker	# 重启docker服务
okarin@LAB:~$ sudo systemctl stop docker	# 停止docker服务
```

删除docker：

```bash
okarin@LAB:~$ sudo apt purge docker-ce -y
okarin@LAB:~$ sudo rm -rf /etc/docker
okarin@LAB:~$ sudo rm -rf /var/lib/docker/
```

docker目录：

```bash
/etc/docker/ 		#docker的认证目录
/var/lib/docker/ 	#docker的应用目录
```

### Docker镜像使用权限问题

问题：docker使用sudo安装，普通用户不使用sudo将无法（无权限）使用。

解决方法：

```bash
#添加docker group（若不存在）
okarin@LAB:~$ sudo groupadd docker
#将用户加入该 group 内，退出并重新登录。
okarin@LAB:~$ sudo gpasswd -a ${USER} docker	# sudo gpasswd -a okarin docker
#重启 docker 服务
okarin@LAB:~$ systemctl restart docker
#切换当前会话到新 group 或者重启会话（必须）
okarin@LAB:~$ newgrp - docker
```

## 二、Docker命令

### 镜像管理

#### 搜索、查看、获取

搜索镜像：

```bash
okarin@LAB:~$ docker search nginx	# docker search [镜像名]
```

获取镜像：

```bash
okarin@LAB:~$ docker pull nginx		# docker pull [镜像名]
```

查看镜像：

```bash
okarin@LAB:~$ docker images nginx	# 查看本地指定镜像
okarin@LAB:~$ docker images -a		# 列出本地所有镜像
okarin@LAB:~$ docker image ls		# 列出本地所有镜像
```

#### 重命名、删除

镜像重命名：

```bash
# docker tag [旧的镜像名称]:[旧的镜像版本][新的镜像名称]:[新的镜像版本]
okarin@LAB:~$ docker tag nginx:latest panda-nginx:v1.0
```

删除镜像：

```bash
okarin@LAB:~$ docker rmi 87a94228f133	# 使用ID删除
okarin@LAB:~$ docker rmi nginx:latest	# 使用名称删除
# 如果一个ID对应多个名称，则应按照 名称:版本 的格式删除镜像。
# 强制删除参数：-f, --force
```

#### 导出、导入

导出镜像：

```bash
# 指定写入的文件名和路径参数：-o, --output filename
okarin@LAB:~$ docker save -o nginx.tar nginx
```

导入镜像：

```bash
okarin@LAB:~$ docker load < nginx.tar
# 或者
okarin@LAB:~$ docker load -i nginx.tar
```

#### 历史、创建

查看镜像历史：

```bash
okarin@LAB:~$ docker history nginx:latest
# 或者
okarin@LAB:~$ docker history 87a94228f133
```

根据模板创建镜像：

```bash
#登录系统模板镜像网站：https://download.openvz.org/template/precreated/
#下载镜像模板，如ubuntu-16.04-x86_64.tar.gz，地址为：
#https://download.openvz.org/template/precreated/ubuntu-16.04-x86_64.tar.gz
#命令格式：
cat 模板文件名.tar | docker import - [自定义镜像名]
#演示效果：
$ cat ubuntu-16.04-x86_64.tar.gz | docker import - ubuntu-mini
```

### 容器管理

容器，即镜像运行时的实例，Ubuntu下存储位置`/var/lib/docker/containers`。

#### 查看、创建、启动

查看容器列表：

```bash
okarin@LAB:~$ docker ps
```

创建待启动容器：

```bash
okarin@LAB:~$ docker create -it --name ubuntu-1 ubuntu ls -a
#说明：镜像后的参数为容器启动后需要在容器中执行的命令，如 ls -a
#参数：
#	-t：分配虚拟终端
#	-i：即使未连接，也保持STDIN打开
#	--name：容器名称
```

启动容器：

```bash
okarin@LAB:~$ docker start -a ubuntu-1
#参数：
#	-a：将当前shell的STDOUT/STDERR连接到容器
#	-i：将当前shell的STDIN连接到容器
```

创建新容器并启动：

```bash
okarin@LAB:~$ docker run --rm --name nginx1 nginx /bin/echo "hello nginx"
#参数：
#	--rm：当容器退出后，自动删除容器
```

以守护进程方式启动容器：

```bash
okarin@LAB:~$ docker run -d nginx
#参数：
#	-d：在后台运行容器并打印容器ID
```

#### （取消）暂停、重启

容器暂停：

```bash
okarin@LAB:~$ docker pause 87a94228f133
```

取消容器暂停：

```bash
okarin@LAB:~$ docker unpause 87a94228f133
```

重启容器：

```bash
okarin@LAB:~$ docker restart -t 20 87a94228f133
#参数：
#	-t：等待指定秒数后执行重启
```

#### 关闭、终止、删除

关闭容器：

```bash
okarin@LAB:~$ docker stop 87a94228f133
```

终止容器：

```bash
okarin@LAB:~$ docker kill 87a94228f133
```

删除容器：

```bash
okarin@LAB:~$ docker rm 87a94228f133
```

强制删除容器：

```bash
okarin@LAB:~$ docker rm -f 87a94228f133
```

批量删除容器：

```bash
okarin@LAB:~$ docker rm -f $(docker ps -a -q)
```

#### 进入、退出

创建并进入容器：

```bash
okarin@LAB:~$ docker run -it --name panda-nginx nginx /bin/bash
```

退出容器：

```bash
root@e64af37e7c87:/# exit
#或者：Ctrl+C
```

手工方式进入容器：

```bash
okarin@LAB:~$ docker exec -it e64af37e7c87 /bin/bash
```

生产方式进入容器（脚本）：

1. 安装nsenter工具：

   ```bash
   okarin@LAB:~$ sudo apt install util-linux -y
   ```

2. 编写脚本docker-enter.sh：

   ```bash
   #!/bin/sh
   if [ -e $(dirname "$0")/nsenter ]; then
       # with boot2docker, nsenter is not in the PATH but it is in the same folder
       NSENTER=$(dirname "$0")/nsenter
   else
       NSENTER=nsenter
   fi
   if [ -z "$1" ]; then
       echo "Usage: `basename "$0"` CONTAINER [COMMAND [ARG]...]"
       echo ""
       echo "Enters the Docker CONTAINER and executes the specified COMMAND."
       echo "If COMMAND is not specified, runs an interactive shell in CONTAINER."
   else
       PID=$(docker inspect --format "``.`State`.`Pid`" "$1")
       if [ -z "$PID" ]; then
           exit 1
       fi
       shift
       OPTS="--target $PID --mount --uts --ipc --net --pid --"
       if [ -z "$1" ]; then
           # No command given.
           # Use su to clear all host environment variables except for TERM,
           # initialize the environment variables HOME, SHELL, USER, LOGNAME, PATH,
           # and start a login shell.
           "$NSENTER" $OPTS su - root
       else
          # Use env to clear all host environment variables.
          "$NSENTER" $OPTS env --ignore-environment -- "$@"
       fi
   fi
   ```

3. 赋予脚本执行权限：

   ```bash
   okarin@LAB:~$ chmod +x docker-enter.sh
   ```

4. 使用脚本进入容器：

   ```bash
   okarin@LAB:~$ ./docker_in.sh 58cb379fd546		# 失败！ Sate、Pid未识别。
   ```

#### 基于容器创建镜像

基于容器创建镜像（方式一）：

```bash
okarin@LAB:~$ docker commit -m 'some info..' -a "panda" e64af37e7c87 nginx:v0.2
#参数：
#	-m：改动信息
#	-a：作者
```

基于容器创建镜像（方式二）：

```bash
okarin@LAB:~$ docker export e64af37e7c87 > nginx.tar
#export与save相比不会保存镜像历史信息。
```

#### 日志、信息、端口、重命名

查看容器运行日志：

```bash
okarin@LAB:~$ docker logs e64af37e7c87
```

查看容器详细信息：

```bash
okarin@LAB:~$ docker inspect e64af37e7c87
```

查看容器端口信息：

```bash
okarin@LAB:~$ docker port e64af37e7c87
```

容器重命名：

```bash
okarin@LAB:~$ docker rename e64af37e7c87 nginx-1
```

### 数据管理

#### 数据卷概念

**数据卷**

docker将宿主机的某个目录，映射到容器，作为数据存储的目录，即容器内数据直接映射到宿主机。

**数据卷作用**

- 容器数据持久化。
- 外部机器和容器间接通信。
- 容器之间数据交换。

**数据卷特点**

- 数据卷是宿主机中的一个目录或文件（推荐使用文件）。
- 当容器目录和数据卷目录绑定后，对于任意一方的修改另一方都会同步。
- 一个容器卷可以被多个容器挂载。
- 一个容器可以挂载多个容器卷。

#### 配置数据卷

创建容器时，使用-v参数设置数据卷。

```bash
docker run ... -v 宿主机目录（文件）:容器内目录（文件） ...
```

注：

- 目录必须绝对路径。
- 若目录不存在，则将自动创建。
- 一个容易可同时挂载多个数据卷。

示例：

```bash
# 创建容器c1并挂载数据卷
okarin@LAB:~$ docker run -it --name=c1 -v /home/okarin/docker-test/data:/root/data_container nginx /bin/bash
# 退出容器c1
root@5d76a41d5a78:/# exit
# 数据卷内创建文件
okarin@LAB:~/docker-test/data$ touch hello.txt
# 启动容器
okarin@LAB:~$ docker start  5d76a41d5a78
# 进入容器
okarin@LAB:~$ docker exec -it 5d76a41d5a78 /bin/bash
# 进入容器查看文件
root@5d76a41d5a78:~/data_container# ls
# hello.txt
# 写入内容并退出容器
root@5d76a41d5a78:~/data_container# echo hello_world > hello.txt 
# 宿主机查看
okarin@LAB:~/docker-test/data$ cat hello.txt
root@5d76a41d5a78:/# exit
# hello_world

# 创建容器c2容器挂载多个数据卷
okarin@LAB:~$ docker run -it --name=c2 -v ~/docker-test/data2:/root/data2 -v ~/docker-test/data3:/root/data3  nginx /bin/bash
root@f2569f0d7d75:/# cd root/
root@f2569f0d7d75:~# ls
# data2  data3

# 两个容器挂载一个数据卷实现通信
# 创建容器c3，挂载~/docker-test/data1目录
okarin@LAB:~/docker-test$ docker run -it --name c3 -v ~/docker-test/data1:/root/data1 nginx /bin/bash
# 创建容器c4，挂载~/docker-test/data1目录
okarin@LAB:~/docker-test$ docker run -it --name c4 -v ~/docker-test/data1:/root/data1 nginx /bin/bash
# c3在数据卷中创建文件
root@6bd30030a7eb:~/data1# echo shared > common.txt
# c4查看数据卷
root@11591b95245c:~/data1# ls
# common.txt
```

#### 数据卷容器概念

使用数据卷容器在多个容器间共享数据，并永久保存数据，必须：

- 创建数据卷容器（自身无需启动）。
- 其他容器挂载数据卷容器。

#### 配置数据卷容器

1. 创建启动c5数据卷容器，使用-v参数设置数据卷：

```bash
#参数-v：容器数据卷目录，不指定宿主机目录时将自动创建。
okarin@LAB:~$ docker run -it --name=c1 -v /volume nginx /bin/bash
```

2. 启动c6、c7容器，使用--volumes-from参数设置数据卷容器：

```bash
okarin@LAB:~$ docker run -it --name=c2 --volumes-from c1 nginx /bin/bash
okarin@LAB:~$ docker run -it --name=c3 --volumes-from c1 nginx /bin/bash
```

#### 数据备份

方案：

1. 创建一个挂载数据卷容器的容器。
2. 该容器挂载宿主机本地目录作为备份数据卷。
3. 将数据卷容器的内容备份到宿主机本地目录挂载的数据卷中。
4. 完成备份操作后销毁创建的容器。

实践：

```bash
#创建备份目录
okarin@LAB:~$ mkdir ./backup/
#创建备份容器
okarin@LAB:~$ docker run --rm --volumes-from f9e6fe1c5d5c -v /home/okarin/backup/:/backup/ nginx tar zcPf /backup/data.tar.gz /data
#验证
okarin@LAB:~$ ls ./backup
okarin@LAB:~$ zcat ./backup/data.tar.gz
```

#### 数据还原

待续。

## 三、Docker应用部署

### 部署MySQL

1. 搜索mysql

```bash
okarin@LAB:~$ docker search mysql
```

2. 拉取mysql

```bash
okarin@LAB:~$ docker pull mysql:latest
```

3. 宿主机创建mysql目录，存储mysql数据

```bash
okarin@LAB:~$ mkdir -p ~/docker/mysql
okarin@LAB:~$ cd docker/mysql/
okarin@LAB:~/docker/mysql$ mkdir conf
okarin@LAB:~/docker/mysql$ cd conf/
okarin@LAB:~/docker/mysql/conf$ vim my.cnf
# 内容如下
[mysqld]
user=mysql
character-set-server=utf8
default_authentication_plugin=mysql_native_password
secure_file_priv=/var/lib/mysql
expire_logs_days=7
sql_mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION
max_connections=1000

[client]
default-character-set=utf8

[mysql]
default-character-set=utf8
```

4. 创建容器，设置端口映射、目录映射（容器自启，且拥有root权限）

```bash
okarin@LAB:~/docker/mysql$ docker run -p 13306:3306 --name my-mysql --restart=always --privileged=true -v $PWD/conf:/etc/mysql -v $PWD/logs:/var/log/mysql -v $PWD/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql
```

5. 进入容器

```bash
okarin@LAB:~/docker/mysql$ docker exec -it my-mysql /bin/bash
```

6. 登入mysl

```bash
root@db44dd99089d:/# mysql -uroot -p123456
mysql> 
```

7. 使用Navicat连接

## 四、dockerfile

### 4.1 docker镜像原理

docker镜像由特殊的文件系统叠加而成，最底端使用bootfs（复用宿主机的bootfs），在bootfs之上叠加rootfs，在rootfs之上可叠加其他镜像文件。

> docker镜像的本质是一个分层文件系统。

### 4.2 镜像制作

#### 方式一：容器转为镜像

命令如下：非挂载的文件将保留。

```bash
docker commit 容器id 镜像名称:版本号
docker save -o 压缩文件名称 镜像名称:版本号
docker load -i 压缩文件名称
```

#### 方式二：dockerfile

通过dockerfile将容器打包为镜像。

### 4.3 dockerfile

#### dockerfile关键字

| 关键字      | 作用                     | 备注                                                         |
| ----------- | ------------------------ | ------------------------------------------------------------ |
| FROM        | 指定父镜像               | 指定dockerfile基于那个image构建                              |
| MAINTAINER  | 作者信息                 | 用来标明这个dockerfile谁写的                                 |
| LABEL       | 标签                     | 用来标明dockerfile的标签 可以使用Label代替Maintainer 最终都是在docker image基本信息中可以查看 |
| RUN         | 执行命令                 | 执行一段命令 默认是/bin/sh 格式: RUN command 或者 RUN ["command" , "param1","param2"] |
| CMD         | 容器启动命令             | 提供启动容器时候的默认命令 和ENTRYPOINT配合使用.格式 CMD command param1 param2 或者 CMD ["command" , "param1","param2"] |
| ENTRYPOINT  | 入口                     | 一般在制作一些执行就关闭的容器中会使用                       |
| COPY        | 复制文件                 | build的时候复制文件到image中                                 |
| ADD         | 添加文件                 | build的时候添加文件到image中 不仅仅局限于当前build上下文 可以来源于远程服务 |
| ENV         | 环境变量                 | 指定build时候的环境变量 可以在启动的容器的时候 通过-e覆盖 格式ENV name=value |
| ARG         | 构建参数                 | 构建参数 只在构建的时候使用的参数 如果有ENV 那么ENV的相同名字的值始终覆盖arg的参数 |
| VOLUME      | 定义外部可以挂载的数据卷 | 指定build的image那些目录可以启动的时候挂载到文件系统中 启动容器的时候使用 -v 绑定 格式 VOLUME ["目录"] |
| EXPOSE      | 暴露端口                 | 定义容器运行的时候监听的端口 启动容器的使用-p来绑定暴露端口 格式: EXPOSE 8080 或者 EXPOSE 8080/udp |
| WORKDIR     | 工作目录                 | 指定容器内部的工作目录 如果没有创建则自动创建 如果指定/ 使用的是绝对地址 如果不是/开头那么是在上一条workdir的路径的相对路径 |
| USER        | 指定执行用户             | 指定build或者启动的时候 用户 在RUN CMD ENTRYPONT执行的时候的用户 |
| HEALTHCHECK | 健康检查                 | 指定监测当前容器的健康监测的命令 基本上没用 因为很多时候 应用本身有健康监测机制 |
| ONBUILD     | 触发器                   | 当存在ONBUILD关键字的镜像作为基础镜像的时候 当执行FROM完成之后 会执行 ONBUILD的命令 但是不影响当前镜像 用处也不怎么大 |
| STOPSIGNAL  | 发送信号量到宿主机       | 该STOPSIGNAL指令设置将发送到容器的系统调用信号以退出。       |
| SHELL       | 指定执行脚本的shell      | 指定RUN CMD ENTRYPOINT 执行命令的时候 使用的shell            |

**示例：基于ubuntu的nginx。**

命令：

```bash
# 创建目录
okarin@LAB:~$ mkdir docker-files/dockerfile/nginx -p
okarin@LAB:~$ cd docker-files/dockerfile/nginx/
okarin@LAB:~/docker-files/dockerfile/nginx$ vim Dockerfile
```

Dockerfile：

```dockerfile
# 构建基于ubuntu的docker定制镜像
# 基础镜像
FROM ubuntu

# 镜像作者
MAINTAINER krain krain.lab@qq.com

# 执行命令
RUN mkdir hello && mkdir world
RUN sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
&& sed -i 's/security.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN apt-get update && apt-get install nginx -y

# 启动命令
CMD ["/user/sbin/nginx", "-g", "daemon off;"]

# 对外端口
EXPOSE 80
```

构建：

```bash
okarin@LAB:~/docker-files/dockerfile/nginx$ docker build -t ubuntu-nginx:v1.0 .
```

> 通过&&合并命令可减小镜像文件大小。

#### 基础指令

**FROM**

指定基础镜像。

**MAINTAINER**

指定作者及联系信息。

**RUN**

执行，官方推荐使用exec方式：`RUN ["mkdir", "hello"]`。

**EXPOSE**

对外端口。

#### 运行时指令

**CMD**

启动·命令：指定容器启动时默认执行的一条命令，指定多条时只执行最后一条。使用docker run时若指定了运行指令，则CMD将被覆盖。

如：nginx指定后台启动方式关闭。

```dockerfile
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
```

**ENTRYPOINT**

入口，不能被覆盖。

```dockerfile
ENTRYPOINT ["/usr/sbin/nginx", "-g", "daemon off;"]
```

**CMD ENTRYPOINT**

可以同时使用CMD和ENTRYPOINT。

```dockerfile
ENTRYPOINT ["/usr/sbin/nginx"]
CMD ["-g"]
```

固定命令使用ENTRYPOINT，可变命令使用CMD。

**ADD**

将宿主机的文件拷贝到容器文件系统指定目录。如果文件是可识别的压缩格式，docker将自动解压；ADD多用于解压文件。

- 如果原路径是文件且以`/`结尾：创建目录，并拷贝文件。
- 如果原路径是文件且不以`/`结尾：目标路径作为文件。

- 如果原路径是目录：创建目录，并拷贝目录下所有文件。
- 如原路径为压缩包：自动解压。

示例：

```dockerfile
ADD ["source.list", "/etc/apt/source.list"]
ADD ["krain.tar.gz", "/hello/"]
```

**COPY**

COPY同ADD，区别在于COPY不会自动解压，COPY多用于拷贝文件。

**VOLUME**

在镜像中创建挂载点。

示例：

```dockerfile
VOLUME ["/var/lib/shared/"]
```

其他容器可以挂载该容器实现数据卷容器的功能。

**ENV**

设置环境变量，可以在RUN之前设置，RUN时调用；容器启动时环境变量被指定。

docker启动时设置环境变量：

```bash
okarin@LAB:~$ docker run -e NIHAO="hello" -itd --name ubuntu-test ubuntu /bin/bash
```

Dockerfile中设置环境变量：

```dockerfile
ENV NIHAO=hello
```

**WORKDIR**

配置docker工作目录。

```dockerfile
WORKDIR /root/workspace/test1
WORKDIR /root/workspace/test2
```

#### 触发器指令

当一个镜像A被另一个镜像B作为基础镜像构建时，将激活触发器并执行。

```dockerfile
ONBUILD COPY ["index.html", "/var/www/okarin/com/"]
```

#### 其他指令

**USER**

指定容器运行时的用户，默认使用root。

**ARG**

指定变量在docker build时使用。

### 4.4 实例：构建go环境

获取ubuntu模板文件。

```bash
cat ubuntu-16.04-x86_64.tar.gz | docker import - ubuntu-mini
```

启动docker容器。

```bash
docker run -itd --name go-test ubuntu-mini /bin/bash
```

进入容器。

```bash
docker exec -it go-test /bin/bash
```

配置国内源。

```bash
cp /etc/apt/soucre.list /etc/apt/soucre.list.old
vim /etc/apt/soucre.list
```

> 配置中科大镜像源。

```bash
sed -i 's/cn.archive.ubuntu.com/mirros.ustc.edu.cn/g' /etc/apt/sources.list
```

更新软件源，安装基本软件。

```bash
apt update
apt install gcc libc6-dev git vim lrzsz -y
```

安装go环境。

```bash
apt insatll golang -y
# 或
tar -c /usr/local -zxf go1.10.linux-amd64.tar.gz
export GOROOT=/usr/local/go
export PATN=$PATH:/usr/local/go/bin
export GOPATH=/root/go
export PATH=$GOPATH/bin/:$PATH
```

## 五、docker compose

任务编排：对多个子任务执行顺序进行确定的过程。

- 单机版：docker-compose

- 集群版：k8s

### 5.1 安装：

```bash
sudo apt install python3-pip -y
sudo pip install docker-compose
```

### 5.2 使用

compose配置文件：docker-compose.yml

```yaml
version: '2'							# compose版本号
service:								# 服务标识符
  web1:									# 子服务名
    image: nginx						# 服务依赖镜像属性
	  ports:							# 服务端口属性
        - "9999:80"						# 宿主机端口:容器端口
      container_name: nginx-web1		# 容器命名
  web2:
    image: nginx
      ports:
        - "8888:80"
      container_name: nginx-web1
```

执行：

```bash
# -d 后台运行
docker-compose up -d
```

查看：

```bash
docker-compose ps
```

停止且删除

```bash
docker-compose down
```

### 5.3 命令

#### compose服务启动、关闭、查看

```bash
#后台启动
docker-compose up -d
#删除服务（整体删除）
docker-compose down
```

#### 容器开启、关闭、删除

```bash
#启动服务
docker-compose start <服务名>
#停止服务
docker-compose stop <服务名>
#删除容器（不会删除网络和数据卷，慎用）
docker-compose rm <服务名> 
```

#### 其他信息

```bash
#查看正在运行的服务
docker-compose ps
#查看日志，使用-f参数持续追踪服务产生的日志。
docker-compose logs -f
#查看服务依赖的镜像
docker-compose images
#进入服务容器
docker-compose exec <服务名> <执行命令>
#查看服务网络
docker network ls
```

### 5.4 docker-compose详解

参见：https://blog.csdn.net/qq_36148847/article/details/79427878

官方示例：

```yaml
version: "3"
services:

  redis:
    image: redis:alpine
    ports:
      - "6379"
    networks:
      - frontend
    deploy:
      replicas: 2
      update_config:
        parallelism: 2
        delay: 10s
      restart_policy:
        condition: on-failure

  db:
    image: postgres:9.4
    volumes:
      - db-data:/var/lib/postgresql/data
    networks:
      - backend
    deploy:
      placement:
        constraints: [node.role == manager]

  vote:
    image: dockersamples/examplevotingapp_vote:before
    ports:
      - 5000:80
    networks:
      - frontend
    depends_on:
      - redis
    deploy:
      replicas: 2
      update_config:
        parallelism: 2
      restart_policy:
        condition: on-failure

  result:
    image: dockersamples/examplevotingapp_result:before
    ports:
      - 5001:80
    networks:
      - backend
    depends_on:
      - db
    deploy:
      replicas: 1
      update_config:
        parallelism: 2
        delay: 10s
      restart_policy:
        condition: on-failure

  worker:
    image: dockersamples/examplevotingapp_worker
    networks:
      - frontend
      - backend
    deploy:
      mode: replicated
      replicas: 1
      labels: [APP=VOTING]
      restart_policy:
        condition: on-failure
        delay: 10s
        max_attempts: 3
        window: 120s
      placement:
        constraints: [node.role == manager]

  visualizer:
    image: dockersamples/visualizer:stable
    ports:
      - "8080:8080"
    stop_grace_period: 1m30s
    volumes:
      - "/var/run/docker.sock:/var/run/docker.sock"
    deploy:
      placement:
        constraints: [node.role == manager]

networks:
  frontend:
  backend:

volumes:
  db-data:
```

## 六、docker私有仓库

参见：https://www.bilibili.com/video/BV1CJ411T7BK?p=25

待续。



































