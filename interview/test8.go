package main

// 下面的代码有什么问题?
import (
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age

}

// fatal error: concurrent map read and map write
func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age

	}
	return -1

}

//正确
func (ua *UserAges) Get(name string) int {
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age

	}
	return -1

}
