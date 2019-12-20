`lsof -i :80` 查看指定端口执行的进程

###  nohup 
nohup [common] > com.log 2>&1 &

nohup zip -r fabzat_server_dev_shared.zip fabzat_server_dev_shared > zip.file 2>&1 &

nohup scp -i ~/data/UUPZ root@119.29.240.131:/sharedfolder/fabzat_server_dev_shared.zip ./ > scp.log 2>&1 &

[云服务root之外的帐号使用密钥对登录](./document/ssh_key.md)