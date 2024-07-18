package main

/*
#cgo LDFLAGS:./rust_extern/lib/libhello.a -ldl -lm
#include "./rust_extern/lib/hello.h"
#include <stdlib.h>
*/
import "C"
import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

func register_sgx_2(subclienturl string, warntime uint64, configversion uint16, deviceowner string) {
		// register 
		subclient_url := C.CString(subclienturl)
		defer C.free(unsafe.Pointer(subclient_url))
		owner := C.CString(deviceowner)
		defer C.free(unsafe.Pointer(owner))

		data := C.register_sgx_2(subclient_url,C.ulong(warntime),C.ushort(configversion),owner)

		fmt.Println("register to bool",data)	
}

var ErrSign = fmt.Errorf("sgx sign failed" )


func sign(msg []byte) (string,error) {
		// sign response
		tosign := hex.EncodeToString(msg)
		str5 := C.CString(tosign)
		defer C.free(unsafe.Pointer(str5))
		signed := C.sign_with_device_sgx_key(str5)
		str := C.GoStringN(signed, 128)
	
		_, err := hex.DecodeString(str)
		if err != nil {
			fmt.Println("Error decoding hex string:", err)
			return "",ErrSign
		}
		return str,nil
}


func register_sgx_test() {
	// register 
	subclient_url := C.CString("wss://Bool-Network-Gamma:9933")
	defer C.free(unsafe.Pointer(subclient_url))
	owner := C.CString("0xDeviceOwner-SubstratePk")
	defer C.free(unsafe.Pointer(owner))

	data := C.register_sgx_test(subclient_url,16,20,owner)
	fmt.Println("register_sgx_test",data)	
}

func sign_test(msg []byte) (string,error) {
	// sign response
	tosign := hex.EncodeToString(msg)
	str5 := C.CString(tosign)
	defer C.free(unsafe.Pointer(str5))
	signed := C.sign_with_device_sgx_key_test(str5)
	str := C.GoStringN(signed, 128)

	// _, err := hex.DecodeString(str)
	// if err != nil {
	// 	fmt.Println("Error decoding hex string:", err)
	// 	return "",ErrSign
	// }
	return str,nil
}

var ErrNopubkey = fmt.Errorf("mismatch pubkey type " )

func getSgxpublickey(keytype uint8) (string,error) {
	//get pubkey
	pubkey := C.getpublickey(C.ushort(keytype))
	pk_str := C.GoStringN(pubkey, 64)

	// decodedBytes_pk, err := hex.DecodeString(pk_str)
    // if err != nil {
    //     fmt.Println("Error decoding hex string:", err)
    //     return "",ErrNopubkey
    // }
	// fmt.Println("decodedBytes_pk is ",decodedBytes_pk)

	return string(pk_str), nil
}

func Verify_sgx_signature(msg []byte, sig string, pk string) uint16{

	msg_hex := hex.EncodeToString(msg)
	msg_cstring := C.CString(msg_hex)
	defer C.free(unsafe.Pointer(msg_cstring))

	verify_result := C.verify_sig_sgx(msg_cstring, C.CString(sig), C.CString(pk))
	
	return uint16(verify_result)
}