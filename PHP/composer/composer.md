Compoer
-------

Composer:``是PHP的一个依赖管理工具``

### 使用Composer
解决和下载依赖:``install``
```bash
composer install
## 或者
php composer.phar install
```

### 自动加载
除了库的下载，`Composer`还准备了一个自动加载文件，它可以加载`Composer`下载的库中所有的类文件。使用它，你只需要将下面这行代码添加到你项目的引导文件中
```php
require 'vendor/autoload.php';
```

## 详细教程
[基本用法](./document/base.md)  
[库(资源包)](./document/libraries.md)  
[命令行](./document/cli.md)  
[composer.json架构](./document/schema.md)  
