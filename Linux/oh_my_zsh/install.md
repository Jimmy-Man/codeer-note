安装Oh my ZSH
===============

Centos 安装 
* 安装zsh
```bash
  yum update
  yum -y install zsh
```
* 检查是否安装成功
```bash
zsh --version
```
~~~bash
zsh 5.0.2 (x86_64-redhat-linux-gnu)
### 显示这个信息表示安装成功
~~~
* 更换系统shell为zsh
```bash
chsh -s $(which zsh)
chsh -l # 查看
chsh -s /bin/zsh #从上条命令查看到的目录进行设置
```
* 退出客户端重新登录
* 安装Oh my ZSH
  - [Oh My ZSH官网](https://ohmyz.sh)
  - 安装脚本
  - ```sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"```
  