package main

// 公開していないパッケージなので、適当な名前をつけている
import (
	"fmt"
	"kyoruni/go-package-test/calculate" // calculate
	"kyoruni/go-package-test/action"   // act
)

func main() {
	num := calculate.Double(3)
	result := act.Attack(num)
	fmt.Println(result) // 6のダメージを与えた！
}
