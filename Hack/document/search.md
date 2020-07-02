

WEB 安全
* 客户端
  + XSS
  + CSRF
  + 点击劫持
  + URL跳转
* 服务器端
  + SQL注入
  + 命令注入
  + 文件操作类


### 常见类型
* 钓鱼Phishing
* 篡改(Tampering)
* 暗链(Hidden hyperlinks)
* Webshell


### 被黑网站常见关键词
* `Hacked by`

### 搜索引擎语法
* `Intitle:keyword` 标题中含有关键词的网页
* `Intext:keyword` 正文中含有关键词的网页
* `Site:domain`在某个域名和子域名下的网页


### CSRF(Cross-site request forgery)跨站请求伪造
概念：利用用户已经登录的身份，在用户毫不知情的情况下，以用户的名义完成非法操作。
* 危害

### URL跳转
定义:``借助未验证的URL跳转，将应用程序引导到不安全的第三方区域，从而导致的安全问题``
* 实现方式
  + Header头跳转
  + Javascript跳转
  + META标签跳转 `<meta http-equiv="Refresh" content="5; url="http://test.com" />`
  
### 点击劫持
通过覆盖不可见的框架误导受害者点击而造成的攻击行为。

* 隐蔽性较高
* 骗取用户操作
* "UI-覆盖攻击"
* 利用iframe或者其它标签的属性


### SQL注入
* 万能密码(Universal password)
   把用户名改成` 用户名'-- `(注意`--后面有空格`)
   原理把SQL拼接语句修改为注释密码后面部分
   ```sql
   select name from user where name = 'admin' -- ' and passowrd = 'djsjflkdjsflkjdsjf' limit 1
   ```
* SQL注入原理
  + SQL Injection 是一种常见的Web安全漏洞，攻击者利用这个漏洞，可以访问或修改数据，或者利用潜在的数据库漏洞进行攻击

* 过程
    1. 获取用户请求参数
    2. 拼接到代码当中
    3. SQL语句按照我们构造参数的语义执行成功
* 必备条件
    1. 可以控制输入的数据
    2. 服务器要执行的代码拼接了控制的数据

* 危害
  + 获取管理员用户名和密码
  + 获取敏感信息
  + 整个数据库：脱库
  + 获取服务器权限
  + 植入webshell,获取服务器后门
  + 读取服务器敏感文件
  + 万能密码
  + ...

### 命令注入
WEB应用如何命令注入
* 调用可执行系统命令的函数
* 函数或函数的参数可控
* 拼接注入命令



* 常见文件操作
  + 文件上传
    - 上传Wellshell
    - 上传木马
  + 文件下载
    - 下载系统任意文件
    - 下载程序代码
  + 文件包含漏洞
    - 本地文件包含
    - 远程文件包含