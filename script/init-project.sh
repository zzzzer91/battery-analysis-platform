#!/bin/bash

ABSOLUTE_PROJECT_PATH="$(cd "$(dirname $0)/..";pwd)"
CONFIG_DIR="${ABSOLUTE_PROJECT_PATH}/conf"
INSTANCE_DIR="${ABSOLUTE_PROJECT_PATH}/instance"
DATABASE_DIR="${ABSOLUTE_PROJECT_PATH}/database"
MEDIA_DIR="${INSTANCE_DIR}/media"
DIST_DIR="${INSTANCE_DIR}/dist"
ENV_FILE="${ABSOLUTE_PROJECT_PATH}/.env"

echo '项目初始化中...'

# 如果文件夹不存在，创建文件夹
if [ ! -d "${INSTANCE_DIR}" ]; then
    mkdir "${INSTANCE_DIR}"
    echo "${INSTANCE_DIR} 创建完毕！"
else
    echo "${INSTANCE_DIR} 已存在！"
fi

if [ ! -d "${DATABASE_DIR}" ]; then
    mkdir "${DATABASE_DIR}"
    echo "${DATABASE_DIR} 创建完毕！"
else
    echo "${DATABASE_DIR} 已存在！"
fi

if [ ! -d "${MEDIA_DIR}" ]; then
    mkdir "${MEDIA_DIR}"
    echo "${MEDIA_DIR} 创建完毕！"
else
    echo "${MEDIA_DIR} 已存在！"
fi

if [ ! -d "${DIST_DIR}" ]; then
    mkdir "${DIST_DIR}"
    echo "${DIST_DIR} 创建完毕！"
else
    echo "${DIST_DIR} 已存在！"
fi

mysql_root_password=''
mysql_database=''
mongo_root_password=''
mongo_database=''
read -p "输入 MySQL root 密码：" mysql_root_password
read -p "输入 MySQL 数据库名：" mysql_database
read -p "输入 Mongo root 密码：" mongo_root_password
read -p "输入 Mongo 数据库名：" mongo_database

# 生成 .env
echo "# docker-compose.yml 中使用的环境变量" > ${ENV_FILE}
echo "# 注意值两边的单双引号，会被当作值的一部分，这在 docker-compose 中会出现问题" >> ${ENV_FILE}
#
echo "PROJECT_DIR=${ABSOLUTE_PROJECT_PATH}" >> ${ENV_FILE}
#
echo "# docker 映射数据库数据" >> ${ENV_FILE}
echo "DATABASE_DATA_DIR=${DATABASE_DIR}" >> ${ENV_FILE}
#
echo "# MySQL" >> ${ENV_FILE}
echo "MYSQL_ROOT_PASSWORD=${mysql_root_password}" >> ${ENV_FILE}
#
echo "# Mongo" >> ${ENV_FILE}
echo "MONGO_INITDB_ROOT_USERNAME=root" >> ${ENV_FILE}
echo "MONGO_INITDB_ROOT_PASSWORD=${mongo_root_password}" >> ${ENV_FILE}
#
echo "# Redis" >> ${ENV_FILE}
#
echo "${ENV_FILE} 生成完毕！"

echo "项目初始化完毕！"