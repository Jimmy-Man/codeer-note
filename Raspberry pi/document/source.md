RaspberryPi源
==============


#### 更换国内源
###### 1. 备份原文件
```bash
sudo cp  /etc/apt/source.list /etc/apt/source.list-bak
```
###### 2. 更换国内源
```bash 
sudo nano /etc/apt/source.list
```
把源地址更换为国内源地址
###### 3. 更新源
```bash
sudo apt-get update
```
###### 4. 更新软件
```bash
sudo apt-get upgrade
```

#### 国内源参考 

中国科学技术大学
Raspbian http://mirrors.ustc.edu.cn/raspbian/raspbian/

阿里云
Raspbian http://mirrors.aliyun.com/raspbian/raspbian/

清华大学
Raspbian http://mirrors.tuna.tsinghua.edu.cn/raspbian/raspbian/

华中科技大学
Raspbian http://mirrors.hustunique.com/raspbian/raspbian/ Arch Linux ARM http://mirrors.hustunique.com/archlinuxarm/

华南农业大学（华南用户）
Raspbian http://mirrors.scau.edu.cn/raspbian/

大连东软信息学院源（北方用户）
Raspbian http://mirrors.neusoft.edu.cn/raspbian/raspbian/

重庆大学源（中西部用户）
Raspbian http://mirrors.cqu.edu.cn/Raspbian/raspbian/

新加坡国立大学
Raspbian http://mirror.nus.edu.sg/raspbian/raspbian

牛津大学
Raspbian http://mirror.ox.ac.uk/sites/archive.raspbian.org/archive/raspbian/
