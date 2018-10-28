package fortest

import "fmt"

// 仅用于测试
// 这里应该是显示详细内容吧
func Sum(numbers []int) int {
	sum := 0
	for n := range numbers {
		fmt.Println(n)
		sum += n
	}
	return sum
}

// may be
// it only suport english
func main() {
	fmt.Println("vim-go")
	nums := []int{1, 2, 9, 4}
	fmt.Println("sum result: ", Sum(nums))
}
