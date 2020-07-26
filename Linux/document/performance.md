
## 硬件及操作系统层监控工具
* vmstat
* sar

### Linux 下的性能分析工具
* top , free ,ps , df 
* sysstat (sar,mpstat,iostat) \dstat \iotop
网络工具 * netsata 、 ethstatus 、 arping 
其它工具
* perf top/perf report/perf record 用来记录、报告、查看性能瓶颈
* pstack 分析Mysql内部堆栈调用



#### free 
free -h 命令后 
used - cached(availbel可再分配内存) 的值如果太大，即有内存涉漏
当发现当前系统已经使用到了`swap`,说明内存长期或曾经不够用,当使用`top`看到system选项比较高

#### iostat 
```shell
iostat -dmx 1
```
* await的值超过5就比较严重
* svctm的值超过0.5比较严重
* util超小越好

#### vmstat
```shell
vmstat -S m 1
```
* r
* b
* 
