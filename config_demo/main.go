package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//ini配置文件解析器
//MysqlConfig mysql配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadFile(fileName string, data interface{}) (err error) {
	//参数校验
	//传进来的data参数必须为指针类型（因为需要在函数中赋值
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data show be a pointer") //新建一个错误
		return
	}
	//传进来的data参数必须是结构体类型指针（因为配置文件中各种键值对需要赋值给结构体的字段
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data show be a struct") //新建一个错误
		return
	}
	//读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("read ini file failed, err:%v\n", err)
	}
	lineSlice := strings.Split(string(b), "\n")
	fmt.Printf("type:%T %s\n", lineSlice, lineSlice)
	//一行一行读数据
	var structName string
	for idx, line := range lineSlice {
		//去掉首尾空格
		line = strings.TrimSpace(line)
		//如果是空行就跳过
		if len(line) == 0 {
			continue
		}
		//如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//如果是[开头就表示是节(section)
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line %d syntax error", idx+1)
				return
			}
			//把这行首尾的[]去掉，取到中间的内容把首尾的空格去掉
			sectionName := strings.TrimSpace(line[1 : len(line)-1]) //[mysql]
			fmt.Printf("len:%d sectionName:%s\n", len(line), sectionName)
			if len(sectionName) == 0 {
				err = fmt.Errorf("line %d syntax error", idx+1)
				return
			}
			//根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					//说明找到了对应的嵌套结构体，把字段名记下来
					structName = field.Name //找到对应的结构体
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			//如果不是[开头就是=分割的键值对
			//以等号分割这行，等号左边key，右边value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line %d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			//定义配置文件中的key和value
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//根据structName去data里面把对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) //拿到嵌套结构体的值信息
			sType := sValue.Type()                     //拿到嵌套结构体的类型信息
			// structObj := v.Elem().FieldByName(structName)
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField

			//遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i) //tag信息存储在类型信息中
				fileType = field
				if field.Tag.Get("ini") == key {
					//找到对应的字段
					fieldName = field.Name
					break
				}

			}
			//如果key=tag，给这个字段赋值
			//根据fieldName去取出这个字段
			if len(fieldName) == 0 {
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			fmt.Println(fieldName, fileObj, fileType.Type.Kind())
			//对其赋值
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)

					return
				}
				fileObj.SetInt(valueInt)

			}

		}
	}
	return
}
func main() {
	var cfg Config
	err := loadFile("./config.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini filed, err:%v\n", err)
		return
	}
	fmt.Println(cfg.MysqlConfig)
}
