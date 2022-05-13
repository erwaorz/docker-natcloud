# docker-natcloud
##### docker共享容器云
This a  simple Docker REST API that providing a friendly API to manage Docker container.  
这是一个简单的Docker容器的REST API服务,提供友好的API供操作容器。
# 本地编译
1.下载安装最新 Go 环境  
2.clone 并进入本项目，下载所需包
```bash
git clone --depth=1 https://github.com/erwaorz/docker-natcloudd.git  
cd docker-natcloud  
go version  
go env -w GOPROXY=https://goproxy.cn,direct  
go env -w GO111MODULE=auto  
go mod tidy  
```
3.编辑 main.go 文件，内容按需修改  
4.按照平台输入命令编译，下面举了一些例子  
```bash
# 本机平台
go build -ldflags "-s -w"  -trimpath
```
### 接口
- [x] 开通容器  
- [x] 删除容器  
- [x] 启动容器  
- [x] 重启容器  
- [x] 停止容器  
- [x] 冻结容器  
- [x] 解冻容器  
- [x] 容器信息  
- [x] NAT端口绑定  
- [x] NAT端口解绑  
- [ ] 域名绑定
- [ ] 域名解绑
