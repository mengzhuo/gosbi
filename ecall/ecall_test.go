package ecall

import "testing"

func TestGetSBISpecVersion(t *testing.T) {
	val, err := Ecall0(0x10, 0)
	if err != nil {
		t.Error(err)
	}
	t.Logf("spec version:%d", val)
}
