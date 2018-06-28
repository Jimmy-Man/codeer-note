# VMWare虚拟机通过主机shadowsocks代理上网
## 环境
宿主：windows 7
VM: kali linux(基本上linux系统都可以)
## 步骤
在宿主机windows上运行shadowsocks.exe并勾选“允许局域网连接”
使用桥接方式运行虚拟机（这时虚拟机与宿主处于同一个局域网）
进入linux系统，System Settings – Network – Network proxy勾选Manual（手动）,地址全部填宿主机IP（局域网网段），设置好代理端口（可在windows下的shadowsocks查看，一般为默认1080）
linux用浏览器访问www.google.com，成功
