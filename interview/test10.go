package main

// 以下代码能编译过去吗？为什么？
import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"

	} else {
		talk = "hi"

	}
	return

}

// 是否能通过编译
func main() {
	var peo People = Stduent{}
	//var peo People = &Stduent{}   // 正确，可以通过编译
	think := "bitch"
	fmt.Println(peo.Speak(think))

}

/*
答案是不能，提示Stduent does not implement People (Speak method has pointer receiver)，修正方法是var peo People = &Stduent{}，赋值为指针

在Go语言中，函数和方法是不一样的，函数是没有接收者的，而方法是有接收者的，属于某个结构体。

接受者有两种：value receivers(按值传递)，pointer receivers(按指针传递，可能改变传入的参数)

之前看到有地方解释道：pointer receivers的方法既能传值也能传指针，包含了value receivers的方法。

这个理解是错误的。

从这道题来看不是这样的，value receivers和pointer receivers是有明确区分的，people interface是value receivers，而student实现了pointer receivers的speak方法，所以student没有实现peple。

事实上pointer receivers的方法只是将m.speak()自动转换为(&m).speak()，这点官方的go tour中就有提到，go编译器帮你省去了取地址的一步。
*/
