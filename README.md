# RentSysServ
住房租赁平台服务器

## 安装

  ```
  go get -u github.com/bokwun/RentSysServ
  ```
  
## 介绍和初始化项目

### 介绍

  Go版本需要>=`Go.1.12`,Gin为`1.4`版本

### 初始项目数据库

  新建`user`数据库，编码为`utf8`
  
  在`user`数据库下创建以下表：
  
  1.个人账户信息表
 ```
CREATE TABLE `person` (
    `id` varchar(40) NOT NULL,
    `password` varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
  ```
  2.租房信息表
  ```
CREATE TABLE `message` (
    `idUser` varchar(40) NOT NULL,
    `xiaoqumc` varchar(50) NOT NULL,
    `shi` tinyint(4) NOT NULL,
    `ting` tinyint(4) NOT NULL,
    `wei` tinyint(4) NOT NULL,
    `mianji` int(11) NOT NULL,
    `diceng` tinyint(4) NOT NULL,
    `gongceng` tinyint(4) NOT NULL,
    `chewei` varchar(10) NOT NULL,
    `zujin` int(11) NOT NULL,
    `quyu` varchar(10) NOT NULL,
    `biaoti` varchar(50) NOT NULL,
    `miaoshu` varchar(1000) NOT NULL,
    `lianxiren` varchar(40) NOT NULL,
    `lianxidh` varchar(50) NOT NULL,
    `dateTime` varchar(50) NOT NULL,
    `picName` varchar(50) NOT NULL,
    KEY `id_user` (`idUser`),
    CONSTRAINT `message_ibfk_1` FOREIGN KEY (`idUser`) REFERENCES `person` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
  ```
  
## 已有API

```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/postmsg              --> main.main.func1 (4 handlers)
[GIN-debug] POST   /api/message              --> main.main.func2 (4 handlers)
[GIN-debug] GET    /api/checkUser/:id        --> main.main.func3 (4 handlers)
[GIN-debug] GET    /api/returnMsg            --> main.main.func4 (4 handlers)
[GIN-debug] GET    /api/returnUserMsg/:idUser --> main.main.func5 (4 handlers)
[GIN-debug] GET    /api/returnMsgWhere       --> main.main.func6 (4 handlers)
[GIN-debug] GET    /api/delMessage           --> main.main.func7 (4 handlers)
[GIN-debug] POST   /api/upload               --> main.main.func8 (4 handlers)
[GIN-debug] Listening and serving HTTP on :8081
```

1. /api/postmsg           处理账户信息表单的提交
2. /api/message           处理房源信息表单的提交
3. /api/checkUser/:id     处理按账号名查询的的请求
4. /api/returnMsg         返回信息
5. /api/returnUserMsg/:id 处理按账户名查询的请求
6. /api/returnMsgWhere    处理按条件查询的请求
7. /api/delMessage        删除数据
8. /api/upload            处理图片上传

