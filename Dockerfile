FROM golang:latest

# 设置工作目录
WORKDIR /app

# 将项目文件复制到镜像中
COPY . .

# 添加环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# 构建应用程序
RUN go build -o main .

# 设置容器启动时执行的命令
CMD ["./main"]
