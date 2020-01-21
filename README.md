# 动力电池数据分析系统

## 亮点

- 100%容器化 ，一键打包项目

- RESTful 设计风格，前后端分离

- 基于 Echarts.js，支持 2D 和 3D 及动态可视化

- 基于 Celery，异步创建计算任务

- 基于 PyTorch，支持创建深度学习训练，并可实时画出训练过程

- 内置一个强大的 Markdown 编辑器，能够插入数学公式和图片，并支持导出为 PDF

- 多用户管理，以及访问控制

## 部分功能演示

### 查询历史数据及三维可视化

![查询历史数据及三维可视化](./doc/imgs/查询历史数据及三维可视化.gif)

### 创建神经网络训练任务

![创建神经网络训练任务](./doc/imgs/创建神经网络训练任务.gif)

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
$ sudo yum install -y yum-utils \
    device-mapper-persistent-data \
    lvm2

# 2、添加仓库
$ sudo yum-config-manager \
    --add-repo \
    http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
# 国外使用 https://download.docker.com/linux/centos/docker-ce.repo

# 3、安装
$ sudo yum install docker-ce docker-ce-cli containerd.io
```

### docker-compose

安装：

```bash
$ pip3 install -U docker-compose
```

### 安装后配置

Docker 需要用户具有 root 权限（当前 rootless 特性处于实验阶段），为了避免每次命令都输入 `sudo`，可以把用户加入 Docker 用户组（[官方文档](https://docs.docker.com/install/linux/linux-postinstall/#manage-docker-as-a-non-root-user)），步骤如下：

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

## 启动项目

### 创建配置文件

在 *conf* 文件夹中根据配置模板创建配置文件，文件名中 example，替换为 release 和 debug。

### 开发环境

1、初始化项目（会生成 *.env* 文件供 docker-compose 使用。docker 数据库文件映射默认在当前项目的 *database* 文件夹下）：

```bash
$ ./init-project.sh
```

2、启动数据库和 Nginx：

```bash
$ ./script/run-debug.sh
```

3、在 goland 中配置 run template，只需指定好环境变量 `CONF_FILE`，代表 debug 配置文件路径。然后运行。

4、在 Pycharm 中配置 run template，选择要执行的 py 模块 celery，输入运行参数 `-A task worker --concurrency=3`，最后指定好环境变量 `CONF_FILE`。然后运行。

5、安装前端依赖：

```bash
$ cd frontend
$ npm install
```

6、启动前端：

```bash
$ npm run serve
```

7、访问 `127.0.0.1:8081`

### 生产环境

1、`./script/pack-project.sh` 打包项目

2、`scp` 到服务器

3、执行 `./script/init-project.sh`，初始化相关配置

4、执行 `./script/run-release.sh` 启动项目

## 数据库管理

### 管理 Mongo

使用 robo3t 软件。

### 管理 Redis

浏览器访问 `<ip>:8079` 端口

## 设计规范及注意点

仅供自己参考。

### 数据库

- 表名一律使用单数
- MySQL，MongoDB 表名使用下划线分割
- Redis key 以业务名（或数据库名）为前缀（防止 key 冲突），用冒号分隔，比如业务名:表名:id
- MySQL，Redis，MongoDB key 名使用小驼峰法命名
- MongoDB 中 key 可以用中文，但最好是英文
- MySQL 中字段不允许为 null，有一些坑，比如查询 `id != 1` 时，null 值不会被匹配，因为 null 与其他值都不相等。要匹配到，要用 `id != 1 or id is null`
- 接上条，用 or 时，MySQL 可能不走索引，所以尽量用 union。

### 前端

- URI 格式规范：1）URI中尽量使用连字符”-“代替下划线”_”的使用。2）URI中统一使用小写字母。3）用复数名词。一个资源URI可以这样定义：`https://www.example.com/api/v1/posts/{postId}/comments/{commentId}?过滤条件`
- 上一条的 URI 中的过滤条件采用小驼峰法命名
- 前端对后端返回的 JSON 字段的顺序一律假设是无序的
- 提交和返回的 JSON 使用小驼峰法命名

### 后端

- 数据库表不由某个 ORM 创建，而是手工创建，统一管理
- 字典的 key 都假设是无序的（即使在 Python 中是有序的）
- 字段前后端都要校验
- 某些不确定情况，直接返回 500。因为 gin panic recover 后会返回 500
- 后端字段合法性校验不依赖于 gin 的 ShouldBindxxx，出于逻辑和方便测试 service 上考虑
- service 中的 xxxService 结构体用于接收用户发送的数据；而 model 中结构体一般是返回给用户的数据格式
- Go 中默认值尽量不要从 0，而是从 1 开始，因为 Go 中初始化值为 0

### 时区

- Python 中的时间类型 `datetime.datetime` 是**本地时区**的，时间字符串使用 `datetime.datetime.strptime` 解析，直接转换成本地时间
- Go 中的时间类型 `time.Time` 也是**本地时区**，Go 中把本地时区的时间字符串解析成 `time.Time` 类型时，要用 `time.ParseInLocation` 而不是 `time.Parse` 

- MySQL 中时间类型一律使用 `datatime`，它是**无时区**的，如果可以，最好保存 UTC 时间
- Python 的 PyMySQL 库插入或读取 `datetime.datetime` 类型时，不会做任何转换
- Go 的 go-sql-driver/mysql 库插入或读取 `time.Time` 类型时，也不会做任何转换
- MongoDB 时间类型用 `ISODate`，本质是 Unix 时间戳，在 Robo3T 软件中默认显示为 UTC 时区格式（可以调为本地时区）
- （**注意**）Python 的 PyMongo 库和 Go 的 go-mongo-driver 库在**默认配置**下对 MongoDB 的 `ISODate` 的处理方式有很大不同，不注意会造成时区的误差，见下
- Python 的 `datatime.datatime` 插入 MongoDB 时，会直接把插入的时间当成 UTC 时间，这时就会有 8 个小时的误差，所以插入当前时间要用 `datetime.datetime.utcnow()`；读取时会直接把 MongoDB 中的 UTC 时区直接读入 `datatime.datatime`，而不做任何转换
- Go 的 `time.Time` 读取或插入 MongoDB 时，会自动转换时区（插入时本地时区转为 UTC 时区，读取时 UTC 时区转为本地时区）

### Git

- commit 时附上版本号，log 中某版本号的最后一个 commit，必须保证可运行

### docker

- dockerfile 中工作目录统一在 */root* 目录下； docker 容器中所有资源文件也都映射在 */root* 目录下

### 杂

- 配置文件名中带有 release 的是生产环境的配置文件，带有 debug 的是开发环境配置文件
- 开发环境需要手动设置环境变量 `CONF_FILE`，指定配置文件路径
- gin 的请求 log 会在请求处理函数结束后打印（记录请求用时），所以请求 websocket 时，打印会很延迟

## TODO

- 完善错误处理，如 `error.New("xxx")` 提出来赋值给全局变量
- 测试
