IO
====

redis 使用epoll通信多路复用
nginx 使用epoll

操作系统内核
用户空间
内核空间
syscall 

演变关系
BIO NIO SELECT EPOLL

程序运行过程中抓取生命周期strace
```bash
strace -ff -o
```

查询java进程命令`jps`java自带,可以查看进程号
```bash
jps
```

```/proc```目录存放进程目录
```/proc/进程号/task```是该进程的所有线程
```/proc/进程号/fd```文件描述符
最基本描述符 基本流
``
0 标准输入
1 标准输出
2 错误输出
``

bash 查看网络
`
netstat -natp
`

`tail -f ooxx.进程号` tail -f 阻塞进程的线程 

线程栈是独立的，堆是共享的