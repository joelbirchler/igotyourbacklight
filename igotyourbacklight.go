package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	current, max := readCurrent(), readMax()

	if len(os.Args) == 2 {
		arg := os.Args[1]

		delta, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			panic(err)
		}

		adjusted := adjust(delta, current, max)
		fmt.Println(adjusted)
	} else {
		percent := int64(100 * current / max)
		fmt.Printf("%v%% (%v/%v)\n", percent, current, max)
	}

}

// FIXME: what if we have a type (constrainedInt?) for the backlight int that keeps it's range in check and can give a percent?
func adjust(deltaPercent int64, current int64, max int64) int64 {
	percent := int64(100*current/max) + deltaPercent
	adjusted := (max * percent) / 100

	switch {
	case adjusted < 0:
		return 0
	case adjusted > max:
		return max
	}

	return adjusted
}

func readMax() int64 {
	return read("max_brightness")
}

func readCurrent() int64 {
	return read("brightness")
}

func read(filename string) int64 {
	data, err := ioutil.ReadFile("/sys/class/backlight/intel_backlight/" + filename)
	if err != nil {
		panic(err)
	}

	content := strings.TrimSuffix(string(data), "\n")

	val, err := strconv.ParseInt(content, 10, 64)
	if err != nil {
		panic(err)
	}

	return val
}

//
// TODO read: /sys/class/backlight/intel_backlight/max_brightness
// TODO write: /sys/class/backlight/intel_backlight/brightness (remember \n)
// TODO: figure out how to test
//
