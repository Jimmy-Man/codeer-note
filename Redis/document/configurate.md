Redis配置文件解析
==============
redis.conf

位置：/usr/location

Units单位
* 配置大小单位，开头定义了一些基本的度量单位，只支持bytes不支持bit
* 对大小写不敏感

INCLUDES
* bind 

GENERAL
* daemonize yes
* tcp-backlog
  - 设置tcp的backlog,backlog其实是一个连接队列,backlog队列总和=未完成3次握手队列+已经完成3次握手队列。
  - 在高并发环境下需要一个高backlog值来避免慢客户端连接问题，注意linux内核会将这个值减小到/proc/sys/net/core/somaxcoon的值，所有需要确认增大somaxconn和tcp_max_syn_backlog两个值来达到想要的效果
* timeout
* Tcp-keepalive
  - 单位为秒，如果设置为0，则不会进行Keepalive检测,建议设置成60
* LogLevel 日志级别
* Logfile "" 保存日志的文件名
* Syslog-enalbe 是否把日志输出到syslog中 默认no
* Syslog-ident 指定syslog里的日志标志
* Syslog-facility 指定syslog设备,值可以是USER或LOCAL0-LOCAL7
* DataBse 16 数据库

SNAPSHOTTING快照
* Save 秒钟 写操作次数
  - RDB是整个内存的压缩过的Snapshot,RDB的数据结构，可以配置复合的快照触发条件 默认:
    + 1分钟内改了1万次
    + 或5分钟内改了10次
    + 或15分钟内改了1次
  - 禁用 如果想禁用RDB持久化的策略，只要不设置任何save指令,或者给save传入一个空字符串参数也可
* Stop-writes-on-bgsave-error
  - `stop-writes-on-bgsave-error yes`
  - 如果配置成no,表示你不在乎数据不一致或者有其它手段发现和控制
* rdbcompression
  - `rdbcompression yes`
  - rdbcompression:对于存储到硬盘中的快照，可以设置是否进行压缩存储。如果是的话，redis会采用LZF算法进行压缩。如果你不想消耗CPU来进行压缩的话，可以设置为关闭此功能。
* rdbchecksum
  - `rdbchecksum yes`
  - rdbchecksum: 在存储快照后，还可以让redis使用CRC64算法进行数据检验，但是这样做会增加大约10%的性能消耗，如果希望获取到最大的性能提升，可以关闭此功能。


REPLICATION复制

SECURITY安全
* `config get requirepass` 查看密码
* `config set requirepass ""`设置密码
* `auth 密码` 

LIMITS限制
* Maxclients
* Maxmemory
* Maxmemory-policy
  - Volatile-lru: 使用LRU算法移除KEY,只对设置了过期时间的键
  - Allkeys-lru: 使用LRU算法移除KEY
  - Volatile-random: 在过期集合中移除随机的key,只对设置了过期时间的键
  - Allkeys-random: 移除随机的key
  - Volatile-ttl: 移除那些TTL值最小的key,即那些最近要过期的key
  - noeviction: 不进行移除[永不移除]。针对写操作,只是返回错误信息
* Maxmemory-sampless
  - 设置样本数量，LRU算法和最小TTL算法都并非是精确的算法，而是估算值，所以你可以设置样本的大小，redis会默认检查这么多个key并选择期中LRU的那个
* APPEND ONLY MODE追加

### 配置redis外网可以访问
redis 默认只可以本机访问,只绑定了127.0.0.1，想要设置外网访问主要设置`redis.config`的`bind`参数
注意:`bind表示的是指定本机可以接受连接的网卡地址并不是外部服务器的IP`

如何确定
