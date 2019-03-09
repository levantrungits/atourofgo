package main

import (
	"fmt"
)

var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
	return 42
}

func init() {
	WhatIsThe = 0
}

func main2() {
	if WhatIsThe == 0 {
		fmt.Println("It's all a lie.")
	}
}