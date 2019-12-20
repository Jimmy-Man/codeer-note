deployer
=========

### 安装 
[官方](https://deployer.org/)
```bash
curl -LO https://deployer.org/deployer.phar
mv deployer.phar /usr/local/bin/dep
chmod +x /usr/local/bin/dep
```

composer 安装
```sh
composer require deployer/deployer --dev
```

### 使用

1. 在程序目录下运行
```
dep init
```
`
  [0 ] Common
  [1 ] Laravel
  [2 ] Symfony
  [3 ] Yii
  [4 ] Yii2 Basic App
  [5 ] Yii2 Advanced App
  [6 ] Zend Framework
  [7 ] CakePHP
  [8 ] CodeIgniter
  [9 ] Drupal
  [10] TYPO3
`

Deployer 部署laravel

1. 进入程序根目录
```sh
# dep init
php vendor/deployer/deployer/bin/dep init
```
从出现的选择中选择Laravel

2. 修改程序有主目录下的deploy.php配置文件
```

```
3. 

### 服务器端配置

软件要求：
```
LNMP|LAMP
git
composer
npm
```

1. php配置被禁用的函数列表
  
  php.ini的配置项disable_functions要检查并从该项中去除`proc_open`，`proc_get_status`，`symlink`


2. 添加部署用户[不建议使用root用户]
   ```sh
   # 以dep为例
   sudo adduser dep
   ```
   修改部署用户的用户组[使部署用户有权限对目录进行修改]
   ```sh
   # 假设web的用户组为www
   sudo usermod -aG www dep
   ```
   我们通常需要将 deployer 用户权限分别设置为创建文件 644 与目录 755，这样一来，deployer 用户可以读写，但是组与其它用户只能读
   ~~~bash
   su dep # 切换到部署用户
   echo "umask 022" >> ~/.bashrc
   ~~~
   将部署用户添加到sudoers中
   ```bash
   vim /etc/sudoers
   #在文件最后添加
   dep ALL=(ALL) NOPASSWD: ALL
   ```
   接下来要对我们的web根目录授权，假设我们的`web`服务的根目录在 `/home/wwwroot/` 下，那么需要将这个目录的用户设置为`dep` ，组设置为`www`的用户`www`
  ```bash
  sudo chown -R dep:www /home/wwwroot
  ```
  为了让`dep`用户在 `/home/wwwroot`下创建的文件与目录集成根目录的权限设定（用户`dep`, 组：`www-data`），我们还需要一步操作
  ```bash
  sudo chmod g+s /home/wwwroot
  ```
  云服务器添加密钥对登录[云服务器添加密钥对登录](../../Linux/document/ssh_key.md)

3. 上传静态文件
  ```bash
  /opt/lampp/bin/php vendor/deployer/deployer/bin/dep vue:upload production -vvv
  ```

  UPDATE `business_types` SET `desction` = replace(`desction`,'192.168.2.30/kwpublic.com','public.jinxinheng.cn');

  UPDATE `businesses` SET `desction` = replace(`desction`,'192.168.2.30/kwpublic.com','public.jinxinheng.cn');