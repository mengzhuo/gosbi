// Code generated by "stringer -type=Errno -linecomment"; DO NOT EDIT.

package ecall

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SBI_SUCCESS-0]
	_ = x[SBI_ERR_FAILED - -1]
	_ = x[SBI_ERR_NOT_SUPPORTED - -2]
	_ = x[SBI_ERR_INVALID_PARAM - -3]
	_ = x[SBI_ERR_DENIED - -4]
	_ = x[SBI_ERR_INVALID_ADDRESS - -5]
	_ = x[SBI_ERR_ALREADY_AVAILABLE - -6]
	_ = x[SBI_ERR_ALREADY_STARTED - -7]
	_ = x[SBI_ERR_ALREADY_STOPPED - -8]
	_ = x[SBI_ERR_NO_SHMEM - -9]
}

const _Errno_name = "Shared memory not availableAlready stoppedAlready startedAlready availableInvalid address(s)Denied or not allowedInvalid parameter(s)Not supportedFailedCompleted successfully"

var _Errno_index = [...]uint8{0, 27, 42, 57, 74, 92, 113, 133, 146, 152, 174}

func (i Errno) String() string {
	i -= -9
	if i < 0 || i >= Errno(len(_Errno_index)-1) {
		return "Errno(" + strconv.FormatInt(int64(i+-9), 10) + ")"
	}
	return _Errno_name[_Errno_index[i]:_Errno_index[i+1]]
}
