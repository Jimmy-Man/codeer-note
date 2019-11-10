Redis数据类型
============

Redis的五大数据类型
* string(字符串)
  - string 是redis最基本的类型，可以理解成跟Memcached一样的类型，一个key对应一个value
  - string 是二进制安全的。意思是redis的string可以包含任何数据。比如jpg图片或者序列化对象
  - string 类型是Redis的最基本的数据类型，一个Redis中字符串value最多可以是512M

* Hash(类似与Java里的Map)
  - Redis hash 是一个键值对集合
  - Redis hash 是一个string类型的field和value的映射表,hash特别适合存储对象。

* List(列表)
  - Redis 列表是简单的字符串列表，按照插入顺序排序，你可以添加一个元素到列表的头部(左边)或者尾部(右边)。
  - 它的底层实际是个链表

* Set(集合)
  - Set是string类型的无序集合。它是通过HashTable实现的。

* Zet(sorted set:有序集合)
  - Redis zset 和set 一样也是string类型元素的集合,且不允许重复的成员。
  - 不同的是每个元素都会关联一个double类型的分数
  - redis下是通过分数来为集合中的成员进行从小到大的排序。zset的成员是唯一的，但分数(sore)却可以重复  

-----

  String 单值单Value
  * `keys *` 
  * `eists key` 判断key是否存在
  * `move key db` 把key从当前库移到指定库
  * `expir key 秒` 为给定的key设置过期时间
  * `ttl key ` 查看还有多少秒过期，-1表示永不过期，-2表示已过期
  * `type key` 查看key的数据类型

`set/get/del/append/strlen`
`Incr/decr/incrby/decrby` 一定要是数字才能进行加减
`getrange/setrange`[获取设置指定区间范围内的值]
`setex(setwith expire) key second键秒值 value` /`setnx(set if not exist) key value `
`mset/mget/msetnx`
`getset`(先get再set)

  ---------

List 
它是一个字符串链表,left、right 都可以插入添加
如果键不存在，创建新的链表
如果键已存在，新增内容
如果值全移除，对应的键也就消失了
链表的操作无论是头和尾效率都极高，但假如是对中间元素进行操作，效率就很惨了
`lpush/rpush/lrange`
`lpop/rpop`
`lindex,按照索引下村获得元素(从上到下)`
`llen`
`lrem key count value` 删除N个Value
`ltrim key 开始index 结束index` 截取指定范围的值后再赋值给key
`rpoplpush` 源列表目的列表
`lset key index value `
`linsert key before/after 值1 值2`

-----

Set 单值多Value
`sadd/smembers/sismember`
`scard` 获取集合里面的元素个数
`srem key value` 删除集合中元素
`srandmember key number` 某个整数(随机出几个位数)
`spop key` 随机出栈
`smove key1 key2 key1里的某个值` 将key1里的某个值赋给key2
`sdiff/sinter/sunion `差集|交集|并集

------

Hash KV模式不变,但V是一个键值对
`hset/hget/hmset/hmget/hgetall/hdel`
`hlen`
`hexists key keyname`
`hkeys/hvals`
`hincrby/hincrbyfloat`
`hsetnx`

-----

Zset有序集合  
  - 在set基础上，加一个score值。key score1 v1 score2 v2
  `zadd/zrange`
  `zrangbyscore key startScore endScore [withscores|limit]`
  `zrem key value` 删除元素
  `zcard/zcount key score区间/zrank key key values值`获得下标值/zscore key 对应的值，获得分数
  `zrevrank key values值` 逆序获得下标值
  `zrevrange`
  `zrevrangebyscore key `
  


