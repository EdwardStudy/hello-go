# Pointer

## 1. operations

指针操作涉及到两个操作符——&和*。

### &

`&`是取址，当地址操作符`&`被应用到一个值上时会取出指向该值的*指针值*。

### *

`*`是取值，当地址操作符`*`被应用到一个*指针值*上时会取出该指针指向的那个值。

当`*`出现在一个类型之前（如`*Person`和`*[3]string`）时就不能被看做是操作符了，而应该被视为一个符号。
如此组合而成的标识符所表达的含义是作为第二部分的那个类型的*指针类型*

## 2. 指针方法

只要一个方法的接收者类型是其所属类型的指针类型而不是该类型本身，那么我就可以称该方法为*指针方法*。

``` go
func (person *Person) Grow() {
    person.Age++
}
```

相对的，如果一个方法的接收者类型就是其所属的类型本身，那么我们就可以把它叫做*值方法*。

``` go
func (person Person) Grow() {
    person.Age++
}
```

不同：值方法的接收者标识符所代表的是该方法当前所属的那个值的一个副本，而不是该值本身。而指针方法的副本还是指向原值。

``` go
p := Person{"Robert", "Male", 33, "Beijing"}
p.Grow()
```

比如上述代码中的Grow方法的接收者标识符*person*代表的是p的值的一个拷贝，而不是p的值。
而指针方法的副本则会指向p。这时的person.Age正是(*person).Age的速记法。

## 3. 接口

一个指针类型拥有以它以及以它的基底类型为接收者类型的所有方法，而它的基底类型却只拥有以它本身为接收者类型的方法。