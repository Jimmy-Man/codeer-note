Vue应用部署
==========
* 修改配置文件 
```

```
* 打包项目
 ```
 npm run build
 ```
* Nginx 代理配置
```conf
    #gzip  on;

    server {
        listen       80;
        server_name  localhost;

        #charset koi8-r;

        #access_log  logs/host.access.log  main;

        root /data/www/aa.com/dist;#react/vue项目的打包后的dist
        # index index.html index.htm;
 
        # location / {
        #         try_files $uri $uri/ /index.html;
        # }

       #代理后台接口
         location /api/ {
             proxy_pass http://127.0.0.1:9080;#转发请求的地址
             proxy_connect_timeout 6000;#链接超时设置
             proxy_read_timeout 6000;#访问接口超时设置
       }
        #error_page  404              /404.html;

        # redirect server error pages to the static page /50x.html
        #
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }
```
* 