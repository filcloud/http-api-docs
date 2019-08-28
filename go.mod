module github.com/ipfs/http-api-docs

require (
	github.com/AndreasBriese/bbloom v0.0.0-20190825152654-46b345b51c96 // indirect
	github.com/Stebalien/go-json-doc v0.0.2
	github.com/filecoin-project/go-filecoin v0.0.1
	github.com/ipfs/go-cid v0.0.3
	github.com/ipfs/go-ipfs-cmdkit v0.0.1
	github.com/ipfs/go-ipfs-cmds v0.0.8
	github.com/ipfs/go-ipfs-files v0.0.3 // indirect
	github.com/ipfs/go-ipld-cbor v0.0.2 // indirect
	github.com/ipfs/go-ipld-format v0.0.2 // indirect
	github.com/ipfs/go-path v0.0.4 // indirect
	github.com/ipfs/go-unixfs v0.0.6 // indirect
	github.com/ipfs/iptb v1.4.0 // indirect
	github.com/libp2p/go-libp2p-peer v0.2.0
	github.com/libp2p/go-libp2p-peerstore v0.1.2
	github.com/multiformats/go-multiaddr v0.0.4
	go4.org v0.0.0-20190313082347-94abd6928b1d // indirect
)

replace github.com/filecoin-project/go-filecoin => ../../filecoin-project/go-filecoin

replace github.com/filecoin-project/go-bls-sigs => ../../filecoin-project/go-filecoin/go-bls-sigs

replace github.com/filecoin-project/go-sectorbuilder => ../../filecoin-project/go-filecoin/go-sectorbuilder

replace github.com/ipfs/go-ipfs-cmdkit => github.com/ipfs/go-ipfs-cmdkit v0.0.1

replace github.com/ipfs/go-ipfs-cmds => github.com/ipfs/go-ipfs-cmds v0.0.1
