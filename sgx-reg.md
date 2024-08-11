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
	cd rust_extern/lib/hello && cargo build --release  &&\
	cd ../../../
	cp rust_extern/lib/hello/target/release/libhello.a rust_extern/lib/
	rm btcd 
    CC=gcc occlum-go build
```

sgx :
```
	cd rust_extern/lib/hello && occlum-cargo build --release  &&\
	cd ../../../  &&\
    cp rust_extern/lib/hello/target/x86_64-unknown-linux-musl/release/libhello.a rust_extern/lib/ &&\
	rm btcd  &&\
    occlum-go build
```

build-occlcum
```
rm -rf occlum_instance && mkdir occlum_instance && cd occlum_instance && \
occlum init && rm -rf image && \
new_json="$(jq '.resource_limits.kernel_space_stack_size = "30MB" |
                .resource_limits.kernel_space_heap_size = "256MB" |
                .resource_limits.kernel_space_heap_max_size = "4096MB" |
                .resource_limits.user_space_size = "1024MB" |
                .resource_limits.user_space_max_size = "4096MB" |
                .resource_limits.init_num_of_threads = 8 |
                .resource_limits.max_num_of_threads = 128 |
                .process.default_heap_size = "320MB" |
                .process.default_stack_size = "30MB" |
                .process.default_mmap_size = "2048MB" |
                .env.untrusted = ["EXAMPLE", "RUST_LOG"] |
                .metadata.debuggable = false |
                .feature.enable_edmm = false' Occlum.json)" && \
echo "${new_json}" > Occlum.json  && \
copy_bom -f ../rust-demo.yaml --root image --include-dir /opt/occlum/etc/template && \
occlum build
```
run:
```
occlum run /bin/btcd --rpcuser=prz --rpcpass=prz --subclienturl=ws://127.0.0.1:9933 --deviceowner=0xf24FF3a9CF04c71Dbc94D0b566f7A27B94566cac --watcherdeviceid=0x75686b813035e46622fe2baf6c9fd7ed826f16d8d9f02117a8035c4f76787a99 --verifysig --sgxenable --notls
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

