Composer 笔记
============

### composer问题：
#### composer update或者install时，报超过内存限制错误的解决方法：
 错误：
 ```php
 Allowed memory size of 536870912 bytes exhausted….
 ```
 解决方法：
 
 1. 修改php.ini中 memory_limit 配置。 
可以通过这个如下命令查看设置是否生效：
```shell
php -r “echo ini_get(‘memory_limit’).PHP_EOL
```

 2. 使用命令：
```shell
### update
php -d memory_limit=-1 `which composer` update -vvv
### install
php -d memory_limit=-1 `which composer` install -vvv
```
也可以指定固定的大小,只要把memory_limit的值替换为想要设置的值则可。
