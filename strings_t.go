package main

import (
	"log"
	"unicode/utf8"
)

func main() {
	lst := "测试一下很长的祝福语和表情包😜😂😳😜🐔🐔"
	lst = "Hello, 世🖖🏿  🖖界"
	log.Println(lst, " 长度: ", len(lst))
	log.Println(utf8.RuneCountInString(lst))
}
