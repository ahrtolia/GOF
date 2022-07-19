package creator

import "fmt"

// Author      : zhaoxiaodong
// Create      : 2022/7/19 15:49
// Description : 简单工厂

//go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。
//NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。
//
//在这个simplefactory包中只有API 接口和NewAPI函数为包外可见，封装了实现细节。

type API interface {
	// Do 工厂通用方法
	Do(name string) string
}

// NewAPI 创建工厂对象
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

// hiAPI 工厂对象
type hiAPI struct{}

//Do hi to name
func (*hiAPI) Do(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

//HelloAPI is another ApiInterface implement
type helloAPI struct{}

//Do hello to name
func (*helloAPI) Do(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
