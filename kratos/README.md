快速开始
https://github.com/go-kratos/kratos/blob/main/README_zh.md

安装依赖
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

安装cli
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
kratos upgrade

创建项目模板
kratos new helloworld
cd helloworld
go mod download

生成所有 proto 源码、wire 等等
go generate ./...

运行程序
kratos run

目录结构
    api：定义了proto文件
    cmd：程序入口
    internel：内部文件
        server：提供grpc和http服务
        service：实现api中定义的接口

项目流程
    定义api：kratos proto add api/bubble/v1/todo.proto，修改proto中内容
    生成代码：make api/kratos proto client api/bubble/v1/todo.proto
            kratos proto server api/bubble/v1/todo.proto -t internal/service
    修改代码：在service中定义自己的usercase，service代码负责接收请求，调用usercase的业务逻辑，返回
            usercase为实际业务逻辑，调用repo进行数据操作
    依赖注入：在cmd目录下运行wire
            依赖注入降低代码的耦合度，外部以参数的形式往构造函数传入依赖项

