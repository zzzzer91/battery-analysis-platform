#!/bin/bash

ABSOLUTE_PROJECT_PATH="$(cd "$(dirname $0)/..";pwd)"

RESOURECE_DIR="${ABSOLUTE_PROJECT_PATH}/resource"
MEDIA_DIR="${RESOURECE_DIR}/media"
FILE_DIR="${RESOURECE_DIR}/file"
FILE_DL_DIR="${FILE_DIR}/dl/model"

ENV_FILE="${ABSOLUTE_PROJECT_PATH}/.env"

echo '项目初始化中...'

mysql_root_password=''
mongo_root_password=''
read -p "输入 MySQL root 密码：" mysql_root_password
read -p "输入 Mongo root 密码：" mongo_root_password

# 如果文件夹不存在，创建文件夹
if [ ! -d "${RESOURECE_DIR}" ]; then
    mkdir "${RESOURECE_DIR}"
fi
if [ ! -d "${MEDIA_DIR}" ]; then
    mkdir "${MEDIA_DIR}"
fi
if [ ! -d "${FILE_DIR}" ]; then
    mkdir "${FILE_DIR}"
fi
if [ ! -d "${FILE_DL_DIR}" ]; then
    mkdir -p "${FILE_DL_DIR}"
fi

# 生成 .env，会覆盖原来的
echo "# docker-compose.yml 中使用的环境变量" > "${ENV_FILE}"
echo "# 注意值两边的单双引号，会被当作值的一部分，这在 docker-compose 中会出现问题" >> "${ENV_FILE}"
#
echo "PROJECT_DIR=${ABSOLUTE_PROJECT_PATH}" >> "${ENV_FILE}"
#
echo "# docker 映射数据库数据" >> "${ENV_FILE}"
echo "DATABASE_DATA_DIR=${ABSOLUTE_PROJECT_PATH}/database" >> "${ENV_FILE}"
#
echo "# MySQL" >> "${ENV_FILE}"
echo "MYSQL_ROOT_PASSWORD=${mysql_root_password}" >> "${ENV_FILE}"
#
echo "# Mongo" >> "${ENV_FILE}"
echo "MONGO_INITDB_ROOT_USERNAME=root" >> "${ENV_FILE}"
echo "MONGO_INITDB_ROOT_PASSWORD=${mongo_root_password}" >> "${ENV_FILE}"
#
echo "# Redis" >> "${ENV_FILE}"
#
echo "${ENV_FILE} 生成完毕！"

echo "项目初始化完毕！"