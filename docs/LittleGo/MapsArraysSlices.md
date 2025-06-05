# Chapter 3 Maps,Arrays and Slices
到目前为止，我们已经看到了一些简单的类型和结构。现在该看看数组、切片和映射了。

## 数组
如果你来自 Python、Ruby、Perl、JavaScript 或 PHP（以及其他更多语言），你可能习惯于使用动态数组。这些数组在添加数据时会自动调整大小。但在 Go 语言中，像许多其他语言一样，数组是固定的。声明一个数组需要我们指定大小，一旦指定大小后，它就不能再增长了：
```git
var scores [10]int
scores[0] = 339
```
上述数组可以存储最多10个分数，使用索引scores[0]到scores[9]。尝试访问数组范围外的索引会导致编译器或运行时错误。

我们可以用值初始化数组：
```git
scores := [4]int{9001, 9333, 212, 33}
```

我们可以使用 len 获取数组的长度。range 可以用来遍历它：
```git
for index, value := range scores {

}
```

数组高效但僵硬。我们往往不知道将要处理的元素数量。为此，我们使用切片。

## 切片
在Go中，将会很少，甚至几乎不用数组。相反，将会使用切片，切片是一种轻量级的结构，它封装并表示数组的一部分/

第一个是我们在创造数组时的一种轻微的变体：
`source ：= []int{1,4,293,4,1}`

与数组声明不同的是，切片不是在方括号内声明长度，为了理解其中的区别，我们看一看另一种创造切片的方法，使用make：
`source ：= make([]int，10)`

我们使用make而不是new，是因为创造一个切片不仅仅只是分配内存（new的作用）。具体而言我们需要为底层数组分配内存而初始化切片。

在上面的例子中，我们初始化了一个长度和容量为10的切片。长度是切片的大小，容量是底层数组的大小。使用make可以分别定义这两者：
`source ：= make([]int，0，10)`

这创造了一个长度为0但容量为10的切片。（如果你注意到了，会发现‘make’和‘len’被重载了。Go是一种让开发者沮丧的语言）

为了更好的理解长度和容量之间的交互，我们来看一些例子：
```git
func main(){
    scores :=make([]int，0，10)
    scores[7] = 9033
    fmt.Println(scores)
}
```

第一个例子崩溃了，为什么？因为切片长度是0。底层数组有10个元素，但我们需要显式地扩展切片来访问元素。扩展切片的一个方法是使用append：
```git
func main(){
    scores :=make([]int, 0, 10)
    scores = append(scores, 5)
    fmt.Println(scores)
}
```

但这样会改变我们原始代码的意图。将长度为0的切片追加元素设置为第一个元素。不知出于什么原因，我们的崩溃代码想要设置索引为7的元素。我们可以重新切片：
```
func main(){
    scores := make([]int, 0, 10)
    scores = scores[0:8]
    scores[7] = 9033
    fmt.Println(scores)
}
```

我们能将切片Resize到多大？最多到其容量，这种情况下是10。你可能在想这实际上并没有解决数组的固定长度问题。结果发现append非常特别。如果底层数组满了，它会创建一个新的更大数组并将值复制过去（这正是PHP、Python、Ruby、JavaScript等语言中的动态数组的工作方式）。这就是为什么在上面使用append的例子中，我们需要将append返回的值重新赋值给scores变量：append如果原数组没有更多空间，可能会创建一个新的值。

如果我告诉你Go使用2倍算法增长数组，你能猜出下面的代码会输出什么吗？
```
func main() {
    scores := make([]int, 0, 5)
    c := cap(scores)
    fmt.Println(c)
    
    for i := 0; i < 25; i++ {
        scores = append(scores, i)
        // if our capacity has changed,
        // Go had to grow our array to accommodate the new data
        if cap(scores) != c {
            c = cap(scores)
            fmt.Println(c)
        }
    }
}
```

分数的初始容量是5。为了容纳25个值，它将需要扩展3次，容量分别为10，20和最后40。



