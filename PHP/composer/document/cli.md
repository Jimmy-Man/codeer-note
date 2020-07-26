命令行
-----


### `composer`常用命令
|命令|描述|
|--|--|
|`composer list`|获取帮助信息|
|`composer init`|以交互的方式创建`composer.json`文件信息|
|`composer install`|从当前目录读取`composer.json`文件,处理依赖关系,并安装到`vendor`目录下|
|`composer update`|获取依赖的最新版本,升级`composer.lock`文件|
|`composer require`|添加新的依赖包到`composer.json`文件,并执行更新|
|`compoesr search`|在当前目录中搜索依赖包|
|`composer show`|列举所有可用的依赖包|
|`composer validate`|检测`composer.json`文件是否有效|
|`composer self-update`|将`composer`工具更新到最新版本|
|`composer create-project`|基于`composer`创建一个新的项目|

获取帮助信息:``composer``或``composeer list ``命令，然后结合``--help``命令来获得更多的帮助信息。

### 全局参数
下列参数可与每一个命令结合使用
* **--verbose**(`-v`)增加反馈信息的详细程度
  + -v 表示正常输出
  + -vv 表示更详细的输出
  + -vvv 则是为了debug
* **--help** (`-h`):显示帮助信息
* **--quiet** (`-q`):禁止输出任何信息
* **--no-interaction** (`-n`):不要询问任何交互问题
* **--working-dir** (`-d`): 如果指定,使用指定的目录作为工作目录
* **--profile**:显示时间和内存使用信息
* **--ansi**: 强制`ANSI`输出
* **--no-ansi**:关闭`ANSI`输出
* **--version** (`-V`)显示`composer`版本信息

### 进程退出代码
* 0:正常
* 1:能用/未知错误
* 2:依赖关系处理错误

### 初始化`init`:以交互的方式创建`composer.json`
当你运行该命令,它会以交互的方式要求你填写一些信息，同时聪明的使用一些默认值
```bash
php composer.phar init
```
#### 初始化参数
* **--name**: 包的名称
* **--description**: 包的描述
* **--author**: 包的作者
* **--homepage**: 包的主页
* **--require**: 需要依赖的其它包，必须要有一个版本约束。并且应遵循`foo/bar:1.0.0`这样的格式
* **--require-dev**:开发版的依赖包,内容格式与`require`相同
* **--stability(`-s`)** `mininum-stability`字段的值

### 安装`install`
`install`命令从当前目录读取`composer.lock`如果没有则读取`composer.json`文件,处理依赖关系，并把其安装到`vendor`目录下.
#### 安装参数
* **--prefer-source**:下载包的方式有两种:`source`和`dist`。对于稳定版本`composer`默认使用`dist`方式。而`source`则表示版本控制源。如果`prefer-source`被启用，composer将使用`source`安装。
* **--prefer-dist**:与`--prefer-source`相反,composer将尽可能从`dist`获取,这将大幅度的加快在`build servers`上的安装
* **--dry-run**:如果你只是想演示而并非实际安装一个包，你可以运行`--dry-run`命令，它将模拟安装并显示将会发生什么。
* **--dev**:安装`require-dev`字段中列出的包(这是一个默认值)
* **--no-dev**:跳过`require-dev`字段中列出的包。
* **--no-scripts**:跳过`composer.json`文件中定义的脚本
* **--no-plugins**:关闭`plugins`
* **--no-progress**:移除进度信息
* **--optimize-autoloader(-o)**:转换PSR-0/4 autoloading到`classmap`可以获得更快的加载支持。生产环境建议这么做。

## 更新 `update`
获取依赖的最新版本，并且升级`composer.lock`文件
```bash
#更新所有包
php composer.phar update
## 更新部分包
php composer.phar update vendor/package vendor/package2
## 通配符进行批量更新：
php composer.phar update vendor/*
```
### 更新参数
* 所有`install`命令的参数都可以在`update`下使用
* **--lock**:仅更新`lock`文件的`hash`，取消有关 lock 文件过时的警告。
* **--with-dependencies**:同时更新白名单内包的依赖关系，这将进行递归更新。

## 申明(新增)依赖 `require`
`require`命令增加新的依赖包到当前目录的`composer.json`文件中。
```bash
## 交互的方式添加依赖
php composer.phar require
## 直接指明依赖
php composer.phar require vendor/package:2.* vendor/package2:dev-master
```
### 申明(新增)依赖参数
* **--prefer-source**:当有可用的包时，从`source`安装。
* **--prefer-dist**:当有可用的包时，从`dist`安装。
* **--dev**:安装`require-dev`字段列出的包。
* **--no-update**:禁用依赖关系的自动更新。
* **--no-progress**:移除进度信息，这可以避免一些不处理换行的终端或脚本出现混乱的显示。
* **--update-with-dependencies** 一并更新新装包的依赖。

## 全局执行`global`
`global`命令允许你在`COMPOSER HOME`目录下执行其它命令。  
如果你将`COMPOSER HOME/vendor/bin`加入到`$PATH`环境变量中,你就可以用它在命令行中安装全局应用:
```bash
php composer.phar global require fabpot/php-cs-fixer:dev-master
```
这样上面的php-cs-fixer就可以在全局范围使用了

## 搜索`search`
为当前项目搜索依赖包,通常它只搜索[packagist.org](https://packagist.org)上的包。
```bash
php composer.phar search monolog
```

### 搜索参数
* **--only-name(-N)**:仅针对指定名称搜索(完全匹配)。

## 展示`show`
列出所有可用的软件包
```bash
## 列出所有软件包
php composer.phar show 
## 列出单个包的详细信息
php composer.phar show monolog/monolog
```
### 展示参数 
* **--installed(-i)**:列出已安装的依赖包。
* **--platform(-p)**:仅列出平台软件包(PHP与它的扩展)
* **--self(-s)**:仅列出当前项目信息

## 依赖性检测`depends`
`depends`命令可以查出已安装在你项目中的某个包，是否正在被其它的包所依赖，并列出他们。
```bash
php composer.phar depends --link-type=require monolog/monolog

nrk/monolog-fluent
poc/poc
propel/propel
symfony/monolog-bridge
symfony/symfony
```
### 依赖性检测-参数
* **--link-type**:检测的类型,默认为`require`也可以是`require-dev`
  
## 有效性检测`validate`
在提交`composer.json`文件，和创建`tag`前，你应该始终运行`validate`命令。它将检测你的`composer.json`文件是否是有效的
```bash
php composer.phar validate
```
### 有效性检测-参数
* **--no-check-all**:Composer是否进行完整的检验

## 依赖包状态检测`status`
如果你经常修改依赖包里的代码，并且它们是从`source`（自定义源）进行安装的，那么 `status`命令允许你进行检查，如果你有任何本地的更改它将会给予提示。
```bash
php composer.phar status -v
```

## 自我更新`self-update`
```bash
php composer.phar self-update
```
### 自我更新-参数
* **--rollback**(`-r`):回滚到你已经安装的最后一个版本。
* **--clean-backups**:在更新过程中删除旧的备份，这使得更新过后的当前版本是唯一可用的备份。

## 更改配置`config`
`config`命令允许你编辑`Composer`的一些基本设置，无论是本地的`composer.json`或者全局的`config.json`文件
```bash
php composer.phar config --list
```
### 更改配置-使用方法
```bash
config [options] [setting-key] [setting-value1] ... [setting-valueN]
```
`setting-key`是一个配置选项的名称，`setting-value1`是一个配置的值。可以使用数组作为配置的值（像`github-protocols`），多个`setting-value`是允许的。
### 更改配置-参数
* **--global (-g)**: 操作位于`$COMPOSER_HOME/config.json`的全局配置文件。如果不指定该参数，此命令将影响当前项目的`composer.json`文件，或`--file`参数所指向的文件。
* **--editor (-e)**: 使用文本编辑器打开`composer.json`文件。默认情况下始终是打开当前项目的文件。当存在`--global`参数时，将会打开全局`composer.json`文件。
* **--unset**: 移除由`setting-key`指定名称的配置选项。
* **--list (-l)**: 显示当前配置选项的列表。当存在`--global`参数时，将会显示全局配置选项的列表。
* **--file="..." (-f)**: 在一个指定的文件上操作，而不是`composer.json`。注意：不能与`--global`参数一起使用。

### 修改来源包
除了修改配置选项，`config`命令还支持通过以下方法修改来源信息：
```bash
php composer.phar config repositories.foo vcs http://github.com/foo/bar
```

## 创建项目`create-project`
相当于执行了一个`git clone`或`svn checkout`命令后将这个包的依赖安装到它自己的`vendor`目录。  
此命令有几个常见的用途：
* 你可以快速的部署你的应用。
* 你可以检出任何资源包，并开发它的补丁。
* 多人开发项目，可以用它来加快应用的初始化。

要创建基于`Composer`的新项目，你可以使用`"create-project"`命令。传递一个包名，它会为你创建项目的目录。你也可以在第三个参数中指定版本号，否则将获取最新的版本。
如果该目录目前不存在，则会在安装过程中自动创建。
```bash
php composer.phar create-project doctrine/orm path 2.2.*
```
### 创建项目-参数 
* **--repository-url**:提供一个自定义的储存库来搜索包，这将被用来代替`packagist.org`.可以是一个指向`composer`资源库的`HTTP URL`，或者是指向某个`packages.json`文件的本地路径。
* **--stability (-s)**:资源包的最低稳定版本，默认为`stable`。

## 打印自动加载索引`dump-autoload`
某些情况下你需要更新`autoloader`，例如在你的包中加入了一个新的类。你可以使用 `dump-autoload`来完成，而不必执行`install`或`update`命令。
### 打印自动加载索引-参数
* **--optimize (-o)**:转换`PSR-0/4 autoloading`到`classmap`获得更快的载入速度。这特别适用于生产环境，但可能需要一些时间来运行，因此它目前不是默认设置。
* **--no-dev**:禁用 autoload-dev 规则。

