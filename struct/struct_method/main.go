package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p *Person) Dream() {
	p.age = 200
	fmt.Printf("%s的梦想是有钱,今年%d岁\n", p.name, p.age)
}
func (p *Person) Run() {
	p.Dream()
}
func main() {
	p := NewPerson("徐亮", 30)
	fmt.Println(p)
	p.Dream()
	fmt.Println(p, p.name)
	p.Run()

}
