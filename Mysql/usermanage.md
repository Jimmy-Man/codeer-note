Mysql创建数据库与用户并授予所有权限
==============================

### 一、以能管理用户和创建数据库的角色或者Root用户登录Mysql
```mysql
mysql -u root -p*******
```
注意-p参数后的密码没有空格

### 二、创建数据库
```mysql
create schema [数据库名称] default character set utf8 collate utf8_general_ci; 
```
或者
```mysql
create database [数据库名称] default character set utf8 collate utf8_general_ci; 
```
例如：
```mysql
create database `my_shop` default character set utf8 collate utf8_general_ci; 
```

### 三、创建用户
命令:
```mysql
CREATE USER '[用户名]'@'host' IDENTIFIED BY '[密码]';
```
  host: localhost 本机可用  %：所有服务器
例子：
```mysql 
CREATE USER 'user1'@'localhost' IDENTIFIED BY '123456';
create user `opencart`@`localhost` IDENTIFIED BY 'Opencartpwd!123';
```
注意密码需要大于8位并有大小写加特殊字符


### 四、授予权限
命令：
```mysql
GRANT privileges ON `databasename`.tablename TO 'username'@'host';
```
说明:
1. privileges：用户的操作权限，如SELECT，INSERT，UPDATE等，如果要授予所的权限则使用ALL
2. databasename：数据库名
3. tablename：表名，如果要授予该用户对所有数据库和表的相应操作权限则可用*表示，如*.*

例子：
```mysql
GRANT SELECT, INSERT ON my_shop.product TO 'user1'@'localhost';
GRANT ALL ON *.* TO 'user1'@'%';
GRANT ALL ON my_shop.* TO 'user1'@'localhost';
```


