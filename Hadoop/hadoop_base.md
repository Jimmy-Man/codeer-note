Hadoop 基础
==========
生成密钥
ssh-keygen -t dsa -P '' -f ~/.ssh/id_dsa

将生成的密钥放置到许可证文件中
cat ~/.ssh/id_dsa.pub >> ~/.ssh/authorized_key


