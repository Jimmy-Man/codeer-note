MySQL主从复制
============

复制基本原理
  * slave会从master读取binlog来进行数据同步
  * 三步骤
    + 1. master将改变记录到二进制日志(binary log).这些记录过程叫做二进制日志事件，binary log events;
    * 2. slave将master的binary log events拷贝到它的中继日志(relay log)
    * 3. slave重做中继日志中的事件，将改变应用到自己的数据库中。Mysql复制是异步的且串行化的。

复制的基本原则 
* 每个slave只有一个master
* 每个slave只能有一个唯一的服务器ID
* 每个master可以有多个slave

复制最大的问题
* 延时

一主一从常见配置
* mysql版本一致且后台以服务运行
* 主从都配置在[mysqld]结点下，都是小写
* 同一网段
* Master主机修改my.cnf/my.ini配置文件
  + [必须]主服务器唯一ID: `server-id=1`
  + [必须]启用二进制日志: `log-bin=/mypath/mysqlbin`
  + [可选]启用错误日志: `log-err=/mypath/mysqlerr`
  + [可选]根目录: `basedir="自己定义的目录地址"`
  + [可选]临时目录: `tmpdir="自己定义的目录"`
  + [可选]数据目录: `datadir="自己定义的目录/Data"`
  + 主机，读写都可以: `read-only=0`
  + [可选]设置不要复制的数据库: `binlog-ignore-db=mysql`
  + [可选]设置需要复制的数据: `binlog-do-db=需要复制的主数据库的名字`

* Slave从机修改my.cnf/my.ini配置文件
  + [必须]从服务器唯一ID
  + [可选]启用二进制文件
* 注意: 因为修改过配置文件，主机+从机都需要重启后台MySQL服务

* 主机上建立帐户并授权slave
  + `GRANT REPLICATION SLAVE ON *.* TO 'dbuser'@'从机器数据库IP' IDENTIFIED BY 'dbUserPWD';`
  + `flush privileges;` 刷新
  + 查询master的状态 `show master status;`
    -  记录下`File`和`Position`的值

* 从机上配置需要复制的主机
  + `
  CHANGE MASTER TO MASTER_HOST='主机IP地址',MASTER_USER='dbuser',
  MASTER_PASSWORD='dbUserPWD',
  MASTER_LOG_FILE='主机File名字',
  MASTER_LOG_POS=主机Postion数字
  `
  + 启动从服务器复制功能: `start slave`
  * 检查是否成功:
    - `show slave status\G;`
    - Slave_IO_Running: Yes;
    - Slave_SQL_Running: Yes;
    - 以上2个参数都是yes，则说明主从配置成功

* 如何停止从服务复制功能: `stop slave;`

