git
====

### git免密操作|免密部署(生成与配置密钥对)
在开发过程中和服务器部署的时候，如果没有配置公钥每次`push`和`pull`的操作都需要输入帐号和密码，这造成了开发效率和部署的问题,下面是创建和配置密钥对的方法:
1. 在开发的机子或者需要部署的机子上先生成密钥对(注意先切换到部署的用户不建议使用root直接部署)
```
ssh-keygen -t rsa -C "xxxx@xxx.com"
```
根据提示三次回车后完成生成密钥对操作

1. 查看已生成的密钥对[密钥对的保存路径可以在第1步中设置]
```bash
# ~/.ssh/id_rsa.pub (默认路径)
cat ~/.ssh/id_rsa.pub
```
`.pub`扩展名为公钥,主要使用它来配置到git服务商(或者自建git)

3. 在git管理平台[`github`|`码云`|`gitLab`等]添加部署公钥(或者ssh密钥)
   * 部署公钥 只能进行pull与clone等只读操作[通常用于服务器部署操作]
   * SSH公钥 可以进行任何读写操作 [通常开发者使用]
  
4. 在开发主机或者部署部署服务器中增加信任列表
```bash
  ## 以码云为例
  ssh -T git@gitee.com
```

5. 修改远程仓库的传输协议(把`HTTPS`改为`SSH`的形式)
   * 本地项目 - 修改项目根目录下的`.git/config`文件中的`remote "origin`选项
   * 部署项目 - 修改部署工具的`repository`源为SSH的方式  
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
6. 测试
  ```bash
  git pull 
  ```

  收工
-------------------------------------------------