Redis主从
=========

是什么？
* 行话: 也是是我们所说的主从复制,主机数据更新后根据配置和策略，自动同步到备机的master/salver机制，Master以写为主,Salve以读为主

能干嘛？
* 读写分离
* 容灾恢复

怎么玩
* 1.配置从(库)不配主(库)
* 2.从库配置: `slaveof 主库IP 主库端口`
  - 每次与master断开后，都需要重新连接,除非你配置进redis.conf文件
  - Info replication 查询主从信息
* 3.修改配置文件细节操作
  + 拷贝多个redis.conf文件
  + 开启daemonize yes 
  + pidfile 文件名字
  + 指定端口
  + logfile Log文件名字
  + dbfilename 备份Dump.rdb名字
* 常用3种配置
  + 一主二仆
    - Init
  + 薪火相传
    - 上一个Slave可以是下一个slave的Master,Slave同样可以接收其他slave的连接和同步请求,那么该slave作为了链条中下一个的master,可以有效减轻master的压力
    - 中途变更转向:会清除之前的数据,重新建立拷贝最新的
    - `slaveof 新主库IP 新主库端口`
  + 返客为主
    - `SLAVEOF no none` 使当前数据库停止与其他数据库的同步，转成主数据

复制原理
* Slave启动成功连接到master后会发送一个sync命令
* Master接到命令启动后台的存盘进程,同时收集所有接收到的用于修改数据集命令，在后台进程执行完毕之后，master将传送整个数据文件到slave,以完成一次完全同步
* 全量复制: 而slave服务在接收到数据库文件数据后，将其存盘并加载到内存中。
* 增量复制: Master继续将新的所有收集到的修改命令依次传给slave,完成同步,但是只要是重新连接master,一次完全同步(全量复制)将被自动执行

哨兵模式(sentinel) 返客为主自动版
* 是什么：
  + 反客为主的自动版，能够后台监控主机是否故障，如果故障了根据投票数自动将从库转换为主库
* 怎么使用：
  + 调整结构，1主多从
  + 目录下新建`sentinel.conf`文件,文件名绝不能错
  + 配置哨兵，填写内容
    - `sentinel monitor 被监控数据库名字(自己起名字) 127.0.0.1 6379 1`
    - 上面一个数字 1,表示主机挂掉后slave投票决定让谁接替成为主机，得票数为多少后成功主机
  + 启动哨兵
    - `redis-sentinel /目录/sentinel.conf`
  + 一组sentinel能同时监控多个Master

复制的缺点
* 复制延时
由于所有的写操作都是先在Master上操作，然后同步更新到Slave上，所以从Master同步到Slave机器有一定的延迟，当系统很繁忙的时候，延迟问题会更加严重，Slave机器数量的增加也会使这个问题更加严重
