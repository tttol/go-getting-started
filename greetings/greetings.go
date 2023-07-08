package greetings

import "fmt"

func Hello(name string) string {
	// 変数宣言と代入を同時に行う`:=`
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
