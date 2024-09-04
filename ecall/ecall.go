//go:build !riscv64

package ecall

func Ecall0(eid uintptr, fid uintptr) (value int64, err error) {
	panic("Not supported GOARCH")
}

func Ecall1(eid uintptr, fid uintptr, arg0 uintptr) (value int64, err error) {
	panic("Not supported GOARCH")
}
