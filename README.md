# my-blog

my-blog

# centos安装MySQL8

```shell
yum install mysql
yum install mysql-devel
```

# 交叉编译：mac上编译成Linux平台可执行程序

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# 以后台形式运行程序

- 在main.go中导入包："github.com/codyguo/godaemon"
- 启动命令添加-d参数设置为true
```shell
./example -d=true 
# 或者 
./example -d true
```
