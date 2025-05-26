# 基础
接下来，将要学习：
- 设置遵循go约定的项目目录
- 启动web服务器并侦听传入的HTTP请求
- 根据请求路径和方法将请求路由到不同的处理程序
- 在路由模式中使用通配符分段
- 向用户发送不同HTTP响应、标头和状态代码
- 以合理可扩展的方式构建项目
- 呈现HTML页面并使用模版继承和保持HTML标记没有重复的样板代码
- 从您的应用程序提供静态的文件，如图像、CSS和JaveScript

## 项目设置和创建模块
在我们编写任何代码之前，您需要在计算机上创建一个 snippetbox 目录，作为此项目的顶级 “home”。我们在整本书中编写的所有 Go 代码，以及其他特定于项目的资源（如 HTML 模板和 CSS 文件）都将位于此处。

本文的项目设置在Projects文件夹下。

### 创建模版
确定项目的路径模块，如果你还不熟悉 Go 模块，你可以把模块路径看作是你项目的规范名称或标识符。

你几乎可以选择任何字符串作为你的模块路径，但需要关注的重要一点是唯一性。为避免将来与其他人的 projects 或 standard library 发生潜在的导入冲突，您需要选择一个全局唯一且不太可能被其他任何内容使用的 module path。在 Go 社区中，一个常见的约定是将模块路径基于您拥有的 URL。

就我而言，对于这个项目来说，一个清晰、简洁且不太可能被其他任何东西使用的模块路径会 snippetbox.alexedwards.net，我将在本书的其余部分使用它。如果可能的话，你应该把它换成你独有的东西。

确定唯一的模块路径后，接下来需要做的是将项目目录转换为模块。

确保你位于项目目录的根目录中，然后运行 go mod init 命令 — 将你选择的模块路径作为参数传入，如下所示：

```git
$ cd $HOME/code/snippetbox
$ go mod init snippetbox.alexedwards.net
```

### 你好世界！
再继续之前，测试一下是否设置正确。在项目目录中创建一个新的main.go，其中包含以下代码：
```git
$ touch main.go
-----------
package main

import "fmt"

func main() {
    fmt.Println("Hello world!")
}
```

使用go run命令，运行并编译当前目录下的代码。
```git
$ go run .
Hello world!
```

### 其他信息
如果你正在创建一个可以被其他人和程序下载和使用的项目，那么最好让你的模块路径等于可以下载代码的位置。

例如，如果您的包托管在 https://github.com/foo/bar 上，则应 github.com/foo/bar 项目的模块路径。

## Web 应用程序基础知识
现在一切都设置正确了，让我们进行 Web 应用程序的第一次迭代。我们将从三个绝对要素开始：
- 先需要一个处理程序。如果您以前使用 MVC 模式构建过 Web 应用程序，则可以将处理程序视为有点像控制器。他们负责执行您的应用程序逻辑以及编写 HTTP 响应标头和正文。
- 第二个组件是路由器（或 Go 术语中的 servemux）。这将存储应用程序的 URL 路由模式与相应处理程序之间的映射。通常，您的应用程序有一个 servemux ，其中包含所有路由。
- 我们最不需要的是 Web 服务器。Go 的一大优点是你可以建立一个 Web 服务器，并将传入的请求作为应用程序本身的一部分来监听。

将这些组件放在 main.go 文件中，以创建一个有效的应用程序。

```git
File: main.go
------------
package main

import (
    "log"
    "net/http"
)

// 定义了一个home处理函数，该函数写入了一个包含"Hello from Snippetbox"作为响应正文的字节
func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from Snippetbox"))
}

func main() {
    // 使用http.NewServeMux()初始化一个新的servemux，然后将home函数注册为"/"URl模式处理程序
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    
    // 打印一条日志，服务器正在启动   
    log.Print("starting server on :4000")
    
    // 使用 http.ListenAndServe（） 函数来启动新的 Web 服务器。我们传入两个参数： 要监听的 TCP 网络地址（在本例中为 “：4000”）和我们刚刚创建的 servemux 。如果 http.ListenAndServe（） 返回一个错误 ，我们使用Fatal（） 函数记录错误消息并退出。
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
```

注：当您运行此代码时，它将启动一个 Web 服务器，侦听本地计算机的端口 4000。每次服务器收到新的 HTTP 请求时，它都会将请求传递给 servemux，然后 servemux 将检查 URL 路径并将请求分派给匹配的处理程序。

保存 main.go 文件，然后尝试使用 go run 命令从终端运行它。

在服务器运行时，打开 Web 浏览器并尝试访问 http://localhost:4000。

如果您返回终端窗口，您可以通过按键盘上的 Ctrl+C 来停止服务器。

### 其他信息
网络地址
您传递给 http.ListenAndServe（） 的格式应为 “host：port”。如果你省略了主机（就像我们对 “：4000” 所做的那样），那么服务器将监听所有计算机的可用网络接口。通常，如果您的计算机有多个网络接口，并且您只想侦听其中一个，则只需在地址中指定一个主机。

在其他 Go 项目或文档中，你有时可能会看到使用命名端口（如 “：http” 或 “：http-alt” 而不是数字）编写的网络地址。如果您使用命名端口，则 http.ListenAndServe（） 函数将在启动服务器时尝试从 /etc/services 文件中查找相关端口号，如果找不到匹配项，则返回错误。

使用go run
在开发过程中，go run 命令是试用代码的便捷方式。它本质上是一个快捷方式，用于编译您的代码，在您的 /tmp 目录中创建一个可执行二进制文件，然后一步运行此二进制文件。

它接受以空格分隔的 .go 文件列表、特定包的路径（其中 . 字符表示当前目录）或完整的模块路径。对于我们目前的应用程序，以下三个命令都是等效的：

```git
$ go run .
$ go run main.go
$ go run snippetbox.alexedwards.net
```

## 路由请求
拥有只有一个路由的 Web 应用程序并不是很令人兴奋......或有用！让我们再添加几个路由，以便应用程序开始形成，如下所示：

<table style="width: 100%;">
    <tr>
        <th align="center">Rounte pattern</th>
        <th align="center">Handle</th>
        <th align="center">Action</th>
    </tr>
    <tr>
        <th align="center">/</th>
        <th align="center">home</th>
        <th align="center">显示首页</th>
    </tr>
    <tr>
        <th align="center">/snippet/view</th>
        <th align="center">snippetView</th>
        <th align="center">显示特定代码段</th>
    </tr>
    <tr>
        <th align="center">/snippet/Create</th>
        <th align="center">snippetCreate</th>
        <th align="center">显示用于创建新代码段的表单</th>
    </tr>
</table>

重新打开main.go文件，并以如下方式更新它：

```git
File: main.go
------------
package main

import (
    "log"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from Snippetbox"))
}

// 添加一个snippetView处理程序函数
func snippetView(w http.ResponseWriter,r*http.Request){
    w.Write([]byte("Display a specific snippet..."))
}

// 添加sniperCreate处理程序函数
func snippetCreate(w http.ResponseWriter,r*http.Request){
    w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)
    
    log.Print("starting server on :4000")
    
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
```

使用go run启动服务器

接下来可以在web浏览器中访问以下链接
- http://localhost:4000/snippet/view
- http://localhost:4000/snippet/create

### 路由模式的末尾斜杠
重要的是要知道 Go 的 servemux 具有不同的匹配规则，具体取决于路由模式是否以尾部斜杠结尾。

我们的两个新路由模式 — “/snippet/view” 和 “/snippet/create” — 不以尾部斜杠结尾。当模式没有尾部斜杠时，只有当请求 URL 路径与模式完全匹配时，它才会被匹配（并调用相应的处理程序）。

当路由模式以尾部斜杠结尾时（如 “/” 或 “/static/” ），它称为子树路径模式。每当请求 URL 路径的开头与子树路径匹配时，就会匹配（并调用相应的处理程序）子树路径模式。如果它有助于您的理解，您可以将子树路径视为有点像它们在末尾有一个通配符，例如 “/**” 或 “/static/**”。

这有助于解释为什么 “/” 路由模式就像一个 catch-all。该模式本质上意味着匹配单个斜杠，后跟任何内容（或根本没有）。

### 限制分支路径
为了防止子树路径模式的行为就像它们在末尾有通配符一样，您可以将特殊字符序列 {$} 附加到模式的末尾，例如“/{$}”或“/static/{$}”。

因此，如果你有路由模式 “/{$}”，它实际上意味着匹配一个斜杠，后面没有其他任何东西。它只会匹配 URL 路径正好为 / 的请求。



