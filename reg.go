package main

/*
#cgo LDFLAGS:./lib/libhello.a -ldl -lm
#include "./lib/hello.h"
#include <stdlib.h>
*/
import "C"
import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

func register_sgx_2() {
		// register 
		subclient_url := C.CString("wsss://dfljdflajlfdj:0000")
		defer C.free(unsafe.Pointer(subclient_url))
		owner := C.CString("0xDeviceOwner010101")
		defer C.free(unsafe.Pointer(owner))
		data := C.register_sgx_2(subclient_url,16,20,owner)
		fmt.Println(data)	
}

func sign(msg []byte) {
		// sign response
		tosign := hex.EncodeToString(msg)
		str5 := C.CString(tosign)
		defer C.free(unsafe.Pointer(str5))
		signed := C.sign_with_device_sgx_key(str5)
		str := C.GoStringN(signed, 128)
	
		decodedBytes, err := hex.DecodeString(str)
		if err != nil {
			fmt.Println("Error decoding hex string:", err)
			return
		}
		fmt.Println("decoded is ",decodedBytes)
}


func register_sgx_test() {
	// register 
	subclient_url := C.CString("wsss://dfljdflajlfdj:0000")
	defer C.free(unsafe.Pointer(subclient_url))
	owner := C.CString("0xDeviceOwner010101")
	defer C.free(unsafe.Pointer(owner))
	data := C.register_sgx_test(subclient_url,16,20,owner)
	fmt.Println(data)	
}

func sign_test(msg []byte) {
	// sign response
	tosign := hex.EncodeToString(msg)
	str5 := C.CString(tosign)
	defer C.free(unsafe.Pointer(str5))
	signed := C.sign_with_device_sgx_key_test(str5)
	str := C.GoStringN(signed, 128)

	decodedBytes, err := hex.DecodeString(str)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return
	}
	fmt.Println("decoded is ",decodedBytes)
}

var ErrNopubkey = fmt.Errorf("mismatch pubkey type " )

func getSgxpublickey(keytype uint8) (string,error) {
	//get pubkey
	pubkey := C.getpublickey( C.ushort(keytype))
	pk_str := C.GoStringN(pubkey, 128)

	decodedBytes_pk, err := hex.DecodeString(pk_str)
    if err != nil {
        fmt.Println("Error decoding hex string:", err)
        return "",ErrNopubkey
    }
	fmt.Println("decodedBytes_pk is ",decodedBytes_pk)

	return string(decodedBytes_pk), nil
}