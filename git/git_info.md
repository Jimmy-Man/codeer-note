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

修改项目根目录下的.git/config `remote "origin"` 从HTTPS方式改为git请求方式
```bash
# .git/config
[remote "origin"]
	url = https://github.com/Jimmy-Man/codeer-note.git

```
改为
```bash
[remote "origin"]
	url = gut@github.com:Jimmy-Man/codeer-note.git
```

-------------------------------------------------