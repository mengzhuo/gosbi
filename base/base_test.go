package base_test

import (
	"fmt"
	"testing"

	. "github.com/mengzhuo/gosbi/base"
	"github.com/mengzhuo/gosbi/ecall"
)

func ExampleGetSpecVersion() {
	major, minor, err := GetSpecVersion()
	if err != ecall.SBI_SUCCESS {
		panic(err)
	}
	fmt.Println(major, minor)
	// Ouput: 0, 2
}

func TestGetSpecVersion(t *testing.T) {
	major, minor, err := GetSpecVersion()
	if err != ecall.SBI_SUCCESS {
		t.Skipf("err=%v", err)
	}
	t.Log(major, minor)

}
