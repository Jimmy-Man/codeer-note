RPM安装
=======

### 查询
* `rpm -qa`  `rpm -qa | grep lib` 


### 安装
* `rpm -ivh ****.rpm`

### 升级
* `rpm -Uvh ***.rpm`

### 刷新
* `npm -Fvh ***.rpm`

### 删除
* `rpm -e ****.rpm`

### 修改RPM数据库
* `rpm -rebuilddb`

----

YUM安装
======

## 设置本地YUM源

* `yum clean all`
* `yum makecache all`
* 