# 开始
在这一章，完成事项：
- 创建项目框架目录结构
- 建立HTTP服务器并侦听请求
- 引入合理的模式管理配置，使用依赖项注入使依赖项可以用于我们的处理程序（？）
- 使用httprouter帮助API端点实现RESful结构

## 项目设立及骨架结构
首先创建一个greenlight目录作为该项目主页，使用mkdir创建(可以随意选择位置，pwd可查看当前位置)：
``` git
$ mkdir -p $HOME/Projects/greenlight
```
切换到该目录（cd命令），使用go mod init命令作为项目启动模组

