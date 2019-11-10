Docker 
=============

#### 配置镜像加速器
1. 配置加速器
```
###  /etc/default/docker
DOCKER_OPTS="--registry-mirror=https://xxxx.mirror.aliyuncs.com"
```
2. 重启docker服务
```shell
sudo service docker restart
```
-------------------------------------------------------------------------------------

获取镜像：
```shell
docker pull [选项] [Docker Registry地址]<仓库名>:<标签>
```
* Docker Registry地址: 地址格式一般是 <域名/IP>[:端口]。默认是Docker HUB。
* 仓库名: 这里的仓库名是两段式名称，即 <用户名>/<软件名> 。如果不给出用户名，则默认为library  ，也就是官方镜
像。

列出镜像
```
docker images
```


#### 运行
```shell
docker run -it --rm ubuntu:14.04 bash
```
* -it  ：这是两个参数，一个是  -i  ：交互式操作，一个是  -t  终端
* --rm  ：这个参数是说容器退出后随之将其删除。默认情况下，为了排障需求，退出的容器并不会立即删除，除非手动  docker rm 

#### 这类无标签镜像也被称为 虚悬镜像(dangling image)
使用以下命令专门显示dangling image
```shell
docker images -f dangling=true
```
一般来说，虚悬镜像已经失去了存在的价值，是可以随意删除的，可以用下面的命令删除
```
docker rmi $(docker images -q -f dangling=true)
```

#### 中间层镜像
默认的  docker images  列表中只会显示顶层镜像，如果希望显示包括中间层镜像在内的所有镜像的话，需要加  -a  参数。
```
docker images -a
```
镜像是多层存储，每一层是在前一层的基础上进行的修改；而容器同样也是多层存储，是在以镜像为基础层，在其基础上加一层作为容器运行时的存储层。
```shell
docker run --name webserver -d -p 80:80 nginx
```
nginx  镜像启动一个容器，命名为  webserver  ，并且映射了 80
端口，这样我们可以用浏览器去访问这个  nginx  服务器


docker exec 命令进入容器
```shell
docker exec -it webserver bash
```
以交互式终端方式进入  webserver  容器，并执行了  bash  命令，也就是获得一个可操作的 Shell

docker diff 查看容器具体改动
```shell
docker diff webserver
```
要知道，当我们运行一个容器的时候（如果不使用卷的话），我们做的任何文件修改都会被记录于容器存储层里。而 Docker 提供了一个  docker commit  命令，可以将容器的存储层保存下来成为镜像。换句话说，就是在原有镜像的基础上，再叠加上容器的存储层，并构成新的镜像。以后我们运行这个新镜像的时候，就会拥有原有容器最后的文件变化。

docker commit(慎用)  的语法格式为：
```shell
docker commit [选项] <容器ID或容器名> [<仓库名>[:<标签>]]
```
docker history  具体查看镜像内的历史记录
```shell
docker history nginx:v2
```
#### 使用 Dockerfile 定制镜像
定制每一层所添加的配置、文件。如果我们可以把每一层修改、安装、构建、操作的命令都写入一个脚本，用这个脚本来构建、定制镜像，那么之前提及的无法重复的问题、镜像构建透明性的问题、体积的问题就都会解决。这个脚本就是Dockerfile。
##### Dockerfile
Dockerfile 是一个文本文件，其内包含了一条条的指令(Instruction)，每一条指令构建一层，因此每一条指令的内容，就是描述该层应当如何构建。

创建Dockerfile文件
```dockerfile
FROM nginx
RUN echo '<h1>Hello, Docker!</h1>' > /usr/share/nginx/html/index
.html
```
* FROM 指定基础镜像
一个Dockerfile中FROM是必备的指令，并且必须是第一条指令。
在Docker Hub上有非常多高质量的官方镜像，可以在其中寻找一个最符合我们最终目标的镜像为基础镜像进行定制。
* 除了选择现有镜像为基础镜像外，Docker 还存在一个特殊的镜像，名为scratch  。这个镜像是虚拟的概念，并不实际存在，它表示一个空白的镜像。
```dockerfile
FROM scratch
```
如果你以scratch为基础镜像的话，意味着你不以任何镜像为基础，接下来所写的指令将作为镜像第一层开始存在。
* RUN 执行命令
RUN  指令是用来执行命令行命令的。RUN  指令在定制镜像时是最常用的指令之一。其格式有两种：
1. shell 格式： RUN <命令> 就像直接在命令行中输入的命令一样
```dockerfile
RUN echo '<h1>Hello, Docker!</h1>' > /usr/share/nginx/html/index
.html
```
2. exec 格式 ： RUN ["可执行文件", "参数1", "参数2"]  ，这更像是函数调用中的格式。

Dockerfile 中每一个指令都会建立一层， RUN  也不例外。
Union FS 是有最大层数限制的，比如 AUFS，曾经是最大不得超过 42 层，现在是不得超过 127 层。
例子：
```dockerfile
FROM debian:jessie
RUN buildDeps='gcc libc6-dev make' \
&& apt-get update \
&& apt-get install -y $buildDeps \
&& wget -O redis.tar.gz "http://download.redis.io/releases/r
edis-3.2.5.tar.gz" \
&& mkdir -p /usr/src/redis \
&& tar -xzf redis.tar.gz -C /usr/src/redis --strip-component
s=1 \
&& make -C /usr/src/redis \
&& make -C /usr/src/redis install \
&& rm -rf /var/lib/apt/lists/* \
&& rm redis.tar.gz \
&& rm -r /usr/src/redis \
&& apt-get purge -y --auto-remove $buildDeps
```
使用&&将各个所需命令串联起来，简化为了 1 层。
Dockerfile支持Shell类的行尾添加 \ 的命令换行方式，以及行首#进行注释的格式。良好的格式，比如换行、缩进、注释等，会让维护、排障更为容易，这是一个比较好的习惯。 
一定要确保每一层只添加真正需要添加的东西，任何无关的东西都应该清理掉。
很多人初学 Docker 制作出了很臃肿的镜像的原因之一，就是忘记了每一层构建的最后一定要清理掉无关文件。

### 构建镜像
在  Dockerfile  文件所在目录执行：
```shell
docker build -t nginx:v3 .
```
上面的点是指定上下文路径
镜像构建上下文(Context)
这就引入了上下文的概念。当构建的时候，用户会指定构建镜像上下文的路径， docker build  命令得知这个路径后，会将路径下的所有内容打包，然后上传给Docker引擎。

一般来说，应该会将  Dockerfile  置于一个空目录下，或者项目根目录下。如果该目录下没有所需文件，那么应该把所需文件复制一份过来
.dockerignore： 
可以用.gitignore一样的语法写一个.dockerignore，该文件是用于剔除不需要作为上下文传递给 Docker引擎的。

### 其它docker build的用法
#### docker build  还支持从URL构建，比如可以直接从Gitrepo中构建：
```shell
docker build  https://github.com/xxx/xxxx.git#:4.5
```
这行命令指定了构建所需的 Git repo，并且指定默认的master分支，构建目录为/4.5/，然后Docker 就会自己去git clone这个项目、切换到指定分支、并进入到指定目录后开始构建。

#### 用给定的 tar 压缩包构建
```shell
docker build https://domain.com/xxx.tar.gz
```
如果所给出的 URL 不是个 Git repo，而是个  tar  压缩包，那么 Docker 引擎会下载这个包，并自动解压缩，以其作为上下文，开始构建。

#### 从标准输入中读取 Dockerfile 进行构建
```shell
docker build - < Dockerfile
```
```shell
cat Dockerfile | docker build -
```
#### 从标准输入中读取上下文压缩包进行构建
```shell
docker build - < xxx.tar.gz
```
如果发现标准输入的文件格式是gzip、bzip2以及xz的话，将会使其为上下文压缩包，直接将其展开，将里面视为上下文，并开始构建。
-------------------------------------------------------------------------------------------------

### Dockerfile指令详解
+ COPY 复制文件
    * COPY <源路径>... <目标路径>
    * COPY ["<源路径>",..."<目标路径>"]
 源路径指的是上下文目录
```dockerfile
COPY test.com.conf /etc/nginx/sites-enable/
COPY config/* /etc/nginx/sites-enable/
```
<源路径> 可以是多个，甚至可以是通配符，其通配符规则要满足Go的filepath.Match规则
<目标路径>可以是容器内的绝对路径，也可以是相对于工作目录的相对路径（工作目录可以用WORKDIR指令来指定）
使用COPY指令，源文件的各种元数据都会保留。比如读、写、执行权限、文件变更时间等。

+ ADD 更高级的复制文件
ADD  指令和COPY的格式和性质基本一致。但是在COPY基础上增加了一些功能。
比如<源路径>可以是一个URL：Docker会下载这个文件，下载后的文件权限自动设置为600[可能需要额外的权限设置]
<源路径>为一个tar压缩文件的话，压缩格式为gzip,bzip2以及xz的情况下，ADD指令将会自动解压缩这个压缩文件到<目标路径>去
尽可能的使用COPY，因为COPY的语义很明确，就是复制文件而已，而ADD则包含了更复杂的功能，其行为也不一定很清晰

+ CMD  容器启动命令
CMD  指令的格式和RUN相似，也是两种格式：
   * shell格式： CMD <命令>
   * exec格式： CMD ["可执行文件", "参数1", "参数2"...]
   * 参数列表格式： CMD ["参数1", "参数2"...]  在指定了ENTRYPOINT指令后，用CMD指定具体的参数。
Docker 不是虚拟机，容器就是进程. CMD指令就是用于指定默认的容器主进程的启动命令的。
在指令格式上，一般推荐使用exec格式,这类格式在解析时会被解析为JSON数组，因此一定要使用双引号"  ，而不要使用单引号。
如果使用shell格式的话，实际的命令会被包装为  sh -c的参数的形式进行执行.
容器中应用在前台执行和后台执行的问题:
Docker不是虚拟机，容器中的应用都应该以前台执行，而不是像虚拟机、物理机里面那样，用 upstart/systemd 去启动后台服务，容器内没有后台服务的概念。
对于容器而言，其启动程序就是容器应用进程，容器就是为了主进程而存在的，主进程退出，容器就失去了存在的意义，从而退出，其它辅助进程不是它需要关心的东西。

+ ENTRYPOINT 入口点
ENTRYPOINT 的格式和RUN指令格式一样，分为exec格式和shell格式。
ENTRYPOINT 的目的和CMD一样，都是在指定容器启动程序及参数。

+ ENV 设置环境变量
格式有两种：
ENV <key> <value>
ENV <key1>=<value1> <key2>=<value2>...
这个指令很简单，就是设置环境变量而已，无论是后面的其它指令，如RUN，还是运行时的应用，都可以直接使用这里定义的环境变量。
```dockerfile
ENV VERSION=1.0 DEBUG=on \
NAME="Kimmy"
```
```dockerfile
ENV NODE_VERSION 7.2.0
RUN curl -SLO "https://nodejs.org/dist/v$NODE_VERSION/node-v$NOD
E_VERSION-linux-x64.tar.xz" \
&& curl -SLO "https://nodejs.org/dist/v$NODE_VERSION/SHASUMS25
6.txt.asc" \
&& gpg --batch --decrypt --output SHASUMS256.txt SHASUMS256.tx
t.asc \
&& grep " node-v$NODE_VERSION-linux-x64.tar.xz\$" SHASUMS256.t
xt | sha256sum -c - \
&& tar -xJf "node-v$NODE_VERSION-linux-x64.tar.xz" -C /usr/loc
al --strip-components=1 \
&& rm "node-v$NODE_VERSION-linux-x64.tar.xz" SHASUMS256.txt.as
c SHASUMS256.txt \
&& ln -s /usr/local/bin/node /usr/local/bin/nodejs
```
+ ARG 构建参数
格式： ARG <参数名>[=<默认值>]
ARG所设置的构建环境的环境变量，在将来容器运行时是不会存在这些环境变量的.
Dockerfile中的ARG指令是定义参数名称，以及定义其默认值。该默认值可以在构建命令docker build中用  --build-arg <参数名>=<值> 来覆盖。

+ VOLUME 定义匿名卷
格式为：
    * VOLUME ["<路径1>", "<路径2>"...]
    * VOLUME <路径>
为了防止运行时用户忘记将动态文件所保存目录挂载为卷，在Dockerfile中，我们可以事先指定某些目录挂载为匿名卷，这样在运行时如果用户不指定挂载，其应用也可以正常运行，不会向容器存储层写入大量数据。
```dockerfile
VOLUME /data
```
这里的/data目录就会在运行时自动挂载为匿名卷，任何向/data中写入的信息都不会记录进容器存储层，从而保证了容器存储层的无状态化。
运行时可以覆盖这个挂载设置:
```shell
docker run -d -v mydata:/data xxxx
```
在这行命令中，就使用了mydata这个命名卷挂载到了/data这个位置，替代了Dockerfile中定义的匿名卷的挂载配置。

+ EXPOSE 声明端口
格式为: EXPOSE<端口1> [<端口2>...]
EXPOSE  指令是声明运行时容器提供服务端口，这只是一个声明，在运行时并不会因为这个声明应用就会开启这个端口的服务
一个是帮助镜像使用者理解这个镜像服务的守护端口，以方便配置映射；另一个用处则是在运行时使用随机端口映射时，也就是  docker run -P时，会自动随机映射EXPOSE的端口。
EXPOSE 仅仅是声明容器打算使用什么端口而已，并不会自动在宿主进行端口映射。

+ WORKDIR 指定工作目录
格式为  WORKDIR <工作目录路径>

使用  WORKDIR  指令可以来指定工作目录（或者称为当前目录），以后各层的当前目录就被改为指定的目录，该目录需要已经存在， WORKDIR  并不会帮你建立目录。

+ USER 指定当前用户
格式： USER <用户名>
USER 是改变之后层的执行RUN,CMD以及ENTRYPOINT这类命令的身份
USER  只是帮助你切换到指定用户而已，这个用户必须是事先建立好的，否则无法切换。
```dockerfile
RUN groupadd -r redis && useradd -r -g redis redis
USER redis
RUN [ "redis-server" ]
```
在执行期间希望改变身份,建议使用gosu
```dockerfile
# 建立 redis 用户，并使用 gosu 换另一个用户执行命令
RUN groupadd -r redis && useradd -r -g redis redis
# 下载 gosu
RUN wget -O /usr/local/bin/gosu "https://github.com/tianon/gosu/releases/download/1.7/gosu-amd64" \
&& chmod +x /usr/local/bin/gosu \
&& gosu nobody true
# 设置 CMD，并以另外的用户执行
CMD [ "exec", "gosu", "redis", "redis-server" ]
```

+ HEALTHCHECK 健康检查
   * HEALTHCHECK [选项] CMD <命令>  ：设置检查容器健康状况的命令
   * HEALTHCHECK NONE  ：如果基础镜像有健康检查指令，使用这行可以屏蔽掉其健康检查指令
HEALTHCHECK  指令是告诉 Docker 应该如何进行判断容器的状态是否正常.
当在一个镜像指定了HEALTHCHECK指令后，用其启动容器，初始状态会为starting，在HEALTHCHECK  指令检查成功后变为healthy  ，如果连续一定次数失败，则会变为unhealthy。
HEALTHCHECK  支持下列选项：
   * --interval=<间隔>  ：两次健康检查的间隔，默认为 30 秒；
   * --timeout=<时长>  ：健康检查命令运行超时时间，如果超过这个时间，本次健康检查就被视为失败，默认 30 秒；
   * --retries=<次数>  ：当连续失败指定次数后，则将容器状态视为unhealthy  ，默认 3 次。
```dockerfile
FROM nginx
RUN apt-get update && apt-get install -y curl && rm -rf /var/lib
/apt/lists/*
HEALTHCHECK --interval=5s --timeout=3s \
CMD curl -fs http://localhost/ || exit 1
```
这里我们设置了每 5 秒检查一次（这里为了试验所以间隔非常短，实际应该相对较长），如果健康检查命令超过 3 秒没响应就视为失败，并且使用  curl -fs http://localhost/ || exit 1  作为健康检查命令。

docker ps 查看容器状态
docker inspect 

+ ONBUILD
格式： ONBUILD <其它指令>
ONBUILD  是一个特殊的指令，它后面跟的是其它指令，比如RUN,COPY等，而这些指令，在当前镜像构建时并不会被执行。只有当以当前镜像为基础镜像，去构建下一级镜像的时候才会被执行.
Dockerfile  中的其它指令都是为了定制当前镜像而准备的，唯有  ONBUILD  是为了帮助别人定制自己而准备的。
```dockerfile
FROM node:slim
RUN "mkdir /app"
WORKDIR /app
ONBUILD COPY ./package.json /app
ONBUILD RUN [ "npm", "install" ]
ONBUILD COPY . /app/
CMD [ "npm", "start" ]
```
-------------------------------------------------------------------------------------------------
#### docker save 与docker load
用以将镜像保存为一个tar文件，然后传输到另一个位置上，再加载进来
#### 删除本地镜像 
```shell
docker rmi [选项] <镜像1> [<镜像2> ...]
```
注意  docker rm  命令是删除容器，不要混淆
```shell
docker rmi nginx
docker rmi 998ed1195e09
```
----------

### 操作 Docker 容器
简单的说，容器是独立运行的一个或一组应用，以及它们的运行态环境。
#### 启动容器
启动容器有两种方式，一种是基于镜像新建一个容器并启动，另外一个是将在终止状态（stopped）的容器重新启动。
因为 Docker 的容器实在太轻量级了，很多时候用户都是随时删除和新创建容器。
##### 新建并启动
所需要的命令主要为  docker run
```shell
sudo docker run -t -i ubuntu:14.04 /bin/bash
```
* -t 选项让Docker分配一个伪终端（pseudo-tty）并绑定到容器的标准输入上
* -i 让容器的标准输入保持打开
当利用  docker run  来创建容器时，Docker 在后台运行的标准操作包括：
检查本地是否存在指定的镜像，不存在就从公有仓库下载
利用镜像创建并启动一个容器
分配一个文件系统，并在只读的镜像层外面挂载一层可读写层
从宿主主机配置的网桥接口中桥接一个虚拟接口到容器中去
从地址池配置一个 ip 地址给容器
执行用户指定的应用程序
执行完毕后容器被终止
##### 启动已终止容器
可以利用  docker start  命令，直接将一个已经终止的容器启动运行。
#### 后台(background)运行
更多的时候，需要让 Docker在后台运行而不是直接把执行命令的结果输出在当前宿主机下。此时，可以通过添加  -d  参数来实现。
docker ps 查看容器信息
要获取容器的输出信息，可以通过  docker logs  命令
docker logs [container ID or NAMES]
##### 终止容器
可以使用  docker stop  来终止一个运行中的容器。
终止状态的容器可以用  docker ps -a  命令看到
处于终止状态的容器，可以通过  docker start  命令来重新启动。
docker restart  命令会将一个运行态的容器终止，然后再重新启动它。
#### 进入容器
在使用  -d  参数时，容器启动后会进入后台。 某些时候需要进入容器进行操作，有很多种方法，包括使用docker attach命令或nsenter工具等。
##### attach 命令
docker attach  是Docker自带的命令。下面示例如何使用该命令。
```shell
docker attach [OPTIONS] CONTAINER
```
但是使用  attach  命令有时候并不方便。当多个窗口同时 attach 到同一个容器的时候，所有窗口都会同步显示。当某个窗口因命令阻塞时,其他窗口也无法执行操作了。
##### nsenter 命令
nsenter 需要先安装
nsenter  启动一个新的shell进程(默认是/bin/bash), 同时会把这个新进程切换到和目标(target)进程相同的命名空间，这样就相当于进入了容器内部。
为了连接到容器，你还需要找到容器的第一个进程的 PID，可以通过下面的命令获。
```shell
PID=$(docker inspect --format "{{ .State.Pid }}" <container>)
```
通过这个 PID，就可以连接到这个容器：
```shell
nsenter --target $PID --mount --uts --ipc --net --pid
```
```shell
$ sudo docker run -idt ubuntu
243c32535da7d142fb0e6df616a3c3ada0b8ab417937c853a9e1c251f499f550
$ sudo docker ps
CONTAINER ID IMAGE COMMAND CREA
TED STATUS PORTS NAMES
243c32535da7 ubuntu:latest "/bin/bash" 18 s
econds ago Up 17 seconds nostalgi
c_hypatia
$ PID=$(docker-pid 243c32535da7)
10981
$ sudo nsenter --target 10981 --mount --uts --ipc --net --pid
```
wget -P ~ https://github.com/yeasy/docker_practice/raw/master/_local/.bashrc_docker

#### 导出和导入容器
##### 导出容器
如果要导出本地某个容器，可以使用  docker export  命令。
```shell
$ sudo docker ps -a
CONTAINER ID IMAGE COMMAND CREA
TED STATUS PORTS NA
MES
7691a814370e ubuntu:14.04 "/bin/bash" 36 h
ours ago Exited (0) 21 hours ago te
st
$ sudo docker export 7691a814370e > ubuntu.tar
```
这样将导出容器快照到本地文件。

##### 导入容器快照
可以使用  docker import  从容器快照文件中再导入为镜像
```shell
$ cat ubuntu.tar | sudo docker import - test/ubuntu:v1.0
$ sudo docker images
REPOSITORY TAG IMAGE ID CREA
TED VIRTUAL SIZE
test/ubuntu v1.0 9d37a6082e97 Abou
t a minute ago 171.3 MB
```
此外，也可以通过指定 URL 或者某个目录来导
```shell
sudo docker import http://example.com/exampleimage.tgz example/imagerepo
```
*注：用户既可以使用  docker load  来导入镜像存储文件到本地镜像库，也可以使用  docker import  来导入一个容器快照到本地镜像库。这两者的区别在于容器快照文件将丢弃所有的历史记录和元数据信息（即仅保存容器当时的快照状态），而镜像存储文件将保存完整记录，体积也要大。此外，从容器快照文件导入时可以重新指定标签等元数据信息。
#### 删除容器
可以使用  docker rm  来删除一个处于终止状态的容器
```shell
docker rm [OPTIONS] CONTAINER [CONTAINER...]
```
如果要删除一个运行中的容器，可以添加  -f  参数。Docker 会发送SIGKILL信号给容器。
##### 清理所有处于终止状态的容器
```shell
docker rm $(docker ps -a -q)
```
----
### 访问仓库
仓库（Repository）是集中存放镜像的地方。
一个容易混淆的概念是注册服务器（Registry）。
实际上注册服务器是管理仓库的具体服务器，每个服务器上可以有多个仓库，而每个仓库下面有多个镜像。从这方面来说，仓库可以被认为是一个具体的项目或目录。例如对于仓库地址dl.dockerpool.com/ubuntu  来说，dl.dockerpool.com  是注册服务器地址， ubuntu  是仓库名。

#### Docker Hub
目前 Docker 官方维护了一个公共仓库 Docker Hub
##### 登录
```shell
docker login
```
##### 基本操作
* 搜索镜像
```shell
docker search [name]
docker search nginx
```
* 下载镜像
```shell
docker pull [name]
docker pull nginx
```
#### 自动创建
自动创建允许用户通过Docker Hub指定跟踪一个目标网站（目前支持 GitHub或 BitBucket）上的项目，一旦项目发生新的提交，则自动执行创建。

要配置自动创建，包括如下的步骤：
* 创建并登录 Docker Hub，以及目标网站
* 在目标网站中连接帐户到 Docker Hub
* 在 Docker Hub中配置一个自动创建
* 选取一个目标网站中的项目（需要含Dockerfile）和分支
* 指定Dockerfile的位置，并提交创建

#### 私有仓库
有时候使用Docker Hub这样的公共仓库可能不方便，用户可以创建一个本地仓库供私人使用。
`docker-registry`  是官方提供的工具，可以用于构建私有的镜像仓库。
##### 安装运行docker-registry
1. 容器运行
```shell
docker run -d -p 5000:5000 registry
```
在安装了Docker后，可以通过获取官方 registry 镜像来运行。
这将使用官方的 registry 镜像来启动本地的私有仓库
```shell
docker run -d -p 5000:5000 -v /opt/data/registry:/var/lib/registry  -v /data/config.yml:/etc/docker/registry/config.yml  registry
```

2. 本地安装
```shell
sudo pip install docker-registry
```

##### 在私有仓库上传、下载、搜索镜像
* 使用docker tag 标记镜像tag
```
docker tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]
```
```
docker tag IMAGE[:TAG] [REGISTRYHOST/][USERNAME/]NAME[:TAG]
```
```
sudo docker tag  fce289e99eb9 192.168.42.128:5000/test
```
* 使用docker push 上传标记的镜像
```
docker push 192.168.42.128:5000/test
```
* 使用docker pull 下载私有仓库镜像
```
docker pull 192.168.42.128:5000/test
```
##### 仓库配置文件
Docker的Registry利用配置文件提供了一些仓库的模板（flavor），用户可以直接使用它们来进行开发或生产部署。
在config_sample.yml  文件中，可以看到一些现成的模板段：
* `common`  ：基础配置
* `local`  ：存储数据到本地文件系统
* `s3`  ：存储数据到 AWS S3 中
* `dev`  ：使用  local  模板的基本配置
* `test`  ：单元测试使用
* `prod`  ：生产环境配置（基本上跟s3配置类似）
* `gcs`  ：存储数据到 Google 的云存储
* `swift`  ：存储数据到 OpenStack Swift 服务
* `glance`  ：存储数据到 OpenStack Glance 服务，本地文件系统为后备
* `glance-swift`  ：存储数据到 OpenStack Glance 服务，Swift 为后备
* `elliptics`  ：存储数据到 Elliptics key/value 存储
示例：
```yml
common:
    loglevel: info
    search_backend: "_env:SEARCH_BACKEND:"
    sqlalchemy_index_database: "_env:SQLALCHEMY_INDEX_DATABASE:sqlite:////tmp/docker-registry.db"
prod:
    loglevel: warn
    storage: s3
    s3_access_key: _env:AWS_S3_ACCESS_KEY
    s3_secret_key: _env:AWS_S3_SECRET_KEY
    s3_bucket: _env:AWS_S3_BUCKET
    boto_bucket: _env:AWS_S3_BUCKET
    storage_path: /srv/docker
    smtp_host: localhost
    from_addr: docker@myself.com
    to_addr: my@myself.com
dev:
    loglevel: debug
    storage: local
    storage_path: /home/myself/docker
test:
    storage: local
    storage_path: /tmp/tmpdockertmp
```
------------------------------------------------------------------------------------------
## Docker 数据管理
在容器中管理数据主要有 两种方式：
* 数据卷（Data volumes）
* 数据卷容器（Data volume containers）

### 数据卷（Data volumes）
数据卷是一个可供一个或多个容器使用的特殊目录，它绕过 UFS，可以提供很多有用的特性：
* 数据卷可以在容器之间共享和重用
* 对数据卷的修改会立马生效
* 对数据卷的更新，不会影响镜像
* 数据卷默认会一直存在，即使容器被删除

类似于 Linux 下对目录或文件进行 mount，镜像中的被指定为挂载点的目录中的文件会隐藏掉，能显示看的是挂载的数据卷
#### 创建一个数据卷
在用`docker run`命令的时候，使用`-v`标记来创建一个数据卷并挂载到容器里。在一次run中多次使用可以挂载多个数据卷
```
docker run -d -P --name web -v /webapp training/webapp python app.py
```
#### 删除数据卷
数据卷是被设计用来持久化数据的，它的生命周期独立于容器
可以在删除容器的时候使用`docker rm -v`  这个命令。
#### 挂载一个主机目录作为数据卷
使用  -v  标记也可以指定挂载一个本地主机的目录到容器中去
```
docker run -d -P --name web -v /src/webapp:/opt/webapp training/webapp python app.py
```
上面的命令加载主机的/src/webapp目录到容器的/opt/webapp目录。这个功能在进行测试的时候十分方便
Docker 挂载数据卷的默认权限是读写，用户也可以通过  :ro  指定为只读
```
docker run -d -P --name web -v /src/webapp:/opt/webapp:ro training/webapp python app.py
```
查看指定容器信息
```
docker inspect [容器名称]
```
#### 挂载一个本地主机文件作为数据卷
`-v`标记也可以从主机挂载单个文件到容器中
```
docker run --rm -it -v ~/.bash_history:/.bash_history ubuntu /bin/bash
```
这样就可以记录在容器输入过的命令了

### 数据卷容器
如果你有一些持续更新的数据需要在容器之间共享，最好创建数据卷容器。
数据卷容器，其实就是一个正常的容器，专门用来提供数据卷供其它容器挂载的。
首先，创建一个名为 dbdata 的数据卷容器：
```
docker run -d -v /dbdata --name dbdata training/postgres echo Data-only container for postgres
```
在其他容器中使用`--volumes-from`来挂载 dbdata 容器中的数据卷
```
docker run -d --volumes-from dbdata --name db1 training/postgres
docker run -d --volumes-from dbdata --name db2 training/postgres
```
可以使用超过一个的`--volumes-from`参数来指定从多个容器挂载不同的数据卷。也可以从其他已经挂载了数据卷的容器来级联挂载数据卷。

如果要删除一个数据卷，必须在删除最后一个还挂载着它的容器时使用`docker rm -v`命令来指定同时删除关联的容器。

-----------------------------------------------------------------------------------------
##Docker 中的网络功能介绍
Docker 允许通过外部访问容器或容器互联的方式来提供网络服务
### 一、外部访问容器
容器中可以运行一些网络应用，要让外部也可以访问这些应用，可以通过`-P`或`-p`参数来指定端口映射。
当使用`-P`标记时，Docker会随机映射一个49000~49900的端口到内部容器开放的网络端口。
可以使用`docker ps`与`docker logs`端口与应用信息
`-p`（小写的）则可以指定要映射的端口，并且在一个指定端口上只可以绑定一个容器。
支持的格式有  ip:hostPort:containerPort | ip::containerPort | hostPort:containerPort
#### 映射所有接口地址
```
docker run -d -p 5000:5000 training/webapp python app.py
```
此时默认会绑定本地所有接口上的所有地址

#### 映射到指定地址的指定端口
可以使用  ip:hostPort:containerPort  格式指定映射使用一个特定地址，比如localhost 地址 127.0.0.1
```
docker run -d -p 127.0.0.1:5000:5000 training/webapp python app.py
```
#### 映射到指定地址的任意端口
使用  ip::containerPort  绑定 localhost 的任意端口到容器的 5000 端口，本地主机会自动分配一个端口
```
docker run -d -p 127.0.0.1::5000 training/webapp python app.py
```
#### 还可以使用 udp 标记来指定 udp 端口
```
docker run -d -p 127.0.0.1:5000:5000/udp training/webapp python app.py
```
#### `docker port`查看映射端口配置
```
docker port nostalgic_morse 5000 127.0.0.1:49155.

 docker port focused_bell
```
注意：
   * 容器有自己的内部网络和ip地址（使用docker inspect可以获取所有的变量，Docker还可以有一个可变的网络配置）
   * `-p` 标记可以多次使用来绑定多个端口
```
docker run -d -p 5000:5000 -p 3000:80 training/webapp python app.py
```
### 二、容器互联
容器的连接（linking）会在源和接收容器之间创建一个隧道，接收容器可以看到源容器指定的信息。
#### 自定义容器命名
连接系统依据容器的名称来执行
`--name`标记可以为容器自定义命名
```
docker run -d -P --name web training/webapp python app.py
```
容器的名称是唯一的,需要先用docker rm来删除之前创建的同名容器
在执行docker run的时候如果添加`--rm` 标记，则容器在终止后会立刻删除
注意， --rm  和  -d  参数不能同时使用
#### 容器互联
使用`--link`参数可以让容器之间安全的进行交互

```shell
## 先创建一个新的数据库容器
docker run -d --name db training/postgres
```
```
docker run -d -P --name web --link db:db training/webapppython app.py
```
`--link`参数的格式为`--link name:alias`，其中name是要链接的容器的名称,alias是这个连接的别名。
Docker 在两个互联的容器之间创建了一个安全隧道，而且不用映射它们的端口到宿主主机上。
用户可以链接多个父容器到子容器，比如可以链接多个 web 到 db 容器上。
Docker 通过 2 种方式为容器公开连接信息：
* 环境变量
* 更新/etc/hosts文件

-------------------------------------------------------------------------------------------------
## 高级网络配置
当Docker启动时，会自动在主机上创建一个`docker0`虚拟网桥，实际上是Linux 的一个bridge，可以理解为一个软件交换机。它会在挂载到它的网口之间进行转发。
### 快速配置指南
下面是一个跟Docker网络相关的命令列表。
其中有些命令选项只有在Docker服务启动的时候才能配置，而且不能马上生效
* -b BRIDGE or --bridge=BRIDGE  --指定容器挂载的网桥
* --bip=CIDR  --定制 docker0 的掩码
* -H SOCKET... or --host=SOCKET...  --Docker 服务端接收命令的通道
* --icc=true|false  --是否支持容器之间进行通信
* --ip-forward=true|false  --请看下文容器之间的通信
* --iptables=true|false  --是否允许 Docker 添加 iptables 规则
* --mtu=BYTES  --容器网络中的 MTU
下面2个命令选项既可以在启动服务时指定，也可以 Docker 容器启动（dockerrun）时候指定。在 Docker 服务启动的时候指定则会成为默认值，后面执行docker run 时可以覆盖设置的默认值。
* --dns=IP_ADDRESS...  --使用指定的DNS服务器
* --dns-search=DOMAIN...  --指定DNS搜索域
最后这些选项只有在  docker run  执行时使用，因为它是针对容器的特性内容
* -h HOSTNAME or --hostname=HOSTNAME  --配置容器主机名
* --link=CONTAINER_NAME:ALIAS  --添加到另一个容器的连接
* --net=bridge|none|container:NAME_or_ID|host  --配置容器的桥接模式
* -p SPEC or --publish=SPEC  --映射容器端口到宿主主机
* -P or --publish-all=true|false  --映射容器所有端口到宿主主机
### 配置 DNS
怎么自定义配置容器的主机名和 DNS配置呢?秘诀就是它利用虚拟文件来挂载到来容器的3个相关配置文件
在容器中使用 mount 命令可以看到挂载信息：
这种机制可以让宿主主机 DNS 信息发生更新后，所有 Docker 容器的 dns 配置通过  /etc/resolv.conf  文件立刻得到更新。
如果用户想要手动指定容器的配置，可以利用下面的选项。
* -h HOSTNAME or --hostname=HOSTNAME设定容器的主机名，它会被写到容器内的/etc/hostname和/etc/hosts。但它在容器外部看不到，既不会在docker ps中显示，也不会在其他的容器的/etc/hosts看到。
* --link=CONTAINER_NAME:ALIAS选项会在创建容器的时候，添加一个其他容器的主机名到/etc/hosts  文件中，让新容器的进程可以使用主机名 ALIAS 就可以连接它。
* --dns=IP_ADDRESS添加DNS服务器到容器的/etc/resolv.conf中，让容器用这个服务器来解析所有不在  /etc/hosts中的主机名。
* --dns-search=DOMAIN  设定容器的搜索域，当设定搜索域为  .example.com时，在搜索一个名为 host 的主机时，DNS 不仅搜索host，还会搜索host.example.com。注意：如果没有上述最后 2 个选项，Docker 会默认用主机上的  /etc/resolv.conf  来配置容器。
### 容器访问控制
容器的访问控制，主要通过 Linux上的iptables防火墙来进行管理和实现。`iptables`是Linux上默认的防火墙软件，在大部分发行版中都自带。
#### 容器访问外部网络
容器要想访问外部网络，需要本地系统的转发支持。在Linux 系统中，检查转发是否打开
```
$sysctl net.ipv4.ip_forward
net.ipv4.ip_forward = 1
```
如果为 0，说明没有开启转发，则需要手动打开。
```
sysctl -w net.ipv4.ip_forward=1
```
如果在启动Docker服务的时候设定`--ip-forward=true`,Docker就会自动设定系统的ip_forward参数为 1。
#### 容器之间访问
容器之间相互访问，需要两方面的支持。
* 容器的网络拓扑是否已经互联。默认情况下，所有容器都会被连接到docker0网桥上。
* 本地系统的防火墙软件-- iptables是否允许通过。
#### 访问所有端口
当启动Docker服务时候，默认会添加一条转发策略到iptables的FORWARD链上。策略为通过（ ACCEPT  ）还是禁止（DROP）取决于配置 --icc=true（缺省值）还是--icc=false  。当然，如果手动指定  --iptables=false  则不会添加iptables规则。
可见，默认情况下，不同容器之间是允许网络互通的。如果为了安全考虑，可以在/etc/default/docker  文件中配置DOCKER_OPTS=--icc=false来禁止它。
#### 访问指定端口
在通过-icc=false关闭网络访问后，还可以通过--link=CONTAINER_NAME:ALIAS选项来访问容器的开放端口。


```shell
##进入容器命令
docker exec -it 容器名 bash

```

 ---
 #### Docker常用命令
 * Docker列出容器IP : `docker inspect 容器ID | grep IPAddress`
 * docker run -d --hostname my-rabbit --name rabbit -e RABBITMQ_DEFAULT_USER=root -e RABBITMQ_DEFAULT_PASS=root -p 15672:15672 -p 5672:5672 -p 25672:25672 -p 61613:61613 -p 1883:1883 rabbitmq:3-management







https://www.pornhub.com/view_video.php?viewkey=ph5ac7df45bbf1a