package main

import "fmt"

//方法
//标识符：变量名、函数名、类型名、方法名
//go语言中如果标识符首字母大写就表示对外部可见、暴露的、共有的
type dog struct {
	name string
}

type person struct {
	name string
	age  int
}

//给自定义类型加方法
//不能给别的包里的类型添加方法，只能给自己的包里类型添加方法
type myInt int

func (m myInt) hello() {
	fmt.Println("这是一个int")
}

//构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

//方法是作用于特定类型的函数
//接收者表示的是调用该方法的具体类型的变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%v,汪汪汪～\n", d.name)
}

//使用值接收者
func (p person) guonian() {
	p.age++
	//fmt.Printf("%v\n", p.age)

}

//使用指针接收者
func (p *person) guonian1() {
	p.age++
	// fmt.Printf("%v\n", p.age)

}
func main() {
	d := newDog("花花")
	fmt.Printf("%v\n", d)
	d.wang()
	p := newPerson("华晨宇", 30)
	fmt.Printf("%v\n", p)
	p.guonian()
	fmt.Printf("%v\n", p)
	p.guonian1()
	fmt.Printf("%v\n", p)
	m := myInt(100)
	m.hello()

}
