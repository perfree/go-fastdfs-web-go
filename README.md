## Go-FastDfs-Web的GoLang版本

> go-fastdfs是一个基于http协议的分布式文件系统，它基于大道至简的设计理念，一切从简设计，使得它的运维及扩展变得更加简单，它具有高性能、高可靠、无中心、免维护等优点。Go-FastDfs-Web为go-fastdfs的web平台,旨在方便用户管理查看go-fastdfs文件

## 开发环境

1. GoVersion ： 1.13.4
2. Beego ： 1.12.0
3. sqlite ： 3

## 运行

```bash
# 下载依赖
go mod download
# 将依赖复制到vendor下
go mod vendor
# 运行
bee run 或者直接运行main.go
默认端口8080
```
如依赖下载失败则设置下代理:
```bash
go env -w GOPROXY=https://goproxy.io
```