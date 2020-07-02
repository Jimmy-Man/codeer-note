LNMP环境安装与搭建 Linux Nginx MySQL PHP
--------------------------------------

#### Centos安装`EPEL`扩展仓库
`EPEL`是`yum`的一个软件源,里面包含了很多基本源里没有的软件.
* 方法一`yum`安装：
```bash
yum install epel-release
```
* 方法二 手动安装
  如果方法一提示epel-release不存在则使用此方法
  + 在下面的网址寻找到对应的版本,架构软件包:
  ```bash
  https://dl.fedoraproject.org/pub/epel/
  ```
  + 在找到的源下复制该`RPM`连接，然后安装
  ```bash
  ## rpm -vih 复制的rpm连接
  rpm -vih https://dl.fedoraproject.org/pub/epel/7/x86_64/Packages/e/epel-release-7-12.noarch.rpm
  ```
  
#### 安装`Nginx`
```bash
yum install nginx
```
查看`nginx`版本
```bash
nginx -v
```
运行`nginx`
```bash
systemctl start nginx
```
开机启动`nginx`
```bash
systemctl enable nginx
```

#### 安装`MySQL`
