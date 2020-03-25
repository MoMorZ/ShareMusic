# ShareMusic

一个前后端分离的网易云同步听歌项目。

同时集成了网易云API的简化Go版本，需要的可以直接使用。

重构自大佬[KAMIENDER](https://github.com/KAMIENDER)。



## Front

安装

`npm install`

热运行

`npm run serve`



## Back

使用`Beego`编写，网易云的API改编自[NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi)。



热运行

`beego run sharemusic`

编译

`go build main.go`