package ecall

type Errno int64

//go:generate stringer -type=Errno -linecomment
const (
	SBI_SUCCESS               Errno = 0  // Completed successfully
	SBI_ERR_FAILED            Errno = -1 // Failed
	SBI_ERR_NOT_SUPPORTED     Errno = -2 // Not supported
	SBI_ERR_INVALID_PARAM     Errno = -3 // Invalid parameter(s)
	SBI_ERR_DENIED            Errno = -4 // Denied or not allowed
	SBI_ERR_INVALID_ADDRESS   Errno = -5 // Invalid address(s)
	SBI_ERR_ALREADY_AVAILABLE Errno = -6 // Already available
	SBI_ERR_ALREADY_STARTED   Errno = -7 // Already started
	SBI_ERR_ALREADY_STOPPED   Errno = -8 // Already stopped
	SBI_ERR_NO_SHMEM          Errno = -9 // Shared memory not available
)
