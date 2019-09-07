# 电池数据分析平台后端

电动汽车电池数据分析平台后端，RESTful 设计风格，前后端分离。

## 项目结构

```
TODO
```

## 前置需求

### Docker

CentOS7 下安装：

```bash
# 1、安装工具
$ sudo yum install -y yum-utils

# 2、添加仓库
$ sudo yum-config-manager \
    --add-repo \
    http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

# 3、安装
$ sudo yum install docker-ce docker-ce-cli containerd.io
```

### docker-compose

安装：

```bash
$ pip install -U docker-compose
```

### 安装后配置

Docker 需要用户具有 sudo 权限，为了避免每次命令都输入 `sudo`，可以把用户加入 Docker 用户组（[官方文档](https://docs.docker.com/install/linux/linux-postinstall/#manage-docker-as-a-non-root-user)），步骤如下：

1、创建 docker 组：

```bash
$ sudo groupadd docker
```

2、把当前用户加入 docker 组：

```bash
$ sudo usermod -aG docker $USER
```

3、检查是否加入成功：

```bash
$ groups $USER
```

4、注销并重新登录当前用户。

### 启动 docker

Docker 是服务器-客户端架构。命令行运行 `docker` 命令的时候，需要本机有 Docker 服务。使用如下命令运行：

```bash
$ sudo systemctl start docker
```

## 项目初始化

### 生成文件

```bash
$ ./init-project.sh
```

会生成 *.env* 文件供 docker-compose 使用。

### 创建配置文件

在 *conf* 文件夹中根据配置模板创建配置文件，去掉 *.example*。

### MySQL 初始化

1、创建相应数据库

2、启动程序后会自动生成 MySQL 表：

## 启动项目

TODO

## 其他

### 管理 MySQL

浏览器访问 `<ip>:8080` 端口。

### 管理 Mongo

使用 robo3t 软件。

### 管理 Redis

使用 GUI 软件

## 设计原则

### 前端

- 前端对后端返回的 JSON 字段的顺序一律假设是无序的

### 后端

- 字段前后端都要校验

- 某些不确定情况，直接返回 500。因为 gin panic recover 后会返回 500

- （TODO）后端字段合法性校验在 service 做（URL 的 Param 的判空在 controller 也要做，因为是 URL 的逻辑），包括 URL 的 Param 和 Query，提交的数据（如 JSON）

- 后端字段合法性校验不依赖于 gin 的 ShouldBindxxx，出于逻辑和方便测试 service 上考虑

### git

- commit 时附上版本号，log 中某版本号的最后一个 commit，必须保证可运行

## TODO

- 完善错误处理，将如 `error.New("xxx")` 提出来作为私有全局变量

- 测试
