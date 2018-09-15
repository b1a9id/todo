package main

import (
	"time"
	"fmt"
)

func main() {
	hour := time.Now().Hour()
	if hour >= 6 && hour < 12 {
		fmt.Println("morning")
	} else if hour < 19{
		fmt.Println("noon")
	} else {
		fmt.Println("night")
	}

	if hour2 := time.Now().Hour(); hour2 >= 6 && hour2 < 12 {
		fmt.Println("morning")
	} else if hour < 19{
		fmt.Println("noon")
	} else {
		fmt.Println("night")
	}

	dayOfWeek := "月"
	switch dayOfWeek {
	case "土", "日":
		fmt.Println("休み")
	default:
		fmt.Println("仕事")
	}
}
