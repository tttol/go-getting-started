package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// 変数宣言と代入を同時に行う`:=`
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
