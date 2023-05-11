package main

import (
	"fmt"
	"time"
)

//timeDemo 本地时间
func timeDemo() {
	now := time.Now()
	fmt.Printf("current time: %v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("year:%v,month:%d,day:%d,hour:%v,minute:%v,second:%v\n", year, month, day, hour, minute, second)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
func timestampDemo() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time: %v\n", now)
	t := now.Unix() //时间戳
	n := now.UnixNano()
	fmt.Printf("t:%v, n:%v\n", t, n)

}
func formatDemo() {
	now := time.Now()
	fmt.Printf("current time: %v\n", now)
	a := now.Format("2006-01-02 15:04:05.000")
	a = now.Format("2006-01-02 03:04")
	fmt.Println(a)

}

func main() {
	// formatDemo()
	// timestampDemo()
	now := time.Now() //本地的时间
	fmt.Println(now)
	//按照东八区的时区和格式去解析一个字符串格式的时间
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed, err:%v\n", err)
		return
	}
	// str := string(now.Format("2006-01-02 15:04:22"))
	// //按照指定格式去解析一个字符串格式的时间
	// time.Parse("2006-01-02 15:04:05", str)
	str := string(now.Format("2006-01-02 15:04:22"))
	fmt.Printf("%T, %v\n", str, str)
	//按照指定时区解析时间，以下解析出东八区时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	//时间对象相减
	fmt.Println(timeObj.Sub(now))

}
