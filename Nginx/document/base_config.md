Nginx 基础配置
-------------

#### Nginx指令和指令块
* 简单指令: ``由名称和参数组成,以空格分隔，以分号结尾``
```config
Main 1;
```
* 指令块: ``由名称和大括号{}内的附加指令组成,不以分号结尾``
```config
http {
    Main 2;
}
```
`http`块是全局参数,对整体产生影响;`server`块是虚拟主机,主要对指定的主机和端口进行配置;`location`块在虚拟主机下根据请求`URI(Uniform Resource Identifier统一资源标识)`进行配置,URI即去掉参数后的URL

#### Nginx基本配置说明
##### main配置
在`http`块之前的配置是全局参数，全局参数对整个`Nginx`块都产生作用,例如:
```config
user www;
error_log;
work_processes 1;
http {

}
```
##### `http`块 与客户端有关的配置
其作用是处理与客户端相关的信息
###### http客户端配置常用的指令

|指令|说明|
|-|-|
|`client_body_buffer_size`|设置读取客户端请求体的缓冲区大小。如果请求体的大小大于缓冲区的大小，则整个或一部分请求体会被写入临时文件。 64位系统默认为`16KB`,32位为`8KB`|
|`client_body_temp_path`|定义存储客户端请求体的临时文件目录，最多可以定义3个子集目录|
|`client_body_timeout`|定义读取客户端请求体的超时时间，即两个连续的读操作之间的时间间隔.如果超时HTTP会抛出`408`错误|
|`client_header_buffer_size`|设置客户端请求头的缓冲区大小,默认为`1KB`|
|`client_max_body_size`|设置客户端请求的最大主体的大小,默认为`1MB`|
|`client_header_timeout`|请求客户端请求头的超时时间|
|`etag`|如果设置为`on`表示静态资源自动生成`ETag`响应头|
|`large_client_header_buffers`|设置大型客户端请求头的缓冲区大小|
|`keepalive_timeout`|设置连接超时时间,服务器将在超过超时时间后关闭http连接|
|`send_timeout`|指定客户端的响应超时时间|
|`server_names_hash_bucket_size`|设置`server_names(Nginx中设置的全部域名)`散列表的桶的大小，默认值取决于处理器缓存行的大小|
|`server_names_hash_max_size`|设置`server_names`散列表的最大值|
|`server_tokens`|启用或禁用在错误页面和服务器响应头字段中标识的Nginx版本|
|`tcp_nodelay`|启用或禁用`TCP_NODELAY`选项,只有当连接保持活动时,才会被启用|
|`tcp_nopush`|仅当`sendfile`时使用,能够将响应头和正文的开始部分一起发送|
其中有些指令可以设置在其它块中
##### `server`块,即虚拟主机部分
如果请求中的`Host`头和`server_name`相匹配，则将请求指向对应的`server`块
```config
server {
    server_name jimmy.com www.jimmy.com;
}
```
`server_name`支持使用通配符正则表达式,支持多域名,服务名称.当有多个`server`块时，会存在匹配优先级的问题,优先级顺序如下:
 1. 精确的名字
 2. 以`*`开头的最长通配符名称,如`*abc.com;`
 3. 以`*`结尾的最长通配符名称,如`www.abc.*;`
 4. 按照文件顺序,第1个匹配到的正则表达式;
 5. 如果没有匹配到对应的`server_name`,则会访问`default_server`
##### `loocation`块
`location`块在`http`块中使用,它的作用是根据客户端请求`URL`去定位不同的应用.当服务器接收到客户端请求后,需要在服务器端指定目录中去寻找客户端请求的资源,这就需要使用请求`URL`匹配对应的`location`指令.
##### URL在`location`块中的匹配规则说明
|配置格式|作用|
|--|--|
|`location=/uri`|`=`表示精确匹配|
|`location^~/uri`|`^~`匹配以某个URL前缀开头的请求,不支持正则表达式|
|`locaton~`|`~`区分大小写的匹配,属于正则表达式|
|`location~*`|`~*`不区分大小写的匹配,属于正则表达式|
|`location /uri`|` `表示前缀匹配,不带修饰符,但是优先级没有正则表达式高|
|`location /`|通用匹配,默认找不到其他匹配时,会进行通用匹配|
|`location @`|命名空间,不提供常规的请求匹配|
优先级: `=` > `^~` > `~` > `/uri` > `/`
如果找不到其他配置,就会进行通用匹配;`@`表示命名空间的位置,通常在重定向时进行匹配,且不会改变URL的原始请求.
可以打开`Debug`模式并观察日志,会看到每个请求的执行过程,包括匹配到对应的`location`的操作.
`location`支持嵌套配置:
```nginx
location /a {
    location /a {

    }
}
```
有些指令只能在`location`块中执行:
* `internal`:表示该`location`块只支持`Nginx`内部的请求访问,如支持`rewrite`、`error_page`等重定向,但不能通过外部的`HTTP`直接访问
* `limit_except`: 限定该`location`块可以执行的`HTTP`方法,如`GET`
* `alias`: 定义指定位置的替换,如可以使用以下配置:
```nginx
location /a/ {
    alias /c/x/a/;
}
```
上述配置表示如果匹配到/a/test.json的请求,在进入`location`块后,会将请求变成/c/x/a/test.json.
#### `include`的使用
`include`用来指定主配置文件包含的其他扩展配置文件.扩展文件的内容也要符合`Nginx`的格式规范.`include`可以出现在全局参数、`location`块、`server`块等任何一个位置
`include`支持通配符:
```nginx
include config/vhost/*.conf;
```
#### 常见配置
```nginx
########### 每个指令必须有分号结束。#################
user www www;  #定义运行Nginx用户和用户组，默认为nobody nobody。
worker_processes 2;  #允许生成的进程数，默认为1
pid /nginx/pid/nginx.pid;   #指定nginx进程运行文件存放地址
error_log log/error.log debug;  #制定日志路径，级别。这个设置可以放入全局块，http块，server块，级别以此为：debug|info|notice|warn|error|crit|alert|emerg
events {
    accept_mutex on;   #设置网路连接序列化，防止惊群现象发生，默认为on
    multi_accept on;  #设置一个进程是否同时接受多个网络连接，默认为off
    #use epoll;      #事件驱动模型，select|poll|kqueue|epoll|resig|/dev/poll|eventport
    worker_connections  1024;    #最大连接数，默认为512
}
http {
    include       mime.types;   #文件扩展名与文件类型映射表
    default_type  application/octet-stream; #默认文件类型，默认为text/plain
    #access_log off; #取消服务日志    
    log_format myFormat '$remote_addr–$remote_user [$time_local] $request $status $body_bytes_sent $http_referer $http_user_agent $http_x_forwarded_for'; #自定义格式
    access_log log/access.log myFormat;  #combined为日志格式的默认值
    sendfile on;   #允许sendfile方式传输文件，默认为off，可以在http块，server块，location块。
    sendfile_max_chunk 100k;  #每个进程每次调用传输数量不能大于设定的值，默认为0，即不设上限。
    keepalive_timeout 65;  #连接超时时间，默认为75s，可以在http，server，location块。

    upstream mysvr {   
      server 127.0.0.1:7878;
      server 192.168.10.121:3333 backup;  #热备
    }
    error_page 404 https://www.baidu.com; #错误页
    server {
        keepalive_requests 120; #单连接请求上限次数。
        listen       4545;   #监听端口
        server_name  127.0.0.1;   #监听地址       
        location  ~*^.+$ {       #请求的url过滤，正则匹配，~为区分大小写，~*为不区分大小写。
           #root path;  #根目录
           #index vv.txt;  #设置默认页
           proxy_pass  http://mysvr;  #请求转向mysvr 定义的服务器列表
           deny 127.0.0.1;  #拒绝的ip
           allow 172.18.5.54; #允许的ip           
        } 
    }
}
```
#### 内置变量
在客户端请求过程中,`Nginx`提供了内置变量来获取`HTTP`或`TCP`的信息.
##### 常见内置变量
|变量名|说明|
|--|--|
|`$arg_name`|指URL请求中的参数,name是参数的名字,如`$arg_id`|
|`$args`|代表URL中所有请求的参数|
|`$binary_remote_addr`|客户端地址以二进制数据的形式出现,通常会和限速模块一起使用|
|`$body_bytes_sent`|发送给客户端的字节数,不包含响应头|
|`$bytes_sent`|发送客给户端的总字节数|
|`$document_uri`|设置`$uri`的别名|
|`$hostname`|运行Nginx的服务器名|
|`$http_referer`|表示请求是从哪个页面链接过来的|
|`$http_user_agent`|客户端浏览器的相关信息|
|`$remote_addr`|客户端IP地址|
|`$remote_port`|客户端端口号|
|`$remote_user`|客户端用户名,通常在`Auth Basic`模块中使用|
|`$request_filename`|请求的文件路径,基于`root alias`指令和`URI`请求生成|
|`$request_time`|请求被`Nginx`接收后,一直到响应数据返回给客户端所用的时间|
|`$request_uri`|请求的`URI`,带 参数|
|`$request`|记录请求的`URL`和`HTTP`|
|`$request_length`|请求的长度,包括请求行、请求头和请求正文|
|`$server_name`|虚拟主机的`server_name`的值,通常是域名|
|`$server_port`|服务器端口|
|`$server_addr`|服务器的IP地址|
|`$request_method`|请求的方式,如`POST`和`GET`|
|`$scheme`|请求协议,如`HTTP`或`HTTPS`|
|`$sent_http_name`|任意响应头,`name`为响应头的名字,注意`name`要小写|
|`$realip_remote_addr`|保留原来的客户端地址,在`real_ip`模块中使用|
|`$server_protocol`|请求采用的协议名称和版本号,常为`HTTP/1.0`或`HTTP/1.1`|
|`$uri`|当前请求的`URI`,在请求过程中`URI`可能被改变,例如在内部重定向或使用索引文件时|
|`$nginx_version`|`Nginx`的版本号|
|`$pid`|`worker`进程的`PID`|
|`$pipe`|如果请求的是`HTTP`流水线(pipelined)发送的,pipe值为`p`,否则为`.`|
|`$connection_requests`|当前通过一个连接获得的请求数量|
|`$cookie_name`|`name`即`Cookie`的名字,可得到`Cooike`的信息|
|`$status`|HTTP请求状态|
|`$msec`|日志写入时间,单位为秒,精度是毫秒|
|`$time_local`|在通用日志格式下的本地时间|
|`$upstream_addr`|请求反向代理到后端服务器的IP地址|
|`$upstream_port`|请求反向代理到后端服务器的端口号|
|`$upstream_response_time`|请求后端服务器消耗的时间|
|`$upstream_status`|请求在后端服务器的`HTTP`响应状态|
|`$geoip_city`|城市名称,在`geoip`模块中使用|

##### 内置变量常见技巧
`Nginx`的内置变量主要用于日志记录和分析,以及业务逻辑的处理.
`$`
