云服务器[阿里云、腾讯云、AWS等]使用密钥对登录
=======================================

### 登录方法 
1. 把下载的私钥修改权限
```shell
chmod 400 xxxx.pem
```
2. 登录
```sh
# ssh -i 私钥文件 user@ip
ssh -i xxx.pem root@192.168.2.5
```

### 其它用户使用私钥登录
1. 拷贝公钥文件到其它用户的.ssh目录下
```sh
# 以www作为普通用户的例子,authorized_keys作为云服务器下发的公钥文件
copy /root/.ssh/authorized_keys /home/www/.ssh
```
2. 修改公鉏文件所属用户与所属组
```sh
# 以www:www为权限组，authorized_keys为公钥文件
chown www:www /home/www/.ssh/authorized_keys
```
3. 登录
```sh
ssh -i xxx.pem www@192.168.2.5
```
