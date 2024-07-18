registration to bool by sgx remote attestation
====

## Build

git submodule fetch `https://github.com/ABMatrix/rust_extern`:
```
git submudule init && git submudule update
```
or
```
git clone --recursive 
```

```
	cd rust_extern/lib/hello && cargo build --release
	cp rust_extern/lib/hello/target/release/libhello.a rust_extern/lib/
    go build
```

## Usage

### test

run node:
```
 ./btcd --regtest --rpcuser=prz --rpcpass=prz --miningaddr=bcrt1p5cj85luz7uhaugxpusgtk3xpyp5wmje6ks5fcl3njzegfdxexuws45r4l6 --verifysig
```

generate some blocks:
```
curl -k  --user prz --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "generate", "params": [10]}'  https://127.0.0.1:18334/
```

sgx publickey:
```
curl -k  --user prz --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getsgxpubkey", "params": []}'  https://127.0.0.1:18334/
```

### sgx

start node:
```
./btcd --regtest --rpcuser=prz --rpcpass=prz --miningaddr=bcrt1p5cj85luz7uhaugxpusgtk3xpyp5wmje6ks5fcl3njzegfdxexuws45r4l6 --verifysig --subclienturl=ws://192.168.0.1:9944 --warntime=20 --configversion=16 --deviceowner=0x1234 --sgxenable
```
call node:

Same as above

