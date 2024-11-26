module github.com/ellemouton/slog

go 1.22.6

require (
	github.com/btcsuite/btclog v0.0.0-20241017175713-3428138b75c7
	github.com/btcsuite/btclog/v2 v2.0.0-20241017175713-3428138b75c7
	github.com/lightningnetwork/lnd v0.18.0-beta.rc4.0.20241125112551-66940024acb3
)

replace (
	github.com/btcsuite/btclog/v2 => github.com/ellemouton/btclog/v2 v2.0.0-20241126072058-e0b4906c5b62
	github.com/lightningnetwork/lnd => github.com/ellemouton/lnd v0.8.0-beta-rc3.0.20241125133716-f07e4dade684
)

require (
	github.com/aead/siphash v1.0.1 // indirect
	github.com/btcsuite/btcd v0.24.3-0.20240921052913-67b8efd3ba53 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.4 // indirect
	github.com/btcsuite/btcd/btcutil v1.1.5 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0 // indirect
	github.com/btcsuite/go-socks v0.0.0-20170105172521-4720035b7bfd // indirect
	github.com/btcsuite/websocket v0.0.0-20150119174127-31079b680792 // indirect
	github.com/btcsuite/winsvc v1.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/decred/dcrd/lru v1.1.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/jessevdk/go-flags v1.4.0 // indirect
	github.com/jrick/logrotate v1.1.2 // indirect
	github.com/kkdai/bstream v1.0.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
