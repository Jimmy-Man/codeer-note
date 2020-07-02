Raspberry Pi 
================

##### [Mac安装Raspberry  Pi操作系统](./document/install.md)
##### [Raspberry Pi更换国内源](./document/source.md)

#####  更改VNC分辨率
```bash
sudo raspi-config
```
1. Advanace Options --> Resolution --> 选择需要的分辨率
2. Finsh 重启树莓派或VNC

##### 配置`wifi`
* WIFI帐号密码信息保存在`/etc/wpa_supplicant/wpa_supplicant.conf`
```conf
# 保存wifi帐号密码文件
/etc/wpa_supplicant/wpa_supplicant.conf
## 新增网络添加如下: 一个network一个连接
network={
	ssid="wifi连接名称"
	psk="wifi连接密码"
	key_mgmt=WPA-PSK 
}
```
* key_mgmt 参数使用WPA-PSK为大部分WIFI接入方式


中科大源
http://mirrors.ustc.edu.cn/raspbian/raspbian/