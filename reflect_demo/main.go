package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Student
}
type Student struct {
	Num   string `json:"num"`
	Score int    `json:"score"`
}
type myInt int64

func (u User) Hello() {
	fmt.Println("test")
}
func testTypeOf(a interface{}) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	// fmt.Printf("%v type:%v kine:%v numfield:%v\n", t, t.Name(), t.Kind(), t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("1:%v 2:%v 3:%v\n", t.Field(i), v.Field(i), t.Field(i).Type.Name())
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))

	}
	fmt.Println("####")
	// if scoreField, ok := t.FieldByName("Id"); ok {
	// 	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	// }
	//反射TypeOf(类型) 和ValueOf(值)
	//
	sValue := v.FieldByName("Student") //拿到嵌套结构体的值信息
	sType := sValue.Type()             //嵌套结构体的类型信息
	for i := 0; i < sValue.NumField(); i++ {
		field := sType.Field(i) //对应结构体字段的值

		fieldvalue := sValue.Field(i)       //对应结构体字段的值
		fmt.Println(field.Name, fieldvalue) //字段名、值
	}
}

//反射 结构体
func reflect_type(a interface{}) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	fmt.Println("TypeOf:", t)
	fmt.Println("TypeOf Name:", t.Name())

	fmt.Println("ValueOf:", v)
	k := t.Kind()
	fmt.Println(k)
	kk := v.Kind()
	fmt.Println(kk)
	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Printf("a is string\n")
	}

	switch kk {
	case reflect.Float64:
		fmt.Println("a is: ", v.Float())
	}

}

func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("TypeOf:", t)
	fmt.Println("TypeOf Name:", t.Name())
	fmt.Println("TypeOf NumField:", t.NumField())
	v := reflect.ValueOf(o)
	fmt.Println(v)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()

		fmt.Printf("%s=%v :%v\n", f.Name, val, f.Type)
	}
	fmt.Println("######")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
func testValueOf(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Printf("%#v %v\n", v, v.Elem().Kind())

	k := v.Elem().Kind()
	switch k {
	case reflect.Struct:
		fmt.Printf("%v\n", k)
	case reflect.Ptr:
		fmt.Printf("%v\n", v.Type().Name())
	case reflect.Int64:
		v.Elem().SetInt(200)
	case reflect.Float32:
		v.Elem().SetFloat(100)
	}

}
func main() {
	// var x float64 = 3.14
	// reflect_type(x)
	// u := User{1, "xiaoming", 22}
	// Poni(u)
	// fmt.Println(u)
	// var aa int64 = 100
	// var a *int64 = &aa
	// fmt.Println(*a)
	// testValueOf(a)
	// fmt.Println(*a)
	// var f = new(float32)
	// *f = 3.14
	// fmt.Println(*f)
	// testValueOf(f)
	// fmt.Println(*f)
	// var b myInt = 100
	// testValueOf(b)
	// var c rune
	// testValueOf(c)
	var u = User{1, "xxx", 22, Student{"111", 100}}
	testTypeOf(u)

	// fmt.Println(reflect.Struct)
}
