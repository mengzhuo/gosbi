#include "textflag.h"

TEXT ·ecall0(SB),NOSPLIT,$0
	MOV	eid+0(FP), A7
	MOV	fid+8(FP), A6
	ECALL
	MOV	A0, errno+16(FP)
	MOV	A1, value+24(FP)
	RET

TEXT ·ecall1(SB),NOSPLIT,$0
	MOV	eid+0(FP), A7
	MOV	fid+8(FP), A6
	MOV	arg0+16(FP), A0
	ECALL
	MOV	A0, errno+24(FP)
	MOV	A1, value+32(FP)
	RET
