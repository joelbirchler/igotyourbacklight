package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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
	err := ioutil.WriteFile("/sys/class/backlight/intel_backlight/brightness", []byte(s), 0644)
	if err != nil {
		panic(err)
	}
}
