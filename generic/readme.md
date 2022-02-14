Go1.18 泛型就会正式释出
Go1.17 泛型代码已合入 master 分支。
需要在 run 和 build 的命令执行时指定 -G 标识

http://www.codebaoku.com/it-go/it-go-233126.html


函数要支持这种泛型行为，需要有2个前提条件：

对于函数而言，需要一种方式来声明这个函数到底支持哪些类型的参数
对于函数调用方而言，需要一种方式来指定传给函数的到底是int类型的map还是float类型的map


每个类型参数都有一个类型限制(type constraint)，类型限制就好比类型参数的meta类型，每个类型限制会指明函数调用时该类型参数允许的类型实参。
https://segmentfault.com/a/1190000041140246


// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

fmt.Printf("Generic Sums: %v and %v\n",
    SumIntsOrFloats[string, int64](ints),
    SumIntsOrFloats[string, float64](floats))
在这段代码里：

调用了上面定义的泛型函数，传递了2种类型的map作为函数的实参。
函数调用时指明了类型实参(方括号[ ]里面的类型名称)，用于替换调用的函数的类型实参。

在接下来的内容里，你会经常看到调用函数时，会省略掉类型实参，因为Go通常(不是一定)可以根据你的代码推断出类型实参。

类型实参的自动推导并不是永远可行的。比如，你调用的泛型函数没有形参，不需要传递实参，那编译器就不能根据实参自动推导，需要在函数调用时在方括号[]里显示指定类型实参。

声明类型限制(type constraint)
泛型函数里的类型限制以接口(interface)的形式做定义，这样类型限制就可以在很多地方被复用。声明类型限制可以帮助精简代码，特别是在类型限制很复杂的场景下。

我们可以声明一个类型限制(type constraint)为接口(interface)类型。这样的类型限制可以允许任何实现了该接口的类型作为泛型函数的类型实参。例如，你声明了一个有3个方法的类型限制接口，然后把这个类型限制接口作用于泛型函数的类型限制，那函数调用时的类型实参必须要实现了接口里的所有方法。


在main函数上面，import语句下面，添加如下代码用于声明一个类型限制

type Number interface {
    int64 | float64
}



https://segmentfault.com/a/1190000041174189


generic types(泛型类型)
类型参数除了用于泛型函数之外，还可以用于Go的类型定义，来实现泛型类型(generic types)。


generic types(泛型类型)
类型参数除了用于泛型函数之外，还可以用于Go的类型定义，来实现泛型类型(generic types)。

看如下代码示例，实现了一个泛型二叉树结构

type Tree[T interface{}] struct {
    left, right *Tree[T]
    data T
}

func (t *Tree[T]) Lookup(x T) *Tree[T] 

var stringTree Tree[string]



type sets(类型集)
类型参数的类型限制约定了该类型参数允许的具体类型。

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
type Ordered interface {
  Integer | Float | ~string
}

类型参数列表不能用于方法，只能用于函数。

type Foo struct {}

func (Foo) bar[T any](t T) {}

|: 表示取并集。比如下例的Number这个interface可以作为类型限制，用于限定类型参数必须是int，int32和int64这3种类型。

~T: ~ 是Go 1.18新增的符号，~T表示底层类型是T的所有类型。~的英文读作tilde。

比如下例的AnyString这个interface可以作为类型限制，用于限定类型参数的底层类型必须是string。string本身以及下面的MyString都满足AnyString这个类型限制。

type AnyString interface{
   ~string
}
type MyString string

类型限制有2个作用：

用于约定有效的类型实参，不满足类型限制的类型实参会被编译器报错。
如果类型限制里的所有类型都支持某个操作，那在代码里，对应的类型参数就可以使用这个操作。

constraint literals(类型限制字面值)
type constraint既可以提前定义好，也可以在type parameter list里直接定义，后者就叫constraint literals。

[S interface{~[]E}, E interface{}]

[S ~[]E, E interface{}]

[S ~[]E, E any]
几个注意点：

可以直接在方括号[]里，直接定义类型限制，即使用类型限制字面值，比如上例。
在类型限制的位置，interface{E}也可以直接写为E，因此就可以理解interface{~[]E}可以写为~[]E。
any是Go 1.18新增的预声明标识符，是interface{}的别名。

constraints包
constraints包定义了一些常用的类型限制，整个包除了测试代码，就1个constraints.go文件，


Type inference(类型推导)

Go泛型有2种类型推导：

function argument type inference: deduce type arguments from the types of the non-type arguments.

通过函数的实参推导出来具体的类型。比如上面例子里的m2 = min(a, b)，就是根据a和b这2个函数实参

推导出来T是float64。

constraint type inference: inferring a type argument from another type argument, based on type parameter constraints.

通过已经确定的类型实参，推导出未知的类型实参。下面的代码示例里，根据函数实参2不能确定E是什么类型，但是可以确定S是[]int32，再结合类型限制里S的底层类型是[]E，可以推导出E是int32，int32满足constraints.Integer限制，因此推导成功。

type Point []int32

func ScaleAndPrint(p Point) {
  r := Scale(p, 2)
  fmt.Println(r)
}

func Scale[S ~[]E, E constraints.Integer](s S, c E) S {
  r := make(S, len(s))
  for i, v := range s {
    r[i] = v * c
  }
  return r
}
类型推导并不是一定成功，比如类型参数用在函数的返回值或者函数体内，这种情况就必须指定类型实参了。

func test[T any] () (result T) {...}
func test[T any] () {
  fmt.Println(T)
}


什么时候使用泛型？
需要使用slice, map, channel类型，但是slice, map, channel里的元素类型可能有多种。
通用的数据结构，比如链表，二叉树等。

什么时候不要使用泛型?
只是单纯调用实参的方法时，不要用泛型。
当函数或者方法或者具体的实现逻辑，对于不同类型不一样时，不要用泛型。比如encoding/json这个包使用了reflect，如果用泛型反而不合适。
Go语言里interface和refelect可以在某种程度上实现泛型，

https://segmentfault.com/a/1190000041174189


为啥当前Go泛型不好实现泛型方法?
package p3不知道p1.S类型，整个程序中如果也没有其它地方调用p1.S.Identity,依照现在的Go编译器的实现，是没有办法为p1.S.Identity[int]生成对应的代码的。

Facilitator模式 
函数实现让人感觉到一种无力感，一种缺乏归宿感，一种没有对象的感觉，而这种实现呢，生成了特定类型的Querier[T],All方法就有泛型的感觉了(虽然实际是Receiver泛型)

https://colobu.com/2021/12/22/no-parameterized-methods/