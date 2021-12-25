# 先用Node镜像编译项目，再将编译好的项目复制进Nginx镜像中打包
# 作为别名为"builder"的编译镜像，下面会用到
FROM node AS builder

# 设置工作目录：/build
WORKDIR /app

# 复制文件
COPY . .
# 安装依赖
RUN npm install
# 构建应用
RUN npm run build

###################
# 接下来创建一个Nginx镜像将打包好的文件放进去
###################
FROM nginx

# 从builder镜像中把/build/dist 拷贝到/app
COPY --from=builder /app/build /app/build

# 删除原有配置文件
RUN rm -f  /etc/nginx/nginx.conf
# 将Nginx配置文件复制进去
COPY nginx.conf /etc/nginx

# 切换工作目录
WORKDIR /web/share
# 开启映射
VOLUME /web/share

# 声明服务端口
EXPOSE 8081
