#!/bin/bash

# 根据传入参数进行编译，参数1：服务名称，即每个服务文件夹名称
function GoBuild {
  echo go building "$1" ...
# 编译
  go build -o "$PROJECT_ROOT_DIR"/target/"$1" "$1"/main.go
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


# 进入服务目录
cd main || exit

SERVICES=$(ls)

# 当参数个数不为0时，需要编译的服务为传入的参数
if [ $# -ne 0 ]; then
    SERVICES=$*
fi

# 遍历所有文件夹名称，即所有服务
for directory in $SERVICES
do
# 调用编译函数
  GoBuild "$directory"
done

# 输出编译结果
echo ">>> all service build finished <<<"
echo successful: $BUILD_SUCCESSFUL_COUNT "|" faild: $BUILD_FAILED_COUNT
