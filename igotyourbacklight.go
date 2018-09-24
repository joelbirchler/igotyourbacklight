package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

func readMax() int {
	return read("max_brightness")
}

func readCurrent() int {
	return read("brightness")
}

func read(filename string) int {
	data, err := ioutil.ReadFile("/sys/class/backlight/intel_backlight/" + filename)
	if err != nil {
		panic(err)
	}

	content := strings.TrimSuffix(string(data), "\n")

	val, err := strconv.ParseInt(content, 10, 64)
	if err != nil {
		panic(err)
	}

	return int(val)
}

func write(b Brightness) {
	s := fmt.Sprintf("%v\n", b.Value())
	err := ioutil.WriteFile("output.txt", []byte(s), 0644)
	if err != nil {
		panic(err)
	}
}
