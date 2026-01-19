准备
GOROOT：go的安装路径，包含了标准库、编译器等
GOPATH：开发时的工作目录，保存第三方库

1.17后引入go mod进行包管理
go.mod中保存了引入的包名，代码中导入时去$GOPATH/pkg/mod中查找;
go install会将可执行文件下载到$GOPATH/bin
go.mod的路径代表了项目的根目录，导入子目录的包后，编译器会顺着根目录找;
大型项目初始化为github.com/path，用于下载


