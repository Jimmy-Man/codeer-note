Redis笔记
========

NoSQL Not only SQL这些数据类型的数据存储不需要固定的模式，无需多余操作就可以横向扩展。
数据间无关系，这就非常容易扩展。

NoSQL数据库都具有非常高的读写性能，尤其在大数据量下，同样表现优秀
这得得益于它的无关系性，数据库结构简单
1秒 读10W次 写8W次

---------------
RDBMS VS NoSQL
---
RDBMS
- 高度组织化结构化数据
- 结构化查询语言(SQL)
- 数据和关系都存储在单独的表中
- 数据操纵语言，数据定义语言
- 严格的一致性
- 基础事务

ER图(1:1/1:N/N:N,主外键等常见)

NoSQL
- 分区表着不仅仅是SQL
- 没有声明性查询语言
- 没有预定义的模式
- 键 - 值对存储，列存储，文档存储，图形数据库
- 最终一致性，而非ACID属性
- 非结构化和不可预知的数据
- CAP定理
- 高性能，高可用性和可伸缩性

NoSQL数据模型 - 聚合模型
  * KV键值
  * Bson
  * 列族
  * 图形

BSON
---------------------
KV+Cache+Persistence
---
大数据时代的3V
- 海量Volume
- 多样Variety
- 实时Velocity

大数据时代的3高
- 高并发
- 高可扩
- 高性能

NoSQL数据库的四大分类
* KV键值[redis+tair,BerkeleyDB+Redis,memcache+redis]
* 文档型数据库(bson格式比较多)[CouchDB,MongoDB]
* 列存储数据库[Cassandra,HBase,分布式文件系统]
* 图关系数据库[它放的不是图形的，放的是关系比如:朋友圈社交网络、广告推荐系统|Neo4J,InfoGrid]

传统ACID
  * A(Automicity)原子性
  * C(Consistency)一致性
  * I(Isolation)独立性
  * D(Durability)持久性

CAP
  * C: Consistency 强一致性
  * A: Availability (可用性)
  * P: Partition tolerance (分区容错性)

CAP原理 CAP+BASE
===============
CAP理论的核心是:`一个分布
式系统不可能同时很好的满足一致性、可用性和分区容错性这三个需求，最多只能同时较好的满足两个。`
根据CAP原理将NoSQL数据库分成了满足CA原则、满足CP原则和满足AP原则三大类:
  * CA - `单点集群，满足一致性，可用性的系统，通常在可扩展上不太强大` RDBMS关系型数据库
  * CP - `满足强一致性，分区容错性的系统，通常性能不是特别高` MongoDB,HBase,Redis
  * AP - `满足可用性，分区容错性的系统，通常可能对一致性要求低一些` CouchDB,Cassandra,DynamoDB

  由于当前网络硬件背景肯定会出现延迟丢包等问题，所以
  `P 分区容忍性是我们必须需要实现的`所以我们只能在C一致性和A可用性之间权衡，没有NoSQL能同时保证这三点
  
  CA 传统关系型数据库MySQL,Oracle
  AP 大多数网站架构的选择
  CP Redis、MongoDB

BASE
BASE: `就是为了解决关系数据库强一致性引起的问题而引起的可用性降低而提出的解决方案`
  * 基本可用 (Basically Available)
  * 软状态 (Soft state)
  * 最终一致性 (Eventually consistent)
它的思想是通过让系统放松对某一时刻数据一致性的要求来换取系统整体伸缩性和性能上改观。

分布式和集群:
* 分布式: `不同的多台服务器上面部署不同的服务模块(工程),他们之间通过Rpc/Rmi之间通信和调用，对外提供服务和组内协作`
* 集群: `不同的多台服务器上面部署相同的服务模块，通过分布式调度软件进行统一的调度，对外提供服务和访问`

[入门概述](./document/Overview.md)
[Redis数据类型](./document/dataType.md)
[Redis配置文件](./document/configurate.md)
[Redis持久化](./document/persistence.md)
[Redis事务](./document/transaction.md)
[Redis发布订阅](./document/subscription.md)
[Redis主从配置](./document/masterSlave.md)