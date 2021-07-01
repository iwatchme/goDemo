package main

import (
	"container/list"
	"flag"
	"fmt"
	"math/rand"
)

//var name string

func init() {
	//flag.StringVar(&name, "name", "iwatchme", "please input the name")
}

var container = []string{"1111", "2222", "3333333"}

type Printer func(contents string) (n int, err error)

func main() {
	// 变量声明方式一
	//var name = flag.String("name", "iwatchme", "please input the name")

	//变量声明方式二， 只能在函数体内部使用
	name:= flag.String("name", "iwatchme", "please input the name")
    flag.Parse()
	fmt.Printf("Hello %s \n", *name)

	container := map[int]string{0: "zero", 1:"one", 2:"two"}

	// interface{}: 表示不包含任何方法的空接口类型
	value, ok := interface{}(container).([]string)

	fmt.Printf("the element is %s \n", container[0])

	fmt.Printf("is %s____%b \n", value, ok)

	// 数组Array 长度不可变 值类型
	a := [2]string{"1", "2"}
	fmt.Printf("len is %d\n",len(a))
	fmt.Printf("cap is %d\n",cap(a))

	//切片Slice 长度可变  引用类型
	//数组可以看成是切片底层的数据结构
	b := []string{"1", "2"}

	fmt.Printf("len is %d\n",len(b))
	fmt.Printf("cap is %d\n",cap(b))

	c := make([]int, 5, 8)
	fmt.Printf("len is %d\n",len(c))
	fmt.Printf("cap is %d\n",cap(c))

	c1 := c[3:5]
	fmt.Printf("len2 is %d\n",len(c1))
	fmt.Printf("cap2 is %d\n",cap(c1))


	l := list.New()
	e4 := l.PushBack(4)
	e3 := l.PushFront(3)
	l.InsertBefore(2, e4)
	l.InsertAfter(5, e3)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// channel
	//chan 是表示通道类型的关键字，而 int 则说明了该通道类型的元素类型
	//3: 通道的容量, 值为0时为非缓冲通道
	// <- 接送操作符。 一个左尖括号紧接着一个减号形象地代 表了元素值的传输方向。
	// 1. 对于同一个通道，发送操作之间是互斥的，接受操作之间也是互斥的
	ch1 := make(chan int, 3)

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	element1 := <- ch1

	fmt.Printf("the first elemetn received from %d \n", element1)


	//表示改通道是单向通道，只能收不能发
	//作用约束其他代码行为
	oneDirectionFunc(ch1)

	inChan2 := oneDirectionFunc2()
	for elem2 := range inChan2 {
		fmt.Printf("The element is %v \n\n", elem2)
	}


demo13()


}


// 12 函数章节
func demo12() {

	var p Printer
	p = printToStd
	p("something")

	/**
	  高阶函数：
	1. 接受其他函数作为参数传入
	2. 把其他函数作为结果返回
	*/

	// 闭包

}

// 13 及其使用方法
func demo13() {
	category := AnimalCategory{species: "cat"}
	animal := Animal{scientificName: "tiger"}
	animal.AnimalCategory = category

	fmt.Printf("the animal category: %s \n\n", category)

	fmt.Printf("the animal: %s \n\n", animal)
}

// AnimalCategory
type AnimalCategory struct {
	kingdom string // 界。
	phylum string // 门。
	class string // 纲。
	order string // 目。
	family string // 科。
	genus string // 属。
	species string // 种。
}

type Animal struct {
	scientificName string // 学名。
	AnimalCategory // 动物基本分类。
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

func oneDirectionFunc(ch chan <- int)  {
    ch <- rand.Intn(1000)
}

func oneDirectionFunc2() <- chan int{
    num := 5
    ch := make(chan int , num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
    return ch
}
