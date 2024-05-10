# Elix

> 一个Kook机器人的事件代理服务

本项目创建的初衷是为了更方便地将kook的机器人的事件转发到下游连接的clients中(类似于[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)
的事件上报服务)，采用websocket方式进行网络通信。

### 功能

- 作为Kook的代理服务，对接下游的处理服务进行处理(目前主要为[kook-api](https://github.com/Aimenry/KookAPI)插件提供api)
- 开箱即用的maimai查分服务

### 使用姿势

1. 下载对应平台的可执行程序，执行

   ```shell
   ./elix
   ```

1. 如果在elix同级目录已经存在配置文件，将会读取对应配置文件，否则将会创建新的配置文件，请修改配置文件中的`token`为自己的机器人的token值

2. 重新执行步骤1，出现以下日志表示已经连接成功，此时去Kook中查看机器人状态应已经处于`online`状态

   ```log
   # 出现此日志表示机器人已经开始正常工作
    INFO[2024-05-11 02:42:19] Heart beat checker inited
   ```

### 配置

```json5
{
  "token": "Your kook bot token",
  "compress": true, // 是否启用zlib压缩传输,建议保持不变
  "api_server_port": 9001, // http api服务的端口
  "ws_proxy_server_port": 9000 // ws代理服务的端口
}
```

### 其他

本项目使用[kook-go](https://github.com/Aimenry/kook-go)作为连接Kook的Sdk进行开发

> 本项目目前还在开发中，如果感觉对自己有帮助还请点个star⭐️，有任何建议或者意见欢迎提issue~**