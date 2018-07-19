### Python 基础　
#### Python 输入输出
``` Python
### 输入输出
name = input("Please inter your name:　")
# print("Hi",name)
```
###
```python
#ord()函数获取字符的整数表示
#chr()函数把编码转换为对应的字符
ord("a")
chr(65)
chr(25991)
```
###

#!/usr/bin/python3
#!-*- coding: UTF-8 -*-

### 输入输出
name = input("Please inter your name:　")
# print("Hi",name)
# print(len(name))

### Python 字符串和编码
#ord()函数获取字符的整数表示
#chr()函数把编码转换为对应的字符
ord("a")
chr(65)
chr(25991)

#### ASCII编码和Unicode编码的区别：ASCII编码是1个字节，而Unicode编码通常是2个字节
#### 于Python的字符串类型是str，在内存中以Unicode表示，一个字符对应若干个字节。如果要在网络上传输，或者保存到磁盘上，就需要把str变为以字节为单位的bytes
#### 以Unicode表示的str通过encode()方法可以编码为指定的bytes
### encode()
### deconde()

'ABC'.encode('ascii')
'中文'.encode('utf-8')
b'\xe4\xb8\xad\xe6\x96\x87'.decode('utf-8')
### 可以传入errors='ignore'忽略错误的字节：
b'\xe4\xb8\xad\xff'.decode('utf-8', errors='ignore')
#### 纯英文的str可以用ASCII编码为bytes，内容是一样的，含有中文的str可以用UTF-8编码为bytes。含有中文的str无法用ASCII编码，因为中文编码的范围超过了ASCII编码的范围，Python会报错。在bytes中，无法显示为ASCII字符的字节，用\x##显示。

st = "你好".encode("utf-8")
print(st)
print(len(st))
### len() 要计算str包含多少个字符
len("ABC")

### 字符格式化 %
'Hi, %s, you have $%d.' % ('Michael', 1000000)
### 或者　使用字符串的format()方法
'Hello, {0}, 成绩提升了 {1:.1f}%'.format('小明', 17.125)



## List(列表) ['',''..]   tuple(元组) ('',''...)
### Start list ==========================================================
#### list是一种有序的集合，可以随时添加和删除其中的元素。
coress = ['Englist','Matchs']

### len()函数可以获得list元素的个数
list_var = ['var1','var2']
len(list_var)
### append() 追加元素到末尾
list_var.append('var3')
### insert(index, value) 可以把元素插入到指定的位置 索引从0开始
list_var.insert(2,'var_insert2')

### 要删除list末尾的元素，用pop()方法
list_var.pop()
### 删除指定位置的元素,用pop(i)
list_var.pop(2)
### End list  ==========================================================

### Start tuple -------------------------------------------------------
###另一种有序列表叫元组：tuple。tuple和list非常类似，但是tuple一旦初始化就不能修改，元组使用小括号
tuple_var = ('Jimmy','Kimmy',['English','Match'])
### End   tuple -------------------------------------------------------

### 字典(dict) {key:value,key2:value2,...}   与集合set set()
### Start 字典dict -----------------------------------------------------------------------------------------------------------------
### Python内置了字典：dict的支持，dict全称dictionary，在其他语言中也称为map，使用键-值（key-value）存储，具有极快的查找速度。
dict_var = {"Jimmy":18,"Tim":12,"Timmy":20}
print(dict_var["Jimmy"])
#### 通过in判断key是否存在 或者　dict提供的get()方法，如果key不存在，可以返回None，或者自己指定的value
if"Jimmy" in dict_var:
    print ''
dict_var.get('Jimmy')
dict_var.get('Jimmy',11)
#### 要删除一个key，用pop(key)方法，对应的value也会从dict中删除
dict_var.pop("Jimmy")

# 和list比较，dict有以下几个特点：

# 查找和插入的速度极快，不会随着key的增加而变慢；
# 需要占用大量的内存，内存浪费多。
# 而list相反：

# 查找和插入的时间随着元素的增加而增加；
# 占用空间小，浪费内存很少。
# 所以，dict是用空间来换取时间的一种方法。

# dict可以用在需要高速查找的很多地方，在Python代码中几乎无处不在，正确使用dict非常重要，需要牢记的第一条就是dict的key必须是不可变对象。

### End 字典dict -----------------------------------------------------------------------------------------------------------------

### Start　set集合
#### set和dict类似，也是一组key的集合，但不存储value。由于key不能重复，所以，在set中，没有重复的key。
set_var  = set()
set_var  = {'Jimmy','Kimmy','Timmy','Jack'}

### 通过add(key)方法可以添加元素到set中，可以重复添加，但不会有效果
set_var.add('toney')
### 通过remove(key)方法可以删除元素
set_var.remove('toney')
###set可以看成数学意义上的无序和无重复元素的集合，因此，两个set可以做数学意义上的交集、并集等操作
s1 = set([1, 2, 3])
s2 = set([2, 3, 4])
s1 & s2  ### 交集
s1 | s2 ### 并集
### set和dict的唯一区别仅在于没有存储对应的value，但是，set的原理和dict一样，所以，同样不可以放入可变对象，因为无法判断两个可变对象是否相等，也就无法保证set内部“不会有重复元素”。试试把list放入set，看看是否会报错

### 条件判断
# if :
# elif:
# elif:
# else:

### 循环 for ... in
#### for...in
for name in tuple_var:
    print(name)
### Python提供一个range()函数，可以生成一个整数序列，再通过list()函数可以转换为list。比如range(5)生成的序列是从0开始小于5的整数 
list(range(101))
### while ...:
while 1:
    print 'hi'

### break 终止循环　continute 语句可以提前结束本轮循环，并直接开始下一轮循环



   





