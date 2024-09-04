package base

import (
	"strconv"

	"github.com/mengzhuo/gosbi/ecall"
)

const ExtensionID = 0x10

const (
	get_spec_version = 0
	get_impl_id      = 1
	get_impl_version = 2
	probe_extension  = 3
	get_mvendorid    = 4
	get_marchid      = 5
	get_mimpid       = 6
)

// GetSpecVersion Returns the current SBI specification version
func GetSpecVersion() (major, minor int64, err error) {
	v, err := ecall.Ecall0(ExtensionID, get_spec_version)
	if err != nil {
		return
	}
	/*
		The minor number of the SBI specification is encoded
		in the low 24 bits, with the major number encoded in
		the next 7 bits
	*/
	major = int64(v>>24) & 0x7f
	minor = int64(v) & 0xffffff
	return
}

var ImplementationName = []string{
	"Berkeley Boot Loader (BBL)",
	"OpenSBI",
	"Xvisor",
	"KVM",
	"RustSBI",
	"Diosix",
	"Coffer",
	"Xen Project",
	"PolarFire Hart Software Services",
}

// GetImplID Returns the current SBI implementation ID, which is different
// for every SBI implementation. It is intended that this implementation
// ID allows software to probe for SBI implementation quirks
func GetImplID() (id int64, err error) {
	return ecall.Ecall0(ExtensionID, get_impl_id)
}

// GetImplName return name of GetImplID
func GetImplName() (name string, err error) {
	i, err := GetImplID()
	if err != nil {
		return "", err
	}
	if int(i) < len(ImplementationName) {
		return ImplementationName[i], nil
	}
	return "Unknown ImplID:" + strconv.Itoa(int(i)), nil
}

// GetImplVersion Returns the current SBI implementation version.
// The encoding of this version number is specific to the SBI implementation.
func GetImplVersion() (version int64, err error) {
	return ecall.Ecall0(ExtensionID, get_impl_version)
}

// ProbeExtension return whether given SBI extension ID (EID) is available or not.
func ProbeExtension(eid uintptr) (available bool, err error) {
	v, err := ecall.Ecall1(ExtensionID, probe_extension, eid)
	if err != nil {
		return
	}
	if v == 1 {
		return true, nil
	}
	return false, nil
}

// GetMVendorID Return a value that is legal for the mvendorid
// CSR and 0 is always a legal value for this CSR.
func GetMVendorID() (vid int64, err error) {
	return ecall.Ecall0(ExtensionID, get_mvendorid)
}

// GetMArchID Return a value that is legal for the marchid CSR
// and 0 is always a legal value for this CSR.
func GetMArchID() (aid int64, err error) {
	return ecall.Ecall0(ExtensionID, get_marchid)
}

// GetMImpID Return a value that is legal for the mimpid CSR
// and 0 is always a legal value for this CSR.
func GetMImpID() (iid int64, err error) {
	return ecall.Ecall0(ExtensionID, get_mimpid)
}
