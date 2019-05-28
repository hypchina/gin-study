# gin-study
此仓库为学习gin之用

## 项目目录结构

    ├── README.md
    ├── app
    │   ├── core
    │   │   ├── extend 
    │   │   ├── helper 快捷方法
    │   │   └── utils  工具类如db,mysql
    │   ├── http       具体应用层
    │   │   ├── controllers 控制器
    │   │   ├── filters     过滤器
    │   │   └── middleware  中间件
    │   └── logic       逻辑层
    │       ├── bean        接口响应类
    │       ├── dao         数据提供/操作
    │       ├── entity      数据实体
    │       ├── enum        常量,枚举,标签
    │       ├── models      数据模型
    │       └── service     数据服务
    ├── conf
    │   └── config.go   配置文件
    ├── go.mod          官方包管理
    ├── go.sum
    ├── main.go         项目入口
    ├── resource        其他资源文件
    │   └── lang.go         语言包
    ├── routes          路由
    │   └── route.go
    └── vendor          包管理目录
    
 ## 项目结构描述
