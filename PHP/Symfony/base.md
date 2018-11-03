Symfony 笔记：

### 安装 

composer 安装
    --website-skeleton是web应用优化版
```ssh
composer create-project symfony/website-skeleton project-name
```
    -- 徽服务 或者 API
```
composer create-project symfony/skeleton peoject-name
cd project-name
composer require symfony/web-server-bundle --dev
```
-----------------------------------------------------------------
### 运行
```
cd project-name
php bin/console server:run
```
Lnmp或者LAMP环境下：
待继.....

---------------------------------------------------------------
#### 检查服务器环境要求

```
cd your-project/
composer require symfony/requirements-checker
```
该组件会在Public目录下创建check.php文件，在浏览器访问该文件排查服务器环境问题
检查完问题后记得删除该组件
```
cd your-project/
composer remove symfony/requirements-checker
```
-----------------------------------------------------------------

#### 检查安全漏洞
```
cd my-project/
composer require sensiolabs/security-checker --dev
```

#### Symfony演示应用程序
https://github.com/symfony/demo
```
composer create-project symfony/symfony-demo
```




