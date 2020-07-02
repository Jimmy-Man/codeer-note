MySQL函数

#### 常用函数
1. 字符函数
   - concat:连接
   - substr:截取子串
   - upper:变大写
   - lower: 变小写
   - replace:替换同
   - length: 获取字节长度
   - trim: 去前后空格
   - lpad: 左填充
   - rpad: 右填充
   - instr: 获取第一次出现的索引
2. 数学函数
   - ceil: 向上取整
   - round: 四舍五入
   - mod: 取模
   - floor:向下取整
   - trancate:截断
   - rand:获取随机数,返回0-1之间的小数
3. 日期函数
   - now : 返回当前日期+时间
   - year : 返回年
   - month : 返回月
   - day : 返回日
   - date_format : 将日期转换为字符
   - curdate : 返回当前日期
   - str_to_date : 将字符转换为日期
   - curtime : 返回当前时间
   - hour : 小时
   - minute : 分
   - second : 秒
   - datediff : 返回2个日期相差的天数
   - monthname : 以英文形式返回月
4. 其它函数
   - version : 当前数据库服务器版本
   - database : 当前打开的数据库
   - user : 当前用户
   - `password('字符')` : 返回该字符的加密字符[已被弃用]
   - md5('字符') : 返回字符的MD5加密字符
5. 流程控制函数
   - `if (条件表达式，表达式1，表达式2)` : 如果表达式成立，返回表达式1，否则返回表达式2
   - case情况1 : 
     ```sql
     case 变量|表达式|字段
     when 常量1 then 值1
     when 常量2 then 值2
     ...
     else 值n
     end 
     ```
   - case情况2 : 
     ```sql
     case 
     when 条件1 then 值1
     when 条件2 then 值2
     ...
     else 值n
     end 
     ```
     
#### 分组函数 
   1. 常用函数
     - max : 最大值
     - min : 最小值
     - sum : 和
     - avg : 平均值
     - count : 计算个数
   2. 特点
      1. 语法 
         select max(字段) from 表名;
      2. 支持的类型
         sum和avg一般用于处理数值
         max,min,count 可以处理任何数据类型
      3. 以上分组函数都忽略`null`
      4. 都可以搭配`distinct`使用,实现去重的统计`select sum( distinct 字段) from 表名`
      5. count()
         count(字段) : 统计该字段非空值的个数
         count(*)  : 统计结果集的行数
         count(1) : 统计结果集的行数
         效率上: MyISAM存储引擎，count(*)效率最高
                Innodb存储引擎，count(*)和count(1)的效率 > count(字段)
      6. 和分组函数一同查询的字段,要求是group by后出现的字段
    