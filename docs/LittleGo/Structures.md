# Chapter 2 Structures
Go 不像 C++、Java、Ruby 和 C# 那样是一种面向对象（OO）语言。它没有对象和继承，因此也没有与 OO 相关的许多概念，如多态性和重载。

Go 有的是结构体，可以与方法关联。Go 还支持一种简单但有效的组合形式。总体来说，这会导致代码更简单，但有时你会怀念 OO 提供的一些功能。（值得一提的是，组合优于继承是一个老生常谈的话题，而 Go 是我使用的第一种在该问题上采取坚定立场的语言。）

虽然 Go 的面向对象编程方式可能与你习惯的不同，但你会发现结构体的定义与类的定义有很多相似之处。一个简单的例子是以下的 Saiyan 结构体：
```git
type Saiyan struct{
    Name string
    Power int
}
```
我们很快就会看到如何向这个结构中添加一个方法，就像你在类中添加方法一样。在我们做那之前，我们必须再深入了解一下声明。

## 声明和初始化
当我们最初研究变量和声明时，我们只关注内置类型，比如整数和字符串。现在我们谈到结构体了，我们需要将讨论扩展到包括指针。

创建我们结构体的一个值的最简单方式是：
```git
goku := Saiyan{
    Name: "GoKu",
    Power: 9000,
}
```
注意：上述结构中的逗号是必需的。如果没有它，编译器将给出错误。你将欣赏到这种一致性，尤其是如果你使用的是一个语言或格式，它强制相反的规则。

我们不需要设置所有字段，甚至不需要设置任何字段。这两种情况都是有效的：
```git
goku := Saiyan{}

// or

goku := Saiyan{Name: "GoKu"}
goku.power = 9000
```
就像未分配值的变量有一个零值一样，字段也是如此。

此外，你可以省略字段名，并依赖于字段声明的顺序（尽管为了清晰起见， 你只应该在字段较少的结构体中这样做）：
```goku := Saiyan{"Goku", 9000}```

以上所有的例子都声明了一个变量goku并给它赋了一个值。

不过，很多时候我们不想直接与我们的值关联一个变量，而是希望有一个指向我们值的变量。指针是一个内存地址；它是找到实际值的位置。这是一种间接性。粗略地说，这就像在房子和去房子的路线之间的区别。

我们为什么要指向值的变量而不是实际的值呢？这取决于Go是如何将参数传递给函数的：通过拷贝。知道了这一点，下面的代码会打印什么？
```git
func main() {
    goku := Saiyan{"Goku", 9000}
    Super(goku)
    fmt.Println(goku.Power)
}

func Super(s Saiyan) {
    s.Power += 10000
}
```

这个的答案是9000，为什么不是19000，因为super对原始的goku进行了副本修改，因此，super的更改没有反映在调用中，应改为这样，让其指向我们的值的指针：
```git
func main() {
    goku := &Saiyan{"Goku", 9000}
    Super(goku)
    fmt.Println(goku.Power)
}

func Super(s *Saiyan) {
    s.Power += 10000
}
```
我们做了两个更改。第一个是使用 & 运算符来获取我们值的地址（它被称为地址运算符）。接下来，我们改变了 Super 所期望的参数类型。它以前期望一个 Saiyan 类型的值，但现在期望一个 *Saiyan 类型的地址，其中 *X 表示 X 类型的指针。显然，Saiyan 和 *Saiyan 之间存在某种关系，但它们是两种不同的类型。

请注意，我们仍然向 Super 传递了 goku 值的一个副本，只是碰巧 goku 的值已经变成了一个地址。这个副本与原始地址相同，这就是间接引用带给我们的。可以将其视为复制去餐馆的路线。你拥有的是一个副本，但它仍然指向与原始地址相同的餐馆。

我们可以通过尝试改变它指向的地方来证明这是一个副本（虽然你通常不会真的想这样做）：
```git
func main() {
    goku := &Saiyan{"Goku", 9000}
    Super(goku)
    fmt.Println(goku.Power)
}

func Super(s *Saiyan) {
    s = &Saiyan{"Gohan", 1000}
}
```
上面再次打印了9000。这是一些语言的行为方式，包括Ruby、Python、Java和C#。在某种程度上，C#只是让这种事实变得明显。

也应该很明显的是，复制指针比复制复杂结构要便宜。在64位机器上，指针是64位大小。如果我们有一个包含许多字段的结构，创建副本可能会很昂贵。然而，指针的实际价值在于它们让你能够共享值。我们是想让Super修改goku的一个副本，还是修改共享的goku值本身？

这并不是说你总是想要一个指针。在这一章结束时，在我们看到我们可以用结构做什么之后，我们将重新审视 指针与值 的问题。

## 结构上的功能

```git
type Saiyan struct {
    Name string
    Power int
}

func (s *Saiyan) Super() {
    s.Power += 10000
}
```

在上述代码中，我们说类型*Saiyan 是 Super 方法的接收者。我们像这样调用 Super：
```
goku := &Saiyan{"Goku", 9001}
goku.Super()
fmt.Println(goku.Power) // will print 19001
```

## 构造函数
结构没有构造函数。相反，你创建一个函数返回所需类型的实例（就像一个工厂）：
```go
func NewSaiyan(name string, power int) *Saiyan {
    return &Saiyan{
        Name: name,
        Power: power,
    }
}
```
这种模式让很多开发者感到不舒服。一方面，这只是一个小的语法变化；另一方面，确实感觉不太模块化。

我们的工厂不需要返回一个指针；这绝对是有效的：

```go
func NewSaiyan(name string, power int) Saiyan {
    return Saiyan{
        Name: name,
        Power: power,
    }
}
```

## New
尽管没有构造函数，Go 语言确实有一个内置的新函数（new），用于分配类型所需的空间。new(X) 的结果与 &X{} 相同：
```go
goku := new(Saiyan)
// same as
goku := &Saiyan{}
```

你是使用哪种方式取决于你，但你会发现，当人们需要初始化字段时，大多数人都更喜欢后者，

因为这种方式通常更容易阅读：
```
goku := new(Saiyan)
goku.name = "goku"
goku.power = 9001

//vs

goku := &Saiyan {
    name: "goku",
    power: 9000,
}
```
无论你选择哪种方法，如果你遵循上述工厂模式，你可以屏蔽掉代码中其他部分对任何分配细节的了解和担忧。


## 结构中的字段
到目前为止的例子中，赛亚人的字段包括Name和Power，类型分别为string和int。

字段可以是任何类型——包括我们尚未探索的其他结构和类型，例如数组、映射、接口和函数。

例如，我们可以扩展赛亚人的定义：
```git
type Saiyan struct (
    Name string
    Power int
    Father *Saiyan
)
```

我们可以通过以下方式初始化：
```git
gohan := &Saiyan{
    Name: "Gohan",
    Power: 1000,
    Father: &Saiyan {
        Name: "Goku",
        Power: 9001,
        Father: nil,
    },
}
```

## 组成
Go 支持组合，即在一个结构中包含另一个结构的行为。在某些语言中，这被称为一个 trait 或 mixin。没有显式的组合机制的语言总是可以通过长方法来实现这一点。在 Java 中，可以通过继承来扩展结构，但在这种选项不可用的情况下，一个 mixin 可以这样编写：
```javascript
public class Person {
    private String name;

    public String getName() {
        return this.name;
    }
}

public class Saiyan {
    // Saiyan is said to have a person
    private Person person;

    // we forward the call to person
    public String getName() {
        return this.person.getName();
    }
    ...
}
```

这可能会变得相当乏味。Person的每种方法都需要在 Saiyan 中重复。Go 避免了这种乏味：
```go
type Person struct {
    Name string
}

func (p *Person) Introduce() {
    fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
    *Person
    Power int
}

// and to use it:
goku := &Saiyan{
    Person: &Person{"Goku"},
    Power: 9001,
}
goku.Introduce()
```

赛亚人的结构有一个类型为*Person的字段。因为我们没有给它显式指定字段名，所以可以隐式访问组成类型的字段和函数。然而，Go编译器还是给了它一个字段名，考虑一下完全有效的：
```git
goku := &Saiyan{
    Person: &Person{"Goku"},
}
fmt.Println(goku.Name)
fmt.Println(goku.Person.Name)
```
上面的代码都会输出“悟空”。

组合比继承更好吗？很多人认为这是一种更稳健的代码共享方式。使用继承时，你的类会紧密耦合到超类上，你往往会关注层次结构而不是行为。

## Overloading
虽然过载不是特定于结构体的，但值得提及一下。简单来说，Go 并不支持过载。因此，你会看到（并编写）很多看起来像 Load、LoadById、LoadByName 等的函数。

然而，因为隐式组合其实只是一个编译器技巧，我们可以通过“覆盖”组合类型的功能来实现类似的效果。例如，我们的 Saiyan 结构体可以有自己的 Introduce 函数：
```git
func (s *Saiyan) Introduce() {
    fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}
```

组成的版本总是可以通过 s.Person.Introduce() 获取。

## 指针与值

当你编写 Go 代码时，自然会问自己这是应该是一个值，还是一个值的指针？有两个好消息。首先，无论我们讨论的是以下哪一项，答案都是相同的：

• 局部变量赋值
• 结构体中的字段
• 函数的返回值
• 函数的参数
• 方法的接收者

其次，如果你不确定，就使用指针。

正如我们之前看到的，传递值是使数据不可变的一种很好的方式（函数对其所做的更改不会反映在调用代码中）。有时，这正是你想要的行为，但更常见的是，这并不是你想要的。

即使你无意更改数据，也要考虑创建大型结构体副本的成本。相反，你可能有小型结构体，例如：
```git
type Point struct {
    X int
    Y int
}
```

在这种情况下，复制结构的成本可能通过可以直接访问X和Y而无需任何间接访问来抵消。

再次强调，这些都是非常微妙的情况。除非你在成千上万甚至可能是数以万计的此类点上进行迭代，否则你不会注意到任何差异。




