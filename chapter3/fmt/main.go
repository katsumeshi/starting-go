package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Golang!")
	fmt.Println("My", "name", "is", "Taro")
	fmt.Printf("数値=%d\n", 5)
	fmt.Printf("10数値＝%d 2数値＝%b 8数値＝%o 16数値＝%x\n", 17, 17, 17, 17)
	fmt.Printf("%d年%d月%d日\n", 2015, 12)
	fmt.Printf("%d年%d月%d日\n", 2015, 12, 25, 17)
	fmt.Printf("数値=%v　文字列=%v　配列=%v\n", 5, "Golang", [...]int{1, 2, 3})
	fmt.Printf("数値=%#v　文字列=%#v　配列=%#v\n", 5, "Golang", [...]int{1, 2, 3})
	fmt.Printf("数値=%T　文字列=%T　配列=%T\n", 5, "Golang", [...]int{1, 2, 3})
	fmt.Print("Hello, World!")
	fmt.Println("Hello, World!")
	fmt.Println(1, 2, 3)
}
