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

运行此命令需要一个模块路径，作为项目的唯一标识符。本项目以greenlight.alexedwards.net作为模块路径。

```angular2html
$ go mod init greenlight.alexedwards.net
```

运行后在该文件夹下生成go.mod文件
- 当项目目录的根目录中有有效的go.mod文件，您的项目就是一个模块。
- 当您在项目目录工作并使用go get下载依赖项时，依赖项的确切版本会记录在go.mod文件中。
- 当在项目运行或构件代码时，Go会使用go.mod文件列举的确切依赖项。如果本地没有，会自动下载。
- 定义了模块路径，用作项目中包的根导入路径的标识符。

## 生成框架目录结构
接下来继续运行以下命令来生成高级骨架结构：

``` git
$ mkdir -p bin cmd/api internal migrations remote
$ touch Makefile
$ touch cmd/api/main.go
```

此时可以注意到grennlight项目目录发生了变化
- bin目录将包含编译的应用程序二进制文件，可以部署到生产服务器。
- cmd/api目录包含项目API应用程序特定代码。作为运行服务器、读取和写入HTTP请求及管理身份验证的代码。
- internal目录包含API使用的各种辅助包。用于数据库交互、进行数据验证、发送电子邮件等代码。基本上，任何不特定于应用程序且可能被重用的代码都将放在此处。
- migrations将包含数据库SQL迁移文件。
- remove目录包含服务器的配置文件和设置脚本
- go.mod声明项目的依赖项、版本和模块路径。
- Makefile包含自动执行常见管理任务的tasks，例如Go代码，构建二进制文件和执行数据库迁移。

注：internal在Go中具有特殊含义和行为：位于该目录下的任何包只能由内部目录的父目录内的代码导入。
   即，防止其他代码库导入和依赖我们内部目录中的包，即使在GitHub等地方公开可用。

## Hello world！
打开cmd/api/main.go文件并添加代码：

```git
File:cmd/api/main.go
----------------
package main

import "fmt"

func main(){
    fmt.Println("Hello,world!")
}
```

注：Go 是区分大小写的语言，注意函数大小写问题

保存此文件，在终端使用go run命令编译并执行cmd/api中的代码。
```git
$ go run ./cmd/api    
```

## 基本HTTP服务器

本节进行启动和运行HTTP服务器的学习。

首先，我们将服务器配置为只有一个终端节点：/v1/healthcheck。此端点返回有关我们API的一些基本信息，包含其当前版本号和作业环境。

<table style="width: 100%">
    <tr>
        <th align="center">URl Pattern</th>
        <th align="center">Handler</th>
        <th align="center">Action</th>
    </tr>
    <tr>
        <th align="center">URl Pattern</th>
        <th align="center">healthcheckHandle</th>
        <th align="center">显示程序信息</th>
    </tr>
</table>

打开cmd/api/main.go文件将‘Hello world！'替换为以下代码：
```git
File:cmd/api/main.go
----------------
package main

import (
    "flag"
    "fmt"
    "log/slog"
    "net/http"
    "os"
    "time"
)

// 声明一个包含应用程序版本号的字符串
const version = "1.0.0"

// 定义了一个config结构来保存我们的所有配置设置。
type config struct {
    port int
    env string
}

// 定义了一个应用程序结构体保存HTTP处理程序和中间的依赖库
type application struct {
    config config
    logger *slog.Logger
}

func main(){
    // 声明config结构实例
    var cfg config

    // 将port env命令行的值读取到config中，使用默认端口4000和development环境。
    flag.InVar(&cfg.port,"port",4000,"API server port")
    flag.StringVar(&cfg.env,"env","development","Environment (development|staging|production)")
    flag.Parse()

    // 初始化一个记录器，将日志写入标准的out系统
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil)

    // 声明应用程序的实体，其中包含config结构和记录器
    app := &application{
        config: cfg,
        logger: logger,
    }

    // 声明一个新的servemux并添加一个请求，该请求分发给/v1/healthcheck，后续章节创建
    mux := http.NewServeMux()
    mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

    // 声明一个HTTP服务器，该服务器侦听配置文件给的端口，进行一些合理的超时设置，并且写入错误记录器
    srv := &http.Server{
        Addr:            fmt.Sprintf(":%d", cfg.port),
        Handler:         mux,
        IdleTimeout:     time.Minute,
        ReadTimeout:     5 * time.Second,
        WriteTimeout:    10 * time.Second,
        ErrorLog:        slog.NewLogLogger(logger.Handler(), slog.LevelError),
    }

    // 开启HTTP服务
    logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

    err := srv.ListenAndServe()
    logger.Error(err.Error())
    os.Exit(1)
}
```

## 创建healthcheck处理程序
接下来创建healthcheckHandler 方法来响应 HTTP 请求。返回一个包含三条信息的纯文本响应。
- 固定的“status： available” 字符串。
- 硬编码版本常量中的 API 版本。
- env 命令行标志中的作环境名称。

创建一个新的cmd/api/healthcheck.go 文件：
```git
$ touch cmd/api/healthcheck.go
```

然后添加以下代码：
```git
File：md/api/healthcheck.go
-------------
package main

import (
    "fmt"
    "net/http"
)

// 该程序写入一个纯文本响应，包含有关应用程序状态、环境和版本信息。
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "status: available")
    fmt.Fprintf(w, "environment: %s\n", app.config.env)
    fmt.Fprintf(w, "version: %s\n", version)
}
```
healthcheckHandler是作为我们程序结构体上的一个方法实现的。
这是一种有效且惯用的方法，无需求助于全局变量或闭包即可使依赖项可供我们的处理程序使用——healthcheckHandler 需要的任何依赖项都可以简单地作为字段包含在应用程序结构体中.

## 示范
通过执行cmd/api包中的代码，可以看到一条日志消息，确定HTTP服务器正在运行。类似于：
```git
% go run ./cmd/api                     
time=2025-05-24T09:20:29.292+08:00 level=INFO msg="starting server" addr=:4000 env=development
```

当服务器运行时，可以通过web浏览器访问http://localhost:4000/v1/healthcheck，可以收到相关响应。

或者可以通过curl从终端发出请求：
```git
% curl -i localhost:4000/v1/healthcheck
HTTP/1.1 200 OK
Date: Sat, 24 May 2025 01:20:32 GMT
Content-Length: 58
Content-Type: text/plain; charset=utf-8

status: available
environment: development
version: 1.0.0
```
注：上述命令中的 -i 标志指示 curl 显示 HTTP 响应标头以及响应正文。

## 其他版本
API版本控制
支持实际业务和用户的 API 通常需要随着时间的推移更改其功能和端点 — 有时以向后不兼容的方式。因此，为避免客户出现问题和混淆，最好始终实现某种形式的 API版本控制

其中最常见两种方式：
1. 在所有 URL 前面加上您的 API 版本，例如 /v1/healthcheck 或 /v2/healthcheck。
2. 通过在请求和响应上使用自定义 Accept 和 Content-Type 标头来传达 API 版本，例如 Accept： application/vnd.greenlight-v1 。

从HTTP语义角度，使用标头传达API更为纯粹。
从用户体验来看，使用URL前缀更好一些，一目了然。

在该教程使用URL路径添加前缀来进行API版本控制。

## API终端节点和RESTful路由
在接下来的几节，逐步构建我们的API，端点如下所示：
<table style="width: 100%;">
    <tr>
        <th align="center">Method</th>
        <th align="center">URL Pateern</th>
        <th align="center">Action</th>
    </tr>
    <tr>
        <th align="center">GET</th>
        <th align="center">/v1/healthcheck</th>
        <th align="center">展示应用运行情况和版本信息</th>
    </tr>
    <tr>
        <th align="center">GET</th>
        <th align="center">/v1/movies</th>
        <th align="center">显示所有电影的详细信息</th>
    </tr>
    <tr>
        <th align="center">POST</th>
        <th align="center">/v1/movies</th>
        <th align="center">创造新的影片</th>
    </tr>
    <tr>
        <th align="center">GET</th>
        <th align="center">/v1/movies/:id</th>
        <th align="center">显示特定电影的详细信息</th>
    </tr>
    <tr>
        <th align="center">PUT</th>
        <th align="center">/v1/movies/:id</th>
        <th align="center">更新特定电影的详细信息</th>
    </tr>
    <tr>
        <th align="center">DELETE</th>
        <th align="center">/v1/movies/:id</th>
        <th align="center">删除特定电影</th>
</table>


