# Go程序编译之后会得到一个可执行的二进制文件，其实在最终的镜像中是不需要go编译器的，也就是说我们只需要一个运行最终二进制文件的容器即可。
# 作为别名为"builder"的编译镜像，下面会用到
FROM golang AS builder

# 为镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 设置工作目录：/src
WORKDIR /src

# 复制项目并下载依赖信息
ADD ./ /src
RUN go mod download

# 将代码编译成二进制可执行文件app
RUN go build -o /eas-admin eas-admin/main.go
RUN go build -o /eas-source-zkill eas-source-zkill/main.go

###################
# 接下来创建一个小镜像
###################
FROM scratch AS eas-admin

# 从builder镜像中把程序拷贝到当前目录
COPY --from=builder eas-admin /

# 声明服务端口
EXPOSE 8080

# 启动容器时运行的命令
ENTRYPOINT ["/eas-admin"]

FROM scratch AS eas-source-zkill

# 从builder镜像中把程序拷贝到当前目录
COPY --from=builder eas-source-zkill /

# 声明服务端口
EXPOSE 8080

# 启动容器时运行的命令
ENTRYPOINT ["/eas-source-zkill"]