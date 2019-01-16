### Xvfb Linux 下模拟图形界面 
Xvfb 是什麼呢，他的名稱是 virtual framebuffer X server for X Version 11， Xvfb 可以直接處理 Window 的圖形化功能，並且不會把圖像輸出到螢幕上，也就是說，就算你的電腦沒有啟動 Xwindow ， 你仍然可以執行任何圖形程式。
与其他显示服务器相比，Xvfb在虚拟内存中执行所有图形操作，而不显示任何屏幕输出。从客户端的角度来看，它的行为与任何其他X显示服务器完全相同，可以根据需要提供请求并发送事件和错误。但是，没有显示输出。此虚拟服务器不需要运行它的计算机具有任何类型的图形适配器，屏幕或任何输入设备。只有一个网络 层是必要的。
[Xvfb官方文档](https://www.x.org/releases/X11R7.6/doc/man/man1/Xvfb.1.xhtml)
https://www.x.org/releases/X11R7.6/doc/man/man1/Xvfb.1.xhtml

#### 安装
```
sudo apg-get install xvfb
```
#### 启动xvfb服务
```
Xvfb -ac :99 -screen 0 1280x1024x8
```
#### 透xvfb执行命令
```
export  DISPLAY=:99

/usr/bin/google-chrome-stable http://www.investopedia.com         //chrome 浏览www.investopedia.com
```
#### 更详细会在后面进行补充
