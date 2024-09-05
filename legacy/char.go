package legacy

import "github.com/mengzhuo/gosbi/ecall"

type Errno uintptr

func ConsolePutchar(b byte) error {
	_, err := ecall.Ecall1(0x1, 0, uintptr(b))
	return err
}
