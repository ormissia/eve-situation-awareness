#!/bin/sh

# 代码中服务目录在项目跟目录中的文件夹名称
SERVICE_DIR_NAME="main"

# 根据传入参数进行编译，参数1：服务名称，即每个服务文件夹名称
function Build {
  echo building "$1" ...
# 编译
  if [ "$2" == "binary" ]; then
    go build -o "$PROJECT_ROOT_DIR"/target/"$1" "$SERVICE_DIR_NAME"/"$1"/main.go
  elif [ "$2" == "docker-build" ]; then
    docker build -t ormissia/"$1" --build-arg SERVICE_NAME="$1" --progress=plain .
  elif [ "$2" == "docker-push" ]; then
    docker push ormissia/"$1"
  elif [ "$2" == "docker-clean" ]; then
    echo docker-clean
  else
    echo "incorrect compile type $1"
    exit
  fi
# 判断编译是否成功
  if [ $? -eq 0 ];then
    echo build "$1" successful
    BUILD_SUCCESSFUL_COUNT=$((BUILD_SUCCESSFUL_COUNT+1))
  else
    echo build "$1" faild
    BUILD_FAILED_COUNT=$((BUILD_FAILED_COUNT+1))
  fi
  echo "--------------------------"
}

# 获取项目的根目录
PROJECT_ROOT_DIR=$(pwd)
echo project root directory is "$PROJECT_ROOT_DIR"

# 编译成功数量
BUILD_SUCCESSFUL_COUNT=0
# 编译失败数量
BUILD_FAILED_COUNT=0

SERVICES=$(ls $SERVICE_DIR_NAME)

# sh build.sh binary/docker $(service)
# 当传入参数为2个时,第二个为服务名
if [ $# -eq 2 ]; then
  SERVICES=$2
fi

# 遍历所有文件夹名称，即所有服务
for directory in $SERVICES
do
# 调用编译函数
  Build "$directory" "$1"
done

# 输出编译结果
echo ">>> all service build finished <<<"
echo successful: $BUILD_SUCCESSFUL_COUNT "|" faild: $BUILD_FAILED_COUNT
