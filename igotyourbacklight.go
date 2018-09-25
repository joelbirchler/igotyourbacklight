package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	brightness := Brightness{value: readCurrent(), max: readMax()}

	if len(os.Args) == 2 {
		arg := os.Args[1]

		delta, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			panic(err)
		}

		brightness.Increment(int(delta))
		write(brightness)
	} else {
		fmt.Println(brightness)
	}
}
