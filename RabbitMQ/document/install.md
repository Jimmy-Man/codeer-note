RabbitMQ安装
===========

### 正常安装
* 官网下载
* 安装Linux必要依赖包
* 下载rabbitMQ必须安装包
* 配置文件修改

-----------
### Docker安装

```
docker run -d --hostname my-rabbit --name rabbit -p 8080:15672 rabbitmq:management
--hostname：指定容器主机名称
--name:指定容器名称
-p:将mq端口号映射到本地

或在运行时设置用户和密码
docker run -d --hostname my-rabbit --name rabbit -e  RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=admin -p 15672:15672 -p 5672:5672 -p 25672:25672 -p 61613:61613 -p 1883:1883 rabbitmq:management
 15672：控制台端口号
 5672：应用访问端口号
```



* 服务启动: `rabbitmq-server start &`
* 服务停止: `rabbitmqctl stop_app`
* 管理插件: `rabbitmq-plugins enable rabbitmq_management`
* 访问地址: `http://127.0.0.1:15672/`