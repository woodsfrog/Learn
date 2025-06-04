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
举个例子：
```git
package main

import(
    "fmt"
)

func main() {
    var power int
    power = 9000
    fmt.Printf("It's over %d\n",power)
}

```
这个例子，声明了一个名为power的变量，类型为int。 默认情况，Go会将变量初始化为零值，整数初始化为0，布尔值初始化为false，字符串初始化为“”等。
接下来，我们将9000赋值给power变量，我们可以将前两行合并：`var power int =9000`

但是仍然很长，Go提供了一种短变量声明：`power := 9000`

这种声明也可以适用于函数：
```git
func main() {
    power := getPower()
}

func  getPower() int{
    return 9001
}
```

对于：=的使用，要注意其适用于声明变量并为其赋值，且一个变量在同一个域内不能被声明两次。

我们使用 :=，但在后续赋值时，我们使用赋值运算符 =。这很有道理，但当你需要在两者之间切换时可能会有些棘手。 如果你仔细阅读错误信息，你会发现变量是复数形式。这是因为 Go 允许你同时赋值多个变量（可以使用 = 或 :=）。
```git
func main () {
    power := 1000
    fmt.Printf("default power is %d\n",power)

    name, power := "GoKu",9000
    fmt.Printf("%s's power is over %d\n", name, power)
}
```

目前，请记住当你声明一个变量并使用其零值时使用 `var NAME TYPE`，当你声明并赋值一个值时使用 `NAME := VALUE`，当你给之前声明的变量赋值时使用 `NAME = VALUE`。

## 函数声明
下面是一个指出函数返回多个值的例子。其中有三个函数：一个没有返回值的函数，一个有一个返回值的函数，以及一个有两个返回值的函数。
```git
func log(message string){
}

func add(a int, b int) int {
}

func power(name string) (int, bool){
}
```
我们在最后一个函数输入：
```git
value, exists := power("goku")
if exists == false {
    // handle this error case
}
```

有时，只注意返回值中的一个，在这种情况下，可以讲其他值赋予_：
```git
_, exists := power("goku")
if exists == false {
    // handle this error case
}
```

这不仅仅是一个约定。_, 空标识符，在这一点上特别之处在于返回值实际上并没有被赋值。这让你可以无论返回类型如何，都可以反复使用 _。

然后还有一种更简短地语法：
```git
func add(a,b int) int {

}
```

能够返回多个值是你经常会用到的功能。你也会经常使用 _ 来丢弃一个值。命名返回值和稍微不那么冗长的参数声明并不常见。尽管如此，你很快就会遇到这些内容，因此了解它们是很重要的。




