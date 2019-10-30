#!/bin/bash

ABSOLUTE_PROJECT_PATH="$(cd "$(dirname $0)/..";pwd)"
GO_DIR_PATH="${ABSOLUTE_PROJECT_PATH}/go"
CMD_PATH="${GO_DIR_PATH}/cmd"
BIN_PATH="${ABSOLUTE_PROJECT_PATH}/bin"

if [ ! -d "${BIN_PATH}" ]; then
  mkdir "${BIN_PATH}"
fi

cd $GO_DIR_PATH
for dir in `ls $CMD_PATH`; do
  # 如果文件夹不存在，创建文件夹
  path="${BIN_PATH}/${dir}"
  if [ ! -d "$path" ]; then
    mkdir "${path}"
  fi
  echo "build => $path/app"
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "${path}/app" "${CMD_PATH}/${dir}/main.go"
done