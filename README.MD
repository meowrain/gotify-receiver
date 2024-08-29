# gotify-client

一个简易的 gotify windows 客户端

用于实时接收 gotify 推送的通知消息, 并在 windows 的通知栏显示.

## 使用方法

修改`config.example.toml`为`config.toml`或者复制一份

```bash
cp config.example.toml ./config.toml
```

填入配置信息,启动程序

```bash
$ ./main.exe
2024-08-29 23:22:22.769 INFO    client/client.go:77     url wss://gotify.meowrain.cn/stream?token=CeB1XI.m1gEt33y
```
配合 [winsw](https://github.com/winsw/winsw) 使用
即可 作为 windows 服务运行, 并实现开机自启

# CHANGES

## [1.0.1] - 2024-08-29

### 添加功能

- 短信验证码解析
- 自动复制验证码到剪切板中

# LINCENSE 

本项目是 [gotify-client](https://gitee.com/shikong-sk/gotify-client) 的一个分支，该项目使用 Apache License 2.0 许可证。
