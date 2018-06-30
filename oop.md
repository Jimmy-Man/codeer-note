# 面向对象概念

## Python语法：

### 定义类
``` Python
class ClassName:
  ##构造函数 
  def __init__ (self,[,args...]):
  
  
  
  ## 析构函数 [垃圾回收机制] 
  def __del__ (self,[,args...]):
  
```
### 定义类的属性
```
  1、直接在类中定义
  2、在构造函数中定义
```
Python 没有访问控制 没有提供私有属性的功能(全靠自觉[规范])
```python
    def __init__(self,name,age,height):
        self.name = name    ##没有下划线  public 可以公开访问
        self._age  = age    ## 一条下划线 protect
        self.__height = height  ## 二条下划线    private 私有属性(其实依然能从外部访问)
```
### 定义类的方法
```python
    ## 定义函数
    def methodName(selfe):  ## 公开方法
        return 'public'
    def _methodNmae2(self):
        return 'protect'
    def __methodName3(self):
        return 'private'
```
### 类的继承与多态
  ```
  继承 子类能继承分类全部功能
  多态 
  多重继续
  
  ```
