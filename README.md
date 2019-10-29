# 动力电池数据分析系统

动力电池数据分析系统，RESTful 设计风格，前后端分离。

## 项目结构

```
.
├── conf                     # 放置配置文件
├── docker                   # dockerfile
├── frontend                 # 前端
├── go                       # Go 程序
│   ├── app                  # 按功能划分模块
│   │   └── main             # 主要功能实现
│   ├── cmd                  # 编译入口
│   └── pkg                  # 一些公用包
│       ├── checker          # 检查用户传来的参数
│       ├── conf             # 配置相关
│       ├── conv             # 类型转换
│       ├── db               # 数据库初始化
│       ├── jd               # json 响应相关封装
│       ├── jtime            # json 序列化时间时能返回指定格式
│       ├── mysqlx           # 增强原生 Mysql 查询功能
│       ├── producer         # 调用 Celery
│       └── security         # 安全相关
├── py                       # Python 程序
│   └── celery               # 异步执行 Python 代码
└── script                   # 脚本文件
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

4、注销并重新登录当前用户

### 启动 docker

Docker 是服务器-客户端架构。命令行运行 `docker` 命令的时候，需要本机有 Docker 服务。使用如下命令运行：

```bash
$ sudo systemctl start docker
```

## 项目初始化

### MySQL 初始化

创建相应数据库。

### 生成相应文件

```bash
$ ./init-project.sh
```

会生成 *.env* 文件供 docker-compose 使用。

### 创建配置文件

在 *conf* 文件夹中根据配置模板创建配置文件，文件名中 example，替换为 release 和 debug。

## 启动项目

### 开发环境

1、启动数据库和 Nginx：

```bash
$ ./script/run-debug.sh
```

2、启动 py-app-celery：

```
在 Pycharm 中配置 run template，
选择要执行的 py 模块 celery，
输入运行参数 `-A task worker --concurrency=2`，
最后指定好环境变量 CONF_FILE。
```

3、启动 go-app-main：

```
在 goland 中配置 run template，只需指定好环境变量 CONF_FILE。
```

4、启动前端：

```bash
npm run serve
```

5、访问 `127.0.0.1:8081`

### 生产环境

1、编译前端：

```bash
npm run build
```

2、把 *frontend* 下编译出来的 *dist* 文件夹放入 *resource* 文件夹中。

3、执行：

```bash
$ ./script/run-release.sh
```

4、访问 `127.0.0.1:3389`

## 数据库管理

### 管理 MySQL

浏览器访问 `<ip>:8080` 端口

### 管理 Mongo

使用 robo3t 软件。

### 管理 Redis

浏览器访问 `<ip>:8079` 端口

## 设计规范

### 数据库

- 表名一律使用单数

- Redis key 以业务名(或数据库名)为前缀(防止key冲突)，用冒号分隔，比如业务名:表名:id

### 前端

- URI 格式规范：1）URI中尽量使用连字符”-“代替下划线”_”的使用。2）URI中统一使用小写字母。3）用复数名词。一个资源URI可以这样定义：`https://www.example.com/api/v1/posts/{postId}/comments/{commentId}`

- 前端对后端返回的 JSON 字段的顺序一律假设是无序的

- 提交和返回的 JSON 使用小驼峰法命名

### 后端

- 数据库表不由某个 ORM 创建，而是手工创建

- 字段前后端都要校验

- 某些不确定情况，直接返回 500。因为 gin panic recover 后会返回 500

- （TODO）后端字段合法性校验在 service 做（URL 的 Param 的判空在 controller 也要做，因为是 URL 的逻辑），包括 URL 的 Param 和 Query，提交的数据（如 JSON）

- 后端字段合法性校验不依赖于 gin 的 ShouldBindxxx，出于逻辑和方便测试 service 上考虑

- service 中的 xxxService 结构体用于接收用户发送的数据；而 model 中结构体一般是返回给用户的数据格式

### Git

- commit 时附上版本号，log 中某版本号的最后一个 commit，必须保证可运行

### 杂

- 配置文件名中带有 release 的是生产环境的配置文件，带有 debug 的是开发环境配置文件

- 开发环境需要手动设置环境变量 `CONF_FILE`，指定配置文件路径

- gin 的请求 log 会在请求处理函数结束后打印（记录请求用时），所以请求 websocket 时，打印会很延迟

## TODO

- 加缓存

- 完善错误处理，将如 `error.New("xxx")` 提出来作为私有全局变量

- 测试
