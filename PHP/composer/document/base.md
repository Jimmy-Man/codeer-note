composer基本用法
--------------

### 安装
* 下载`composer.phar`可执行文件
* 检查`composer`是否正常工作
```bash
php composer.phar
## 或
composer
```

### `composer.json`项目安装
要开始在你的项目中使用`Composer`，你只需要一个`composer.json`文件。该文件包含了项目的依赖和其它的一些元数据。

### 关于`require` Key
`composer.json`文件中指定`require`的值,即告诉`composer`项目需要依赖那些包
```json
// composer.json
{
    "require": {
        "monolog/monolog": "1.0.*",
        "包名称":"包版本"
    }
}
```
`require`需要一个``包名称``(例如 monolog/monolog)映射到`包版本` （例如 1.0.*)的对象。
* 包名称:``包名称由供应商名称和其项目名称构成``
* 包版本 

|名称|实例|描述|
| -- | -- | -- |
|确切的版本号| 1.0.2 | 指定包的确切版本 |
| 范围 | `>=1.0`<br>`>=1.0,<=2.0`<br>`>=1.0,<1.1|>=1.2`  | 通过使用比较操作符可以指定有效的版本范围。有效的运算符：>、>=、<、<=、!=。你可以定义多个范围，用逗号隔开，这将被视为一个逻辑AND处理。一个管道符号|将作为逻辑OR处理。AND 的优先级高于 OR。|
| 通配符 | `1.0*` | 你可以使用通配符*来指定一种模式。1.0.*与>=1.0,<1.1是等效的。 |
| `~`赋值运算符 | `~1.2` | 这对于遵循语义化版本号的项目非常有用。~1.2相当于>=1.2,<2.0。想要了解更多，请阅读下一小节。 |

* 下一重要版本(`~`波浪号运算符)  
  对于遵循`语义化版本号`的项目最有用  
  一个常见的用法是标记你所依赖的最低版本，像 ~1.2 （允许1.2以上的任何版本，但不包括2.0）  
  你还会看到它的另一种用法，使用 ~ 指定最低版本  
* 稳定性
  默认情况下只有稳定的发行版才会被考虑在内。如果你也想获得 RC、beta、alpha 或 dev 版本，你可以使用`稳定标志` `@`。你可以对所有的包做 最小稳定性 设置，而不是每个依赖逐一设置。
* 安装依赖包
  ```bash
  php composer.phar install
  ```
  包会安装在`vendor`下。
  `install`命令将创建一个`composer.lock`文件到你项目的根目录中。

  ### `composer.lock` - 锁文件
  在安装依赖后`Composer`将把安装时确切的版本号列表写入`composer.lock`文件  
  *提交版本库时，`composer.lock`与`composer.json`文件也需要提交*  
  这是非常重要的，因为`install`命令将会检查锁文件是否存在，如果存在，它将下载指定的版本（忽略`composer.json`文件中的定义）。  
  如果不存在`composer.lock`文件`Composer`将读取`composer.json`并创建锁文件。
 * 更新,[由于`composer.lock`文件的存在,如果项目所使用的依赖有新的版本，也不会获取任何更新]  
   如果需要更新使用`update`命令:
   ```bash
   php composer.phar update
   ```
   这将获取最新匹配的版本(根据你的`composer.json`文件)并将新版本更新`composer.lock`进锁文件  
   只想安装或更新一个依赖
   ```bash
   php composer.phar update monolog/monolog [...]
   ```
### Packagist
[Packagist](https://packagist.org/)是`Composer`的主要资源库。 一个`Composer` 的库基本上是一个包的源：记录了可以得到包的地方。  
Packagist 的目标是成为大家使用库资源的中央存储平台  

### 自动加载
对于库的自动加载信息，`Composer`生成了一个`vendor/autoload.php`文件。你可以简单的引入这个文件，你会得到一个免费的自动加载支持。
```php
require 'vendor/autoload.php';
```
你可以在`composer.json`的`autoload`字段中增加自己的`autoloader`:
```json
{
    "autoload": {
        "psr-4": {"Acme\\": "src/"}
    }
}
```
`Composer`将注册一个`PSR-4 autoloader`到`Acme`命名空间。  
你可以定义一个从命名空间到目录的映射。此时`src`会在你项目的根目录，与`vendor`文件夹同级。例如`src/Foo.php`文件应该包含`Acme\Foo`类。  

添加`autoload`字段后，你应该再次运行`install`命令来生成`vendor/autoload.php`文件。  
引用这个文件也将返回`autoloader`的实例，你可以将包含调用的返回值存储在变量中，并添加更多的命名空间。  
这对于在一个测试套件中自动加载类文件是非常有用的，例如:
```php
$loader = require 'vendor/autoload.php';
$loader->add('Acme\\Test\\', __DIR__);
```
除了`PSR-4`自动加载`classmap`也是支持的。这允许类被自动加载，即使不符合`PSR-0`规范。详细请查看[自动加载-参考](https://docs.phpcomposer.com/04-schema.html#autoload)。  
注意： `Composer`提供了自己的`autoloader`。如果你不想使用它，你可以仅仅引入 `vendor/composer/autoload_*.php`文件，它返回一个关联数组，你可以通过这个关联数组配置自己的`autoloader`。