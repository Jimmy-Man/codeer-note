MySQL
=====

#### 查找分析慢查询

* 慢查询工具：`pt-query-digest`
* mysql慢查询日志: `可开启mysql的慢查询记录日志功能[不建议]`
* show profile: `set profiling=1;开启，服务器上执行的所有语句会检测消耗的时间,存到临时表中`
* show profile操作: `show profiles; show profile for query [query_id临时表ID]`
* 使用show status: `show status 会返回一些读数器，show global status 查看服务级别的所有计数 [这些计数可以猜测出哪些操作代价较高或者消耗时间多]`
* show processlist: `show processlist 观察是否有大量线程处于不正常的状态或者特征`
* explain [desc]: `explain 分析单条SQL语句`
* mysqldumpslow 
* 全局查询日志
  + 配置与启用 
    - 在mysql的my.cnf中，配置如下:
    - `general_log =1 `开启
    - `general_log_file=/path/logfile` 记录日志文件的路径
    - `log_output=FILE` 输出格式
  + 编码启用
    - `set global general_log =1`
    - `set global log_output='Table'`
    - 设置后，你所编写的SQL语句，将会记录到mysql库里的general_log表,可以用以下命令查看
    - `select * from mysql.general_log`
  + 永远不要在生产环境开启这个功能

---
#### 分区表
    工作原理: `对用户而言,分区表是一个独立的逻辑表,但是底层MySQL将其分成了多个物理子表,这对用户来说是透明的,每一个分区表都会使用一个独立的表文件`

    `创建表时使用"partition by"子句定义每个分区存放的数据,执行查询时,优化器会根据分区定义过滤那些没有我们需要数据的分区,这样查询只需要查询所需数据所在的分区即可`

#### SQL查询安全方案
1. 使用预处理防止SQL注入
2. 写入数据库的数据要进行特殊字符转义
3. 查询错误信息不要返回给用户,将错误记录到日志

---
##### 安全设置
* 定期做数据备份
* 不给查询用户root权限,合理分配权限
* 关闭远程访问数据库的权限
* 修改root口命,不使用默认口命,使用较复杂的口命
* 删除多余的用户
* 修改root的用户名称
* 限制一般用户浏览其他库
* 限制用户对数据文件的访问权限

##### 备份与恢复
* 备份工具: `XtraBackup` 、`mysqldump`

###数据引擎
Percona
xtradb
[MySQL索引优化](./document/index.md)
[MySQL锁机制](./document/lock.md)
[Mysql主从复制](./document/masterSlave.md)

VPN
====
virtual private network
vpn 通信标准类型
* IPsec
* PPTP
* L2TP
* OpenVPN
* SSLVPN

PPTP 和L2TP最普遍
winsows 内置L2TP VPN服务器功能

---
Windows下操作拔号连接与VPN连接的命令
* rasphone
* rasdial


211.154.155.29:19500/vpn/getlist

