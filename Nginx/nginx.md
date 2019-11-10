Nginx 
=====

### HTTP




### Nginx安装
#### Linux服务器
##### VMware网络配置
+ 桥接模式 VMnet0 虚拟机与物理机连接到同一个局域网段，而且处于同一IP段
+ NAT模式 VMnet8  默认模式,只有物理机能上网虚拟机就能上网 
+ 权主机械 VMnet1 只有物理机能上网，虚拟机不能上网，只能在VMnet1虚拟网互访
#### Linux网络配置
##### 动态IP
```shell
#暂时启动eth0网卡
ifup eth0
```
centos 网络配置文件
```shell
## /etc/sysconfig/network-scrips/ifcfg-eth0
##是否自动启动
ONBOOT =yes
#用过设置获取IP的方式静态还是动态 dhcp为动态获取 static为静态
BOOTPROTO = dhcp
#以下配置只在静态IP地址时需要添加设置
# IP地址
IPADDRESS = 192.168.78.6
NETMASK = 255.255.255.0
GATEWAY = 192.168.78.2
DNS1 = 192.168.78.2
```
service network reload 使配置生效

------------------------------------------------------------------
Nginx 基本配置

1.用户和组[设置的是worker_process的用户和组]
    user www;
    group www;
-------------------------------------------------
2. 自定义错误页面
  ```config
  error_page 404 /40x.html
  error_page 404 403 402 /40.html
  errop_page 404 https://www.***.html
  error_page 404=200 /40x.html
  ```
### 访问控制  
nginx提供了2个用于配置访问控制的指令：
* allow 用于设置允许访问的权限
* deny 用于设置禁止访问的权限
用法：
allow或者deny 后面跟上允许或禁止的IP、IP段或者all即可
注意事项：
* 同一配置块下,若同时出现多个权限指令(allow|deny),则先出现的权限设置生效,并且会对后出现的设置进行覆盖,未覆盖的范围依然有效。
* 当多个块[http|server|location]中都出现了权限设置指令，内层块中的权限级别比外层块中设置的权限级别高。
```
##禁止所有
deny all;
#允许特定的IP
allow 192.168.33.100;
```
#### location 块
=   根据其后的指定模式进行精准匹配
~   使用正则表达式完成location的匹配，区分大小写
~*  使用正则表达式完成location的匹配，不区分大小写
^~  不使用正则表达式，完成以指定模式开头的location匹配
@   用于定义一个location块,且该块不能被外部客户端所访问,只能被Nginx内部配置指令所访问


#### 日志文件[访问日志、错误日志]
访问日志：
    log_format 

    * $remote_addr  客户端的IP地址
    * $remote_user  客户端用户,如果没有则为空
    * $time_local   访问时间与时区
    * $request      请求的URI和HTTP协议，如GET/HTTP1.1
    * $status       记录返回的http状态码,如200
    * $body_bytes_sent 发送给客户端的文件主体内容的大小,如999
    * $http_referer 来路URL地址
    * $http_user_agent  客户端浏览器信息
    * $http_x_forwarded_for 客户端IP地址列表(包括中间经过的代理)

###### 注意事项：
    nginx默认开启访问日志功能,且log_format指令的配置权可用在http块内
    若在访问过程中需要记录子请求的日志记录，则可以将log_subrequest指令设置为on

    ```config
    #关闭访问日志
    access_log off;
    #
    access_log /var/log/nginx/access.log;
    ```
错误日志：
    ```
    #关闭错误日志
    error_log /dev/null;
    error_log /var/logs/nginx/error.log;
    error_log /var/logs/nginx/error.log notice;
    error_log /var/logs/nginx/error.log info;
    ```
error_log 指令的第一个参数指定日志存放的路径，第二个参数用于指定错误详细程度的等级。[debug|info|notice|warn|error|crit] 默认等级为error
###### 注意事项： 
    error_log 可以在main、http、service、location块中设置

> /dev/null 2>&1

### 虚拟主机
    虚拟主机技术是指在一台物理主机服务上划分多个磁盘空间，每个磁盘空间都是一台虚拟主机，每台虚拟主机都可以独立对外提供Web服务，且互不干扰。
    [意味着可以把不同域名或多个网站放在同一台服务器上]

    在http块下增加server块实现增加虚拟主机.

1. 基于端口配置虚拟主机
    它的原理就是一个nginx监听多个端口，根据不同的端口号，来区分不同的网站。
    修改liten来设置不同端口。
    listen 80;
    listen 8001;
2. 基于IP配置虚拟主机

3. 基于域名配置虚拟主机
    


    #开启目录列表[开启后才能显示目录结构]
    autoindex on;
    当autoindex 设置为on后，还可以通过autoindex_exact_size指令设置精准显示文件大小还是大概显示文件大小;通过autoindex_localtime指令设置最后修改时候的显示格式;
    autoindex_exact_size off;
    autoindex_localtime on;



### Nginx + PHP
编译安装PHP
1. 到官网下载最新或者合适的PHP版本
2. 解压后进入php目录, `./configure` 命令用于编译安装
    `./configure --help` 可以查看详细编译选项
以下列出常用选项



|选项|说明|
|----|----|
| --prefix | 安装目录,默认/usr/loca |
| --enable-fpm | 开户PHP的FPM功能,提供PHP FastCGI管理器 |
| --with-zlib | 包含zlib库,支持数据库压缩和解压缩 |
| --with-zip | 开启zip功能 |
| --enable-mbstring | 开启mbstring功能,用于多字节字符串处理 |
| --with-mcrypt | 包含mcrypt加密支持(依赖libmcrypt) |
| --with-mysql | 包含MySQL数据库访问支持 |
| --with-mysqli | 包含增强版MySQL数据库访问支持 |
| --with-pdo-mysql | 包含基于PDO(PHP Data Project)的MySQL数据库访问支持 |
| --with-gd | 包含GD库支持,用于PHP图像处理 |
| --with-jpeg-dir | 包含JPGE图像格式处理库(依赖libjpeg-devel) |
| --with-png-dir | 包含PNG图像格式处理库(依赖libpng-devel) |
| --with-freetype-dir | 包含FreeType字体图像处理库(依赖freetype-devel) |
| --with-curl | 包含curl支持(依赖curl-devel) |
| --with-openssl | 包含OpenSSL支持(依赖openssl-devel) |
| --with-mhash | 包含mhash加密支持 |
| --enable-bcmath | 开启精准计算功能 |
| --enable-opcache | 开启opcache功能,一种PHP的代码优化器 |

以上前缀为with的选项依赖于系统的共享库，如果系统中没有则需要安装依赖包.
1) 通过yum安装依赖
```shell
yum -y install libxml2-devel openssl-devl curl-devel libjpeg-devel libpng-devel libpng-devel freetype-devel
```

2) 安装PHP
```shell
##进入PHP目录
./configure --prefix /usr/local/php --enable-fpm --with-zlib --with-zip --enable-mbstring --with-mcrypt --with-mysql --with-mysqli --with-pdo-mysql --with-gd --with-jpeg-dir --with-png-dir --with-freetype-dir --with-curl --with-openssl --with-mhash --enable-bcmath --enable-opcache
```
执行以上代码后，如果有提示缺少某些依赖库的，只要安装相关依赖库即可.

3) 测试PHP是否安装成功
```shell
php -r "echo 'Hello World!';"
```
#### Nginx与PHP整合
Nginx与PHP整合离不开FastCGI.
对于Nginx而言,PHP是一个外部程序，而非一个Nginx模块。为了web服务器的功能扩展性更强，就出现了CGI(Common Gateway Interface 公共网关接口)规范。
CGI是Web服务器与外部程序之间的接口标准，用于2种不同程序之间的信息传递。

在实现配置Nginx时,可以利用Location块规则,实现不同URI请求采取不同的处理方式.
PHP提供的PHP-FPM(FastCGI Process Manager)就是一个FastCGI进程管理器。
PHP-FPM位于PHP安装目录下的sbin目录中.在启动PHP-FPM之前要先配置php-fpm.config文件

##### PHP配置文件
PHP的配置文件主要包括php.ini 和php-fpm.config,在更改这2个文件之后都需要重启PHP-FPM服务。

php-fpm.cofig
