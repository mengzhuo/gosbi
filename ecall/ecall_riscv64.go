package ecall

func ecall0(eid uintptr, fid uintptr) (errno int64, value int64)

func Ecall0(eid uintptr, fid uintptr) (value int64, err error) {
	errno, value := ecall0(eid, fid)
	if errno != 0 {
		err = Errno(errno)
	}
	return value, err
}

func ecall1(eid uintptr, fid uintptr, arg0 uintptr) (errno int, value int64)

func Ecall1(eid uintptr, fid uintptr, arg0 uintptr) (value int64, err error) {
	errno, value := ecall1(eid, fid, arg0)
	if errno != 0 {
		err = Errno(errno)
	}
	return value, err
}
