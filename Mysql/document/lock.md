Mysql锁机制
==========

定义
* 锁是计算机协调多个进程或线程并发访问某一资源的机制 
* 在数据访问中，除传统计算资源(如CPU/RAM/IO等)的急用外，数据也是一种供许多用户共享的资源。如何保证数据并发访问的一致性、有效性是所有数据库必须解决的一个问题，锁冲突也是影响数据库并发访问性能的一个重要因素。从这个角度来说，锁对数据库而言显得尤其重要，也更复杂。

锁的分类
* 从数据操作的类型(读/写)分
  + 读锁(共享锁): 针对同一份数据，多个读操作可以同时进行而不会互相影响。
  + 写锁(排它锁): 当前写操作没有完成前，它会阻断其他写锁和读锁
* 从对数据操作的粒度分
  + 表锁
  + 行锁

* 开销、加锁速度、死锁、粒度、并发性能只能就具体应用的特点来说哪种锁更合适

三锁
* 表锁(偏读)
  + 特点
    - 偏向MyISAM存储引擎，开销小,加锁小，加锁快，无死锁，锁定粒度大，发生锁冲突的概率最高，并发度最低。
  + 使用
    - `lock talbe tableName read(write),tableName2 read(write)`
    - `show open talbes` 查看是否有被锁的表
    - `unlock tables` 释放表锁 
    
   + 表锁分析
     - 查看那些表被锁了 `show open tables`
     - 如何分析表锁定 `show status like 'table%'` 
       + 可以通过检查`table_locks_waited`和`table_locks_immediate`状态变量来分析系统上的表锁定
       + `Table_locks_immediate`: 产生表级锁定的次数，表示可以立即获取锁的查询次数，每次即获取锁值加1;
       + `Table_locks_waited`: 出现表级锁定争用而发生的等待的次数(不能立即获取锁的次数，每等待一次锁值加1)，此值高则说明存在着较严重的表级锁争用情况；
       `MyISAM的读写锁调试是写优先，这也是myisam不适合做写为主表的引擎。因为写锁后，其他线程不能做任何操作，大量的更新会使查询很难得到锁，从而造成永远阻塞`
   + 结论
     - MyISAY在执行查询语句前，会自动给涉及的所有表加读锁，在执行增删改操作前，会自动给涉及的表加写锁
     - 1.在对MyISAM表的读操作(加读锁),不会阻塞其他进程对同一表的读请求，但会阻塞对同一表的写请求。只有当读锁释放后，才会执行其它进程的写操作。
     - 2. 对MyISAM表的写操作(加写锁)，会阻塞其它进程对同一表的读和写操作，只有当写锁释放后，才会执行其它进程的读写操作。
     - `简而言之，就是读锁会阻塞写，但不会堵塞读，而写锁则会把读和写都堵塞`

* 行锁(偏写)
  + 特点
    - 偏向InnoDB存储引擎，开销大，加锁慢；会出现死锁；锁定粒度最小，发生锁冲突的概率最低，并发度也最高。
    - InnoDB与MyISAM的最大不同有2点：1是支持事务(TRANSACTION);2是采用了行级锁

       事务是由一组SQL语句组成的逻辑处理单元，事务具有以下4个属性，通常简称为事务的ACID属性。
       * 原子性(Atmicity)
       * 一致性(Consistent)
       * 隔离性(Isolation)
       * 持久性(Durable))

       并发事务处理带来的问题
         * 更新丢失(Lost Update)
         * 脏读(Dirty Reads)
         * 不可重复读(Non-Repeatable Reads)
         * 纪读(Phantom Reads)
    
       事务隔离级别

        | 隔离级别 | 读数据一致性 | 脏读 | 不可重复读 | 幻读 |
        | :-: | :-: | :-: | :-: | :-: | :-: | 
        |||||
    
    + 测试
      - 关闭自动提交 `set autocommit = 0;`

    `无索引行锁升级为表锁` 
    间隙锁的危害 
      + 当我们用范围条件而不是相等条件检索数据，并请求共享或排他锁时，Innodb会给符合条件的已有数据记录的索引项加锁；对于键值在条件范围内但并不存在的记录，叫做"间隙"(GAP);
      Innodb也会对这个"间隙"加锁，这种锁机制就是所谓的间隙锁(Next-Key)
  + 结论
    - Innodb存储引擎由于实现了行级锁定，虽然在锁定机制的实现方面所带来的性能损耗可能比表级锁定会要更高一些，但是在整体并发处理能力方面要远远优于MyISAM的表级锁定的。当系统并发发量较高的时候，Innodb的整体性能和MyISAM相比就会有比较明显的优势了。
    - 但是，InnoDB的行级锁定同样也有其脆弱的一面，当我们使用不当的时候，可能会让Innodb的整体性能表现不仅不能比MyISAM高，甚至可能会更差。
    
  + 行锁分析
    - `show status like like 'innodb_row_lock%'` 通过检查InnoDB_row_lock状态变量来分析系统上的行锁的争夺情况
      - `Innodb_row_lock_current_waits` 当前正在等待锁定的数量
      - `Innodb_row_lock_time`: 从系统启动到现在锁定总时间长度
      - `Innodb_lock_time_avg`: 每次等待所花平均时间
      - `Innodb_lock_time_max`: 从系统启动到现在等待最长的一次所花费的时间
      - `Innodb_lock_waits`: 系统启动后到现在总共等待的次数
   对于这5个状态变量，比较重要的主要是:
        innodb_row_lock_time_avg(等待平均时长)
        innodb_row_lock_waits(等待总次数)
        innodb_row_lock_time(等待总时长) 这三项。
    如果发现有问题则可以使用showprofile命令查看与分析

  + 优化建议
    - 尽可能让所有数据检索都通过索引来完成，避免无索引行锁升级为表锁
    - 合理设计索引，尽量缩小锁的范围
    - 尽可能较少检索条件，避免间隙锁
    - 尽量控制事务大小，减少锁定资源量和时间长度
    - 尽可能低级别事务隔离

* 页锁
  开销和加锁时间界于表锁和行锁之间；会出现死锁；锁定粒度界于表锁和行锁之间，并发度一般