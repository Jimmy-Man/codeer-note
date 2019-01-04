## Linux系统下挂载硬盘

#### 1.1查看新磁盘

```
[root@cgsl ]# fdisk –l
```
找到新添加的磁盘的编号为/dev/vdc

#### 1.2硬盘分区  进入fdisk模式

```
[root@cgsl ]# /sbin/fdisk /dev/vdc
```
输入n进行分区

```
[root@cgsl ]# Command (m for help): n
```
选择分区类型

```
[root@cgsl ]# Select (default p): p
```
选择分区个数

```
[root@cgsl ]# Partition number (1-4, default 1): 1
```
一直回车，最后输入q退出。

#### 三、格式化分区  将新分区格式化为ext3文件系统

```
[root@cgsl ]# mkfs -t ext3 /dev/vdc
```
#### 四、挂载硬盘 

 1.创建挂载点，在根目录下创建storage目录

``
[root@cgsl ]# mkdir /storage
```
2.将/dev/vdc挂载到/storage下

```
[root@cgsl ]# mount /dev/vdc /storage
```
3.设置开机启动自动挂载  新创建的分区不能开机自动挂载，每次重启机器都要手动挂载。设置开机自动挂载需要修改/etc/fstab文件

```
#vi /etc/fstab
```
在文件的最后增加一行  /dev/vdc /storage ext3 defaults 1 2

注意要查询一下自己的硬盘格式，用df -T查询。
