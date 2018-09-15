package main

import "fmt"

func main() {
	for i := 1; i < 100; i++  {
		if i / 2 != 0 {
			fmt.Println(i)
		}
	}

	dayOfWeeks := [...]string{"Mon", "Tue", "Web", "Thu", "Fri", "Sat", "Sun"}
	for arrayIndex, dayOfWeek := range dayOfWeeks {
		fmt.Printf("%d番目の曜日は%s曜日です。\n", arrayIndex + 1, dayOfWeek)
	}
}
