常用模块详解
----------


### 定制`HTTP`头信息
定制`HTTP`头信息,是实际业务中一个很重要的功能.例如将请求结果缓存在浏览器,代理到后端服务器过程生成唯一的ID进行识别等.
#### 使用`ngx_http_headers_module`设置响应头
`ngx_http_header_module`是Nginx编译时自带的模块,主要包含`add_header`和`expires`两个指令.
* ##### `expires`
语法: 
```nginx
expires [modified] time;
expires epoch|max|off;
```
默认值:
expires off;
环境: `http`、`server`、`location`、`if in location`
用途: 设置`Expires`和`Cache-Control`响应头字段,主要作用是控制缓存时间,如在浏览器上的缓存时间、CDN的缓存时间.参数值可以是正数、负数或零.
```nginx
expires -1; # 输出的响应头是cache-control: no-cache;表示不缓存
expires 1h; # 输出的响应头是cache-control: max-age=3600 表示缓存1h,max-age的单位为秒
```
高级用法,根据Content-Type的类型来定义不同的缓存时间
```nginx
map $sent_http_content_type $expires {
    default off; # 其它不缓存
    application/pdf 1h; # PDF缓存1小时
    ~image/ 10h; # image/ 缓存10小时
}
```
* ##### `add_header`
语法:
```nginx
add_header name value [aways]; #添加aways后可以在任何HTTP状态下输出响应头,否则只能在200,201,204,206,301,302,303,304,307,308时输出
```
默认值: 无
环境: `http`、`server`、`location`、`if in location`
用途: 添加自定义的响应头.

* ##### 实战经验
  + ###### `expires 1h `
`expires 1h;`表示在浏览器上缓存1小时,但要注意代码端不要缺少`Last_Modified`响应头,无法在浏览器上进行缓存.
  + ###### `add_header`
如果使用`aways`参数,即使出现异常时,也会输出`add_header`响应头信息.
注意不要建立重复的响应头，引起不必要的Bug

#### 使用`headers-more-nginx`控制请求头和响应头
`add_header`指令,只适合用来添加响应头,如需对`HTTP`请求头进行处理,可以使用第三方模块`headers-more-nginx`,它可以用来添加、删除、修改`HTTP`请求头和响应头.
##### 1. 安装
git地址:`https://github.com/openresty/headers-more-nginx-module.git`
```shell
git clone https://github.com/openresty/headers-more-nginx-module.git
cd nginx-1*/ && ./configurate --prefix=/opt/nginx --add-module=/path/to/header-more-nginx-module
make && make install
``` 
