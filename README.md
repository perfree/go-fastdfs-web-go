## GoFastDfsWeb - Golang版本
> go-fastdfs 是一个简单的分布式文件存储，具有高性能，高可靠，免维护等优点，支持断点续传，分块上传，小文件合并，自动同步，自动修复。本项目为go-fastdfs的web管理端

[前往 Go-Fastdfs](https://github.com/sjqzhang/go-fastdfs)

[Go-Fastdfs web管理平台 - Java版本](https://github.com/perfree/go-fastdfs-web)

## Conf
conf/app.ini
```
[app]
RunMode = release
HttpPort = 8088
SqlFile = DataBase.db
LogFile = go-fastDfs-web-go.log
SessionSecret = 123456
AppVer = dev1.0.0
AppVerDate = 2021-09-29
```
## Run
```
go run main.go 
```
