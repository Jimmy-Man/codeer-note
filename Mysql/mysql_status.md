### mysql查看sql执行情况的几种方法

mysql系统变量分为全局变量和会话变量，全局变量的修改影响到整个服务器，会话变量修改只影响当前的会话。

* 1.查看log日志是否开启
```
show variables like 'general_log'
set GLOBAL general_log='ON';
SET GLOBAL general_log_file = '/tmp/mysql.log'
```
不使用的时候记得关掉，否则会大量占用磁盘空间。
* 2.show processlist命令查看了当前正在执行的sql语句,同时可以查看用户的当前连接
```
show processlist
```
* 3.查看慢日志
  - show variables like '%slow_query_log%';
  - show variables like 'long_query_time%';设置慢日志记录什么样的SQL，默认10s
  - log-queries-not-using-indexes：未使用索引的查询也被记录到慢查询日志中,一般也开启这个变量
* 4.show status查看mysql运行状态（to learn）
比如查看mysql中有多少条慢查询记录： 
```
show global status like '%Slow_queries%';
```
