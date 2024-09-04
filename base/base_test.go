package base_test

import (
	"fmt"

	. "github.com/mengzhuo/gosbi/base"
)

func ExampleGetSpecVersion() {
	major, minor, err := GetSpecVersion()
	if err != nil {
		panic(err)
	}
	fmt.Println(major, minor)
	// Ouput: 0, 2
}
