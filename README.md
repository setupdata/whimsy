# Whimsy Picture

奇妙图床，your photo album on the cloud。

项目根目录的`nac.syso` 为windows下获得管理员权限的文件

配置文件设置
1. system 系统设置
   1. env 环境值 线上环境：`public` 开发环境：`develop` 开发环境日志等级为debug模式
   2. Ip 监听地址
   3. addr 端口号
   4. dbType 数据库类型
2. log 日志设置 使用`logrus`
   1. path 日志路径
   2. name 日志文件名
   3. type 日志文件后缀
   4. format 日志格式 `text`或`json`
   5. level 日志等级 默认为`Info`
      1. `Panic`：记录日志，然后panic。
      2. `Fatal`：致命错误，出现错误时程序无法正常运转。输出日志后，程序退出代码为1。
      3. `Error`：错误日志，需要查看原因。
      4. `Warn`：警告信息，提醒程序员注意。
      5. `Info`：关键操作，核心流程的日志。
      6. `Debug`：一般程序中输出的调试信息。
      7. `Trace`：很细粒度的信息，一般用不到。
   6. maxAge 文件最大保存时间，单位为分钟 43200为30天
   7. rotationTime 日志切割时间间隔，单位为分钟 1400为一天
3. mysql 数据库设置
   1. Username 用户名
   2. Password 密码
   3. Path 数据库路径
   4. Port 端口号
   5. Dbname 数据库名称
   6. Config 数据库连接配置
   7. LogMode 数据库日志等级
   
```json
{
  "system": {
    "env": "develop",
    "Ip": "0.0.0.0",
    "addr": "8080",
    "dbType": "mysql"
  },
  "log": {
    "path": "./log/",
    "name": "pic-vue",
    "type": "log",
    "format": "text",
    "level": "info",
    "maxAge": 43200,
    "rotationTime": 1440
  },
  "mysql": {
    "Username": "root",
    "Password": "Abc@123321",
    "Path": "127.0.0.1",
    "Port": "3306",
    "Dbname": "picture",
    "Config": "charset=utf8mb4&parseTime=True&loc=Local",
    "LogMode": "info"
  }
}
```
