#!/bin/bash

# 把整个项目需要的东西打包，直接可以scp到服务器

# 切换目录
ABSOLUTE_PROJECT_PATH="$(cd "$(dirname "$0")/.."||exit; pwd)"
cd "$ABSOLUTE_PROJECT_PATH" || exit

BUILD_PATH="build"
# 删除已有文件夹
if [ -d "${BUILD_PATH}" ]; then
  rm -r "${BUILD_PATH}"
fi

PROEJCT_NAME="${ABSOLUTE_PROJECT_PATH##*/}"
BUILD_PROJECT_DIR="${BUILD_PATH}/${PROEJCT_NAME}"
mkdir -p "${BUILD_PROJECT_DIR}"


echo '编译前端'
sh "script/build-frontend.sh" || exit

echo '编译 go 程序'
sh "script/build-go.sh" || exit

echo 'copy 二进制文件'
cp -r "bin" "$BUILD_PROJECT_DIR" || exit

echo 'copy 资源文件'
cp -r "resource" "$BUILD_PROJECT_DIR" || exit

echo 'copy Python 代码'
cp -r "py" "$BUILD_PROJECT_DIR" || exit
# 删除无用文件
rm -r "$BUILD_PROJECT_DIR/py/.idea"

echo 'copy Docker file'
cp -r "docker" "$BUILD_PROJECT_DIR" || exit

echo 'copy 脚本文件'
BUILD_SCRIPT_DIR="${BUILD_PROJECT_DIR}/script"
mkdir $BUILD_SCRIPT_DIR
cp "script/init-project.sh" "$BUILD_SCRIPT_DIR" || exit
cp "script/run-release.sh" "$BUILD_SCRIPT_DIR" || exit
cp "script/stop-release.sh" "$BUILD_SCRIPT_DIR" || exit

echo 'copy 配置文件'
BUILD_CONF_DIR="${BUILD_PROJECT_DIR}/conf"
mkdir $BUILD_CONF_DIR
cp -r "conf/app.release.yml" "$BUILD_CONF_DIR" || exit
cp -r "conf/nginx.release.conf" "$BUILD_CONF_DIR" || exit

echo 'copy docker-compose 文件'
cp "docker-compose.release.yml" "$BUILD_PROJECT_DIR" || exit

echo '打包项目'
cd "$BUILD_PATH" || exit
tar -zcf "${PROEJCT_NAME}.tar.gz" "${PROEJCT_NAME}" || exit

echo '项目打包完成！'