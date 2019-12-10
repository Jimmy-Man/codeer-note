git
====


-------------------------------------------------
生成sshkey:
```
ssh-keygen -t rsa -C "xxxx@xxx.com"
```
* 按提示按三次回车
* 查看public key
```
# ~/.ssh/id_rsa.pub (默认路径)
cat ~/.ssh/id_rsa.pub
```
* 在git管理平台添加部署公钥
* 主机端增加信任列表
```
  ## 以码云为例
  ssh -T git@gitee.com
```
-------------------------------------------------