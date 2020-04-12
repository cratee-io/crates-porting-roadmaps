# crates-rsgx

This repository implements a tool building dependency graph about SGX-adapted crates based on [sgx-world](https://github.com/dingelish/sgx-world).

## Goal

Port all crates listed in the Dependency Graph section based on different versions of 
[teaclave-sgx-sdk](https://github.com/apache/incubator-teaclave-sgx-sdk).

## Porting Workflow

1. Fork the project
2. Upgrade to rust 2018
3. Format the project
4. Commit the changes all above 
5. Replace the `std` library with `sgx_tstd`
6. Add tests under the `tests/sgx` folder
7. Make sure tests run well
8. Make another commit to save the changes introduced by porting
9. Push the commit to a branch of pattern `rsgx{major.minor.patch}`, where `major`, `minor` and
`patch` corresponds to the tags as those of [teaclave-sgx-sdk](https://github.com/apache/incubator-teaclave-sgx-sdk). For example, if porting to teaclave-sgx-sdk@v1.1.1, then we should
push to branch `rsgx1.1.1`.

## Dependency Graph

The crate of smaller line number don't depend on those of larger line number.

> `versions` specify the teaclave-sgx-sdk version the porting bases on.

no. | crate | versions
----|-------|---------
000 | [adler32](https://github.com/mesalock-linux/adler32-rs-sgx/blob/master/Cargo.toml) | [v1.1.1](https://github.com/sammyne/adler32-rs/tree/rsgx1.1.1)
001 | [integer-encoding](https://github.com/mesalock-linux/integer-encoding-rs-sgx/blob/master/Cargo.toml) | [v1.1.1](https://github.com/sammyne/integer-encoding-rs/tree/rsgx1.1.1)
002 | [matches](https://github.com/mesalock-linux/rust-std-candidates-sgx/blob/master/matches/Cargo.toml) | [v1.1.1](https://github.com/sammyne/rust-std-candidates/tree/rsgx1.1.1)
003 | [inventory](https://github.com/mesalock-linux/inventory-sgx/blob/master/Cargo.toml) | support `no_std`
004 | [parity-scale-codec-derive](https://github.com/mesalock-linux/parity-scale-codec-sgx/blob/master/derive/Cargo.toml) | macro crate
005 | [hex-literal](https://github.com/mesalock-linux/rustcrypto-utils-sgx/blob/master/hex-literal/Cargo.toml) | support `no_std`
006 | [thiserror](https://github.com/mesalock-linux/thiserror-sgx/blob/master/Cargo.toml) | [v1.1.1](https://github.com/sammyne/thiserror/tree/rsgx1.1.1)
007 | [thread-id](https://github.com/mesalock-linux/thread-id-sgx/blob/master/Cargo.toml) | [v1.1.1](https://github.com/sammyne/thread-id/tree/rsgx1.1.1)
008 | [threadpool](https://github.com/mesalock-linux/rust-threadpool-sgx/blob/master/Cargo.toml) | 
009 | [itoa](https://github.com/mesalock-linux/itoa-sgx/blob/master/Cargo.toml) | 
010 | [ppv-lite86](https://github.com/mesalock-linux/cryptocorrosion-sgx/blob/master/utils-simd/ppv-lite86/Cargo.toml) | 
011 | [color_quant](https://github.com/mesalock-linux/color_quant-sgx/blob/master/Cargo.toml) | 
012 | [md5](https://github.com/mesalock-linux/md5-sgx/blob/master/Cargo.toml) | 
013 | [threshold-secret-sharing](https://github.com/mesalock-linux/rust-threshold-secret-sharing-sgx/blob/master/Cargo.toml) | 
014 | [matrixmultiply](https://github.com/mesalock-linux/matrixmultiply-sgx/blob/master/Cargo.toml) | 
015 | [unicase](https://github.com/mesalock-linux/unicase-sgx/blob/master/Cargo.toml) | 
016 | [iovec](https://github.com/mesalock-linux/iovec-sgx/blob/master/Cargo.toml) | 
017 | [array_tool](https://github.com/mesalock-linux/array_tool-sgx/blob/master/Cargo.toml) | 
018 | [memchr](https://github.com/mesalock-linux/rust-memchr-sgx/blob/master/Cargo.toml) | 
019 | [etcommon-hexutil](https://github.com/mesalock-linux/etcommon-rs-sgx/blob/master/hexutil/Cargo.toml) | 
020 | [hex](https://github.com/mesalock-linux/rust-hex-sgx/blob/master/Cargo.toml) | 
021 | [ring](https://github.com/mesalock-linux/ring-sgx/blob/v0.16.5/Cargo.toml) | 
022 | [termcolor](https://github.com/mesalock-linux/termcolor-sgx/blob/master/Cargo.toml) | 
023 | [base64](https://github.com/mesalock-linux/rust-base64-sgx/blob/master/Cargo.toml) | 
024 | [block-padding](https://github.com/mesalock-linux/rustcrypto-utils-sgx/blob/master/block-padding/Cargo.toml) | 
025 | [take_mut](https://github.com/mesalock-linux/take_mut-sgx/blob/master/Cargo.toml) | 
026 | [slab](https://github.com/mesalock-linux/slab-sgx/blob/master/Cargo.toml) | 
027 | [byteorder](https://github.com/mesalock-linux/byteorder-sgx/blob/master/Cargo.toml) | 
028 | [ucd-util](https://github.com/mesalock-linux/ucd-generate-sgx/blob/master/ucd-util/Cargo.toml) | 
029 | [net2](https://github.com/mesalock-linux/net2-rs-sgx/blob/master/Cargo.toml) | 
030 | [rdrand](https://github.com/mesalock-linux/rust_rdrand-sgx/blob/master/Cargo.toml) | 
031 | [subtle](https://github.com/mesalock-linux/subtle-sgx/blob/master/Cargo.toml) | 
032 | [lzw](https://github.com/mesalock-linux/lzw-sgx/blob/master/Cargo.toml) | 
033 | [rustc-serialize](https://github.com/mesalock-linux/rustc-serialize-sgx/blob/master/Cargo.toml) | 
034 | [byte-slice-cast](https://github.com/mesalock-linux/byte-slice-cast-sgx/blob/master/Cargo.toml) | 
035 | [aead](https://github.com/mesalock-linux/rustcrypto-traits-sgx/blob/master/aead/Cargo.toml) | 
036 | [assert_matches](https://github.com/mesalock-linux/assert_matches-sgx/blob/master/Cargo.toml) | 
037 | [utf8-ranges](https://github.com/mesalock-linux/utf8-ranges-sgx/blob/master/Cargo.toml) | 
038 | [percent-encoding](https://github.com/mesalock-linux/rust-url-sgx/blob/master/percent_encoding/Cargo.toml) | 
039 | [rle-decode-fast](https://github.com/mesalock-linux/rle-decode-helper-sgx/blob/master/Cargo.toml) | 
040 | [anyhow](https://github.com/mesalock-linux/anyhow-sgx/blob/master/Cargo.toml) | 
041 | [fnv](https://github.com/mesalock-linux/rust-fnv-sgx/blob/master/Cargo.toml) | 
042 | [num-traits](https://github.com/mesalock-linux/num-traits-sgx/blob/master/Cargo.toml) | 
043 | [serde_derive](https://github.com/mesalock-linux/serde-sgx/blob/master/serde_derive/Cargo.toml) | 
044 | [crc32fast](https://github.com/mesalock-linux/rust-crc32fast-sgx/blob/master/Cargo.toml) | 
045 | [failure](https://github.com/mesalock-linux/failure-sgx/blob/master/Cargo.toml) | 
046 | [quick-error](https://github.com/mesalock-linux/quick-error-sgx/blob/master/Cargo.toml) | 
047 | [heapsize](https://github.com/mesalock-linux/heapsize-sgx/blob/master/Cargo.toml) | 
048 | [itertools](https://github.com/mesalock-linux/rust-itertools-sgx/blob/master/Cargo.toml) | 
049 | [thread_local](https://github.com/mesalock-linux/thread_local-rs-sgx/blob/master/Cargo.toml) | 
050 | [miniz_oxide](https://github.com/mesalock-linux/miniz_oxide-sgx/blob/master/miniz_oxide/Cargo.toml) | 
051 | [inflate](https://github.com/mesalock-linux/inflate-sgx/blob/master/Cargo.toml) | 
052 | [data-url](https://github.com/mesalock-linux/rust-url-sgx/blob/master/data-url/Cargo.toml) | 
053 | [aho-corasick](https://github.com/mesalock-linux/aho-corasick-sgx/blob/master/Cargo.toml) | 
054 | [sct](https://github.com/mesalock-linux/sct.rs/blob/mesalock_sgx/Cargo.toml) | 
055 | [webpki](https://github.com/mesalock-linux/webpki/blob/mesalock_sgx/Cargo.toml) | 
056 | [parity-wasm](https://github.com/mesalock-linux/parity-wasm-sgx/blob/master/Cargo.toml) | 
057 | [blobby](https://github.com/mesalock-linux/rustcrypto-utils-sgx/blob/master/blobby/Cargo.toml) | 
058 | [block-buffer](https://github.com/mesalock-linux/rustcrypto-utils-sgx/blob/master/block-buffer/Cargo.toml) | 
059 | [jpeg-decoder](https://github.com/mesalock-linux/jpeg-decoder-sgx/blob/master/Cargo.toml) | 
060 | [snap](https://github.com/mesalock-linux/rust-snappy-sgx/blob/master/Cargo.toml) | 
061 | [protected_fs_rs](https://github.com/mesalock-linux/protected_fs_rs/blob/master/Cargo.toml) | 
062 | [universal-hash](https://github.com/mesalock-linux/rustcrypto-traits-sgx/blob/master/universal-hash/Cargo.toml) | 
063 | [gif](https://github.com/mesalock-linux/image-gif-sgx/blob/master/Cargo.toml) | 
064 | [rust-crypto](https://github.com/mesalock-linux/rust-crypto-sgx/blob/master/Cargo.toml) | 
065 | [num-integer](https://github.com/mesalock-linux/num-integer-sgx/blob/master/Cargo.toml) | 
066 | [serde](https://github.com/mesalock-linux/serde-sgx/blob/master/serde/Cargo.toml) | 
067 | [libflate](https://github.com/mesalock-linux/libflate-sgx/blob/master/Cargo.toml) | 
068 | [gzip-header](https://github.com/mesalock-linux/gzip-header-sgx/blob/master/Cargo.toml) | 
069 | [humantime](https://github.com/mesalock-linux/humantime-sgx/blob/master/Cargo.toml) | 
070 | [elastic-array](https://github.com/mesalock-linux/elastic-array-sgx/blob/master/Cargo.toml) | 
071 | [tiff](https://github.com/mesalock-linux/image-tiff-sgx/blob/master/Cargo.toml) | 
072 | [regex](https://github.com/mesalock-linux/regex-sgx/blob/master/Cargo.toml) | 
073 | [webpki-roots](https://github.com/mesalock-linux/webpki-roots/blob/mesalock_sgx/Cargo.toml) | 
074 | [digest](https://github.com/mesalock-linux/rustcrypto-traits-sgx/blob/master/digest/Cargo.toml) | 
075 | [crypto-mac](https://github.com/mesalock-linux/rustcrypto-traits-sgx/blob/master/crypto-mac/Cargo.toml) | 
076 | [stream-cipher](https://github.com/mesalock-linux/rustcrypto-traits-sgx/blob/master/stream-cipher/Cargo.toml) | 
077 | [block-cipher-trait](https://github.com/mesalock-linux/rustcrypto-traits-sgx/blob/master/block-cipher-trait/Cargo.toml) | 
078 | [num-iter](https://github.com/mesalock-linux/num-iter-sgx/blob/master/Cargo.toml) | 
079 | [serde_bytes](https://github.com/mesalock-linux/serde-bytes-sgx/blob/master/Cargo.toml) | 
080 | [indexmap](https://github.com/mesalock-linux/indexmap-sgx/blob/master/Cargo.toml) | 
081 | [num-rational](https://github.com/mesalock-linux/num-rational-sgx/blob/master/Cargo.toml) | 
082 | [chrono](https://github.com/mesalock-linux/chrono-sgx/blob/master/Cargo.toml) | 
083 | [bitvec](https://github.com/mesalock-linux/bitvec-sgx/blob/master/Cargo.toml) | 
084 | [sha1](https://github.com/mesalock-linux/rust-sha1-sgx/blob/master/Cargo.toml) | 
085 | [bytes](https://github.com/mesalock-linux/bytes-sgx/blob/master/Cargo.toml) | 
086 | [unicode-bidi](https://github.com/mesalock-linux/unicode-bidi-sgx/blob/master/Cargo.toml) | 
087 | [encoding_rs](https://github.com/mesalock-linux/encoding_rs-sgx/blob/master/Cargo.toml) | 
088 | [log](https://github.com/mesalock-linux/log-sgx/blob/master/Cargo.toml) | 
089 | [bincode](https://github.com/mesalock-linux/bincode-sgx/blob/master/Cargo.toml) | 
090 | [wabt](https://github.com/mesalock-linux/wabt-rs-sgx/blob/v0.9-core/Cargo.toml) | 
091 | [serde-big-array](https://github.com/mesalock-linux/serde-big-array-sgx/blob/master/Cargo.toml) | 
092 | [linked-hash-map](https://github.com/mesalock-linux/linked-hash-map-sgx/blob/master/Cargo.toml) | 
093 | [serde_cbor](https://github.com/mesalock-linux/cbor-sgx/blob/master/Cargo.toml) | 
094 | [erased-serde](https://github.com/mesalock-linux/erased-serde-sgx/blob/master/Cargo.toml) | 
095 | [smallvec](https://github.com/mesalock-linux/rust-smallvec-sgx/blob/master/Cargo.toml) | 
096 | [deflate](https://github.com/mesalock-linux/deflate-rs-sgx/blob/dev/Cargo.toml) | 
097 | [etcommon-rlp](https://github.com/mesalock-linux/etcommon-rs-sgx/blob/master/rlp/Cargo.toml) | 
098 | [whirlpool](https://github.com/mesalock-linux/rustcrypto-hashes-sgx/blob/master/whirlpool/Cargo.toml) | 
099 | [blake-hash](https://github.com/mesalock-linux/cryptocorrosion-sgx/blob/master/hashes/blake/Cargo.toml) | 
100 | [groestl-aesni](https://github.com/mesalock-linux/cryptocorrosion-sgx/blob/master/hashes/groestl/Cargo.toml) | 
101 | [sha2](https://github.com/mesalock-linux/rustcrypto-hashes-sgx/blob/master/sha2/Cargo.toml) | 
102 | [md2](https://github.com/mesalock-linux/rustcrypto-hashes-sgx/blob/master/md2/Cargo.toml) | 
103 | [md4](https://github.com/mesalock-linux/rustcrypto-hashes-sgx/blob/master/md4/Cargo.toml) | 
104 | [jh-x86_64](https://github.com/mesalock-linux/cryptocorrosion-sgx/blob/master/hashes/jh/Cargo.toml) | 
105 | [ripemd320](https://github.com/mesalock-linux/rustcrypto-hashes-sgx/blob/master/ripemd320/Cargo.toml) | 
106 | [gost94](https://github.com/mesalock-linux/rustcrypto-hashes-sgx/blob/master/gost94/Cargo.toml) | 
107 | [ripemd160](https://github.com/mesalock-linux/rustcrypto-hashes-sgx/blob/master/ripemd160/Cargo.toml) | 
108 | [groestl](https://github.com/mesalock-linux/rustcrypto-hashes-sgx/blob/master/groestl/Cargo.toml) | 
109 | [sha-1](https://github.com/mesalock-linux/rustcrypto-hashes-sgx/blob/master/sha1/Cargo.toml) | 
110 | [md-5](https://github.com/mesalock-linux/rustcrypto-hashes-sgx/blob/master/md5/Cargo.toml) | 
111 | [streebog](https://github.com/mesalock-linux/rustcrypto-hashes-sgx/blob/master/streebog/Cargo.toml) | 
112 | [sha3](https://github.com/mesalock-linux/rustcrypto-hashes-sgx/blob/master/sha3/Cargo.toml) | 
113 | [hmac](https://github.com/mesalock-linux/rustcrypto-MACs-sgx/blob/master/hmac/Cargo.toml) | 
114 | [blake2](https://github.com/mesalock-linux/rustcrypto-hashes-sgx/blob/master/blake2/Cargo.toml) | 
115 | [salsa20-core](https://github.com/mesalock-linux/rustcrypto-stream-ciphers-sgx/blob/master/salsa20-core/Cargo.toml) | 
116 | [c2-chacha](https://github.com/mesalock-linux/cryptocorrosion-sgx/blob/master/stream-ciphers/chacha/Cargo.toml) | 
117 | [twofish](https://github.com/mesalock-linux/rustcrypto-block-ciphers-sgx/blob/master/twofish/Cargo.toml) | 
118 | [kuznyechik](https://github.com/mesalock-linux/rustcrypto-block-ciphers-sgx/blob/master/kuznyechik/Cargo.toml) | 
119 | [ctr](https://github.com/mesalock-linux/rustcrypto-stream-ciphers-sgx/blob/master/ctr/Cargo.toml) | 
120 | [ofb](https://github.com/mesalock-linux/rustcrypto-stream-ciphers-sgx/blob/master/ofb/Cargo.toml) | 
121 | [des](https://github.com/mesalock-linux/rustcrypto-block-ciphers-sgx/blob/master/des/Cargo.toml) | 
122 | [rc2](https://github.com/mesalock-linux/rustcrypto-block-ciphers-sgx/blob/master/rc2/Cargo.toml) | 
123 | [cmac](https://github.com/mesalock-linux/rustcrypto-MACs-sgx/blob/master/cmac/Cargo.toml) | 
124 | [aesni](https://github.com/mesalock-linux/rustcrypto-block-ciphers-sgx/blob/master/aes/aesni/Cargo.toml) | 
125 | [cfb-mode](https://github.com/mesalock-linux/rustcrypto-stream-ciphers-sgx/blob/master/cfb-mode/Cargo.toml) | 
126 | [aes-soft](https://github.com/mesalock-linux/rustcrypto-block-ciphers-sgx/blob/master/aes/aes-soft/Cargo.toml) | 
127 | [cast5](https://github.com/mesalock-linux/rustcrypto-block-ciphers-sgx/blob/master/cast5/Cargo.toml) | 
128 | [magma](https://github.com/mesalock-linux/rustcrypto-block-ciphers-sgx/blob/master/magma/Cargo.toml) | 
129 | [blowfish](https://github.com/mesalock-linux/rustcrypto-block-ciphers-sgx/blob/master/blowfish/Cargo.toml) | 
130 | [threefish-cipher](https://github.com/mesalock-linux/cryptocorrosion-sgx/blob/master/block-ciphers/threefish/Cargo.toml) | 
131 | [cfb8](https://github.com/mesalock-linux/rustcrypto-stream-ciphers-sgx/blob/master/cfb8/Cargo.toml) | 
132 | [block-modes](https://github.com/mesalock-linux/rustcrypto-block-ciphers-sgx/blob/master/block-modes/Cargo.toml) | 
133 | [serde_json](https://github.com/mesalock-linux/serde-json-sgx/blob/master/Cargo.toml) | 
134 | [yasna](https://github.com/mesalock-linux/yasna.rs-sgx/blob/master/Cargo.toml) | 
135 | [parity-scale-codec](https://github.com/mesalock-linux/parity-scale-codec-sgx/blob/master/Cargo.toml) | 
136 | [prost](https://github.com/mesalock-linux/prost-sgx/blob/master/Cargo.toml) | 
137 | [http](https://github.com/mesalock-linux/http-sgx/blob/master/Cargo.toml) | 
138 | [protobuf](https://github.com/mesalock-linux/rust-protobuf-sgx/blob/v2.8/protobuf/Cargo.toml) | 
139 | [kvdb](https://github.com/mesalock-linux/parity-common-sgx/blob/master/kvdb/Cargo.toml) | 
140 | [env_logger](https://github.com/mesalock-linux/env_logger-sgx/blob/master/Cargo.toml) | 
141 | [rustls](https://github.com/mesalock-linux/rustls/blob/mesalock_sgx/Cargo.toml) | 
142 | [rand](https://github.com/mesalock-linux/rand-sgx/blob/v0.6.5_sgx1.1.1/Cargo.toml) | 
143 | [mio](https://github.com/mesalock-linux/mio-sgx/blob/v0.6_sgx_1.1.1/Cargo.toml) | 
144 | [getrandom](https://github.com/mesalock-linux/getrandom-sgx/blob/master/Cargo.toml) | 
145 | [toml](https://github.com/mesalock-linux/toml-rs-sgx/blob/master/Cargo.toml) | 
146 | [typetag](https://github.com/mesalock-linux/typetag-sgx/blob/master/Cargo.toml) | 
147 | [unicode-normalization](https://github.com/mesalock-linux/unicode-normalization-sgx/blob/master/Cargo.toml) | 
148 | [png](https://github.com/mesalock-linux/image-png-sgx/blob/master/Cargo.toml) | 
149 | [etcommon-bigint](https://github.com/mesalock-linux/etcommon-rs-sgx/blob/master/bigint/Cargo.toml) | 
150 | [hmac-drbg](https://github.com/mesalock-linux/hmac-drbg-rs-sgx/blob/master/Cargo.toml) | 
151 | [chacha20](https://github.com/mesalock-linux/rustcrypto-stream-ciphers-sgx/blob/master/chacha20/Cargo.toml) | 
152 | [salsa20](https://github.com/mesalock-linux/rustcrypto-stream-ciphers-sgx/blob/master/salsa20/Cargo.toml) | 
153 | [daa](https://github.com/mesalock-linux/rustcrypto-MACs-sgx/blob/master/daa/Cargo.toml) | 
154 | [aes-ctr](https://github.com/mesalock-linux/rustcrypto-stream-ciphers-sgx/blob/master/aes-ctr/Cargo.toml) | 
155 | [aes](https://github.com/mesalock-linux/rustcrypto-block-ciphers-sgx/blob/master/aes/aes/Cargo.toml) | 
156 | [skein-hash](https://github.com/mesalock-linux/cryptocorrosion-sgx/blob/master/hashes/skein/Cargo.toml) | 
157 | [jsonwebtoken](https://github.com/mesalock-linux/jsonwebtoken-sgx/blob/master/Cargo.toml) | 
158 | [kvdb-memorydb](https://github.com/mesalock-linux/parity-common-sgx/blob/master/kvdb-memorydb/Cargo.toml) | 
159 | [jsonpath_lib](https://github.com/mesalock-linux/jsonpath-sgx/blob/master/Cargo.toml) | 
160 | [http_req](https://github.com/mesalock-linux/http_req-sgx/blob/master/Cargo.toml) | 
161 | [uuid](https://github.com/mesalock-linux/uuid-sgx/blob/master/Cargo.toml) | 
162 | [quickcheck](https://github.com/mesalock-linux/quickcheck-sgx/blob/master/Cargo.toml) | 
163 | [num-complex](https://github.com/mesalock-linux/num-complex-sgx/blob/master/Cargo.toml) | 
164 | [rusty-leveldb](https://github.com/mesalock-linux/rusty_leveldb_sgx/blob/sgx/Cargo.toml) | 
165 | [num-bigint-dig](https://github.com/mesalock-linux/num-bigint-sgx/blob/dig/Cargo.toml) | 
166 | [rsa](https://github.com/mesalock-linux/rustcrypto-RSA-sgx/blob/master/Cargo.toml) | 
167 | [wasmi](https://github.com/mesalock-linux/wasmi-sgx/blob/master/Cargo.toml) | 
168 | [idna](https://github.com/mesalock-linux/rust-url-sgx/blob/master/idna/Cargo.toml) | 
169 | [image](https://github.com/mesalock-linux/image-sgx/blob/master/Cargo.toml) | 
170 | [etcommon-bloom](https://github.com/mesalock-linux/etcommon-rs-sgx/blob/master/bloom/Cargo.toml) | 
171 | [etcommon-trie](https://github.com/mesalock-linux/etcommon-rs-sgx/blob/master/trie/Cargo.toml) | 
172 | [etcommon-block-core](https://github.com/mesalock-linux/etcommon-rs-sgx/blob/master/block-core/Cargo.toml) | 
173 | [libsecp256k1](https://github.com/mesalock-linux/libsecp256k1-rs-sgx/blob/master/Cargo.toml) | 
174 | [ndarray](https://github.com/mesalock-linux/ndarray-sgx/blob/master/Cargo.toml) | 
175 | [num](https://github.com/mesalock-linux/num-sgx/blob/master/Cargo.toml) | 
176 | [url](https://github.com/mesalock-linux/rust-url-sgx/blob/master/Cargo.toml) | 
177 | [etcommon-block](https://github.com/mesalock-linux/etcommon-rs-sgx/blob/master/block/Cargo.toml) | 
178 | [rulinalg](https://github.com/mesalock-linux/rulinalg-sgx/blob/master/Cargo.toml) | 
179 | [rust-base58](https://github.com/mesalock-linux/rust-base58-sgx/blob/master/Cargo.toml) | 
180 | [sputnikvm](https://github.com/mesalock-linux/sputnikvm-sgx/blob/master/Cargo.toml) | 
181 | [rusty-machine](https://github.com/mesalock-linux/rusty-machine-sgx/blob/master/Cargo.toml) | 
-------
