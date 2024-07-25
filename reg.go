package main

/*
#cgo LDFLAGS:./rust_extern/lib/libhello.a -ldl -lm
#include "./rust_extern/lib/hello.h"
#include <stdlib.h>
*/
import "C"
import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"unsafe"
)

func register_sgx_2(subclienturl string, warntime uint64, configversion uint16, deviceowner string) {
	// register
	subclient_url := C.CString(subclienturl)
	defer C.free(unsafe.Pointer(subclient_url))
	owner := C.CString(deviceowner)
	defer C.free(unsafe.Pointer(owner))

	data := C.register_sgx_2(subclient_url, C.ulong(warntime), C.ushort(configversion), owner)

	fmt.Println("register to bool", data)
}

var ErrSign = fmt.Errorf("sgx sign failed")

func sign(msg []byte) (string, error) {
	// sign response
	tosign := hex.EncodeToString(msg)
	str5 := C.CString(tosign)
	defer C.free(unsafe.Pointer(str5))
	signed := C.sign_with_device_sgx_key(str5)
	str := C.GoStringN(signed, 128)

	_, err := hex.DecodeString(str)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return "", ErrSign
	}
	return str, nil
}

func register_sgx_test() {
	// register
	subclient_url := C.CString("wss://Bool-Network-Gamma:9933")
	defer C.free(unsafe.Pointer(subclient_url))
	owner := C.CString("0xDeviceOwner-SubstratePk")
	defer C.free(unsafe.Pointer(owner))

	data := C.register_sgx_test(subclient_url, 16, 20, owner)
	fmt.Println("register_sgx_test", data)
}

func sign_test(msg []byte) (string, error) {
	// sign response
	tosign := hex.EncodeToString(msg)
	str5 := C.CString(tosign)
	defer C.free(unsafe.Pointer(str5))
	signed := C.sign_with_device_sgx_key_test(str5)
	str := C.GoStringN(signed, 128)
	return str, nil
}

var ErrNopubkey = fmt.Errorf("mismatch pubkey type ")

func getSgxpublickey(keytype uint8) (string, error) {

	pubkey := C.getpublickey(C.ushort(keytype))
	pk_str := C.GoStringN(pubkey, 64)

	return string(pk_str), nil
}

func Verify_sgx_signature(msg []byte, sig string, pk string) uint16 {

	msg_hex := hex.EncodeToString(msg)
	msg_cstring := C.CString(msg_hex)
	defer C.free(unsafe.Pointer(msg_cstring))

	verify_result := C.verify_sig_sgx(msg_cstring, C.CString(sig), C.CString(pk))

	return uint16(verify_result)
}

type ResponseSGX struct {
	Result json.RawMessage `json:"resp"`
	Sig    string          `json:"sig"`
	PK     string          `json:"pubkey"`
}

type tmpResultSGX struct {
	Result interface{} `json:"result"`
	Sig    string      `json:"sig"`
}

func convert_to_sgx_result(result []byte) ([]byte, error) {

	jsonBytes, err := JSONRemarshal(result)
	if err != nil {
		rpcsLog.Error("JSONRemarshal", err)
		return nil, err
	}

	var signature string
	var pk string
	var err1 error
	var err2 error
	if SGXmode {
		signature, err1 = sign(jsonBytes)
		pk, err2 = getSgxpublickey(0)

	} else {
		signature, err1 = sign_test(jsonBytes)
		pk, err2 = getSgxpublickey(1)

	}

	if err1 != nil {
		rpcsLog.Error("sgx sign", err1)
		return nil, err1
	}

	if err2 != nil {
		rpcsLog.Error("sgx pk", err2)
		return nil, err2
	}

	if VerifySignature {
		verifyres := Verify_sgx_signature(jsonBytes, signature, pk)
		rpcsLog.Infof("verify signature result %d", verifyres)
	}

	res := ResponseSGX{
		Result: result,
		Sig:    signature,
		PK:     pk,
	}

	return json.Marshal(&res)
}

func JSONRemarshal(bytes []byte) ([]byte, error) {
	var ifce interface{}
	err := json.Unmarshal(bytes, &ifce)
	if err != nil {
		return nil, err
	}
	return json.Marshal(ifce)
}
