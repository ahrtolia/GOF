package creator

import "sync"

// Author      : zhaoxiaodong
// Create      : 2022/7/19 16:34
// Description : 单例模式

// 使用懒惰模式的单例模式，使用双重检查加锁保证线程安全

//Singleton 是单例模式类
type Singleton struct{}

var singleton *Singleton
var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}
