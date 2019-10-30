#!/bin/bash

ABSOLUTE_PROJECT_PATH="$(cd "$(dirname "$0")/.." || exit;pwd)"

GO_CODE_DIR="${ABSOLUTE_PROJECT_PATH}/go"
CMD_DIR="${GO_CODE_DIR}/cmd"
BIN_DIR="${ABSOLUTE_PROJECT_PATH}/bin"

if [ ! -d "${BIN_DIR}" ]; then
  mkdir "${BIN_DIR}"
fi

cd "$GO_CODE_DIR" || exit
for dir in `ls $CMD_DIR`; do
  # 如果文件夹不存在，创建文件夹
  path="${BIN_DIR}/${dir}"
  if [ ! -d "$path" ]; then
    mkdir "${path}"
  fi
  echo "build => $path/app"
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "${path}/app" "${CMD_DIR}/${dir}/main.go"
done

echo '编译完成！'