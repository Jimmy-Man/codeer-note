docker 常见问题

### 非root权限运行docker
 1. 添加docker组:[如果没有]
  ```
    sudo groupadd docker
  ```
 2. 将需要的用户添加到group内[以下代码以用户为user1为例]
  ```
    sudo gpasswd -a user1 docker
  ```
 3. 重启docker服务
  ```
    sudo service docker restart
    //或
    sudo /etc/init.d/docker restart
  ```
  4. 刷新docker组成员
 ```
    newgrp - docker
 ```