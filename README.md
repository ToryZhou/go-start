#### Welcome to go start

# Idea config guide
### 1.在idea安装 go 语言插件<br>
settings->plugins->browse repositories->go->install
### 2.下载go语言环境
[下载 golang](https://golang.org/doc/install?download=go1.10.1.linux-amd64.tar.gz)

### 3.copy到program目录并解压
```sbtshell
tar xvf go1.10.1.linux-amd64
export PATH=$PATH:/home/toryzhou/Program/go1.10.1.linux-amd64/go/bin
```

### 4.设置go sdk 到go目录
```
/home/toryzhou/Program/go1.10.1.linux-amd64/go
Setup GOROOT //go的安装目录
```

### 5.设置运行目录到项目地址
Configure GOPATH //项目path设置项目所在目录<br>
The Go path is used to resolve import statements.<br>
go get {import}, To download the import into the GO path

```
go env GOPATH
go get github.com/hyperledger/fabric/core/chaincode/shim
```

素材来源于
- [菜鸟教程](http://www.runoob.com/go)
- [go-training](https://github.com/go-training/training)