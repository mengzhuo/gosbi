package ecall

func ecall0(eid uintptr, fid uintptr) (errno int64, value int64)

//go:nosplit
func Ecall0(eid uintptr, fid uintptr) (value int64, err Errno) {
	errno, value := ecall0(eid, fid)
	err = Errno(errno)
	return value, err
}

func ecall1(eid uintptr, fid uintptr, arg0 uintptr) (errno int, value int64)

//go:nosplit
func Ecall1(eid uintptr, fid uintptr, arg0 uintptr) (value int64, err Errno) {
	errno, value := ecall1(eid, fid, arg0)
	err = Errno(errno)
	return value, err
}
