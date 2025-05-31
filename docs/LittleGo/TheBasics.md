# Chapter 1 The Basics
Go 是一种编译的静态类型语言，具有 Clike 语法和垃圾回收功能。

## 静态类型
静态类型意味着变量必须具有特定类型（int、string、bool、[]byte 等）。这是 通过在声明变量时指定类型来实现，或者在许多情况下，让编译器推断类型

## C-like 语法
Go 不仅在语法方面，而且在目的方面都比 C# 或 Java 更接近 C。

## Garbage Collected

## 运行go代码
`go run xxx.go` 运行go代码命令。

go run 是一个方便的命令，用于编译和运行您的代码。它使用一个临时目录来构建程序，执行它，然后自行清理。
`go run --work xxx.go`可以看到编译步骤。

要显式编译代码，请使用 go build。这将生成一个可执行的 main，您可以运行它。在 Linux / OSX 上，不要忘记您需要在executable 替换为 dotslash，因此您需要键入 ./main
`go build xxx.go`
开发时，您可以使用 go run 或 go build。但是，当您部署代码时，您需要部署通过 go build 获取二进制文件并执行该二进制文件。

## main
在go语言中，程序的入口点必须是 main 包中名为 main 的函数。

## imports
Go 有很多内置函数，比如 println，可以不用引用就使用。也有一些需要导入使用第三方库。在go中，导入keyword用于声明文件代码使用的包。
```git
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
    os.Exit(1)
    }
    fmt.Println("It's over", os.Args[1])
}
```
接下来，使用`go run main.go 9000`运行。

我们现在使用 Go 的两个标准包：fmt 和 os。我们还引入了另一个内置函数 len。 len返回字符串的大小，或字典中值的数量，或者，正如我们在这里看到的，一个数组。 

Go 对导入包很严格。如果你导入了一个包但不使用它，它不会编译。

## 变量和声明
在 Go 中处理变量声明和赋值最明确的方法也是最冗长的




