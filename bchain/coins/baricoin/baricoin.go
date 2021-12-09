package baricoin

import (
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

const (
	// MainnetMagic is mainnet network constant
	MainnetMagic wire.BitcoinNet = 0x66658273
	// TestnetMagic is testnet network constant
	TestnetMagic wire.BitcoinNet = 0x98617266
	// RegtestMagic is regtest network constant
	RegtestMagic wire.BitcoinNet = 0x98617266
)

var (
	// MainNetParams are parser parameters for mainnet
	MainNetParams chaincfg.Params
	// TestNetParams are parser parameters for testnet
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{26}
	MainNetParams.ScriptHashAddrID = []byte{21}
	MainNetParams.Bech32HRPSegwit = "bari"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{66}
	TestNetParams.ScriptHashAddrID = []byte{18}
	TestNetParams.Bech32HRPSegwit = "baritn"
}

// BaricoinParser handle
type BaricoinParser struct {
	*btc.BitcoinParser
}

// NewBaricoinParser returns new BaricoinParser instance
func NewBaricoinParser(params *chaincfg.Params, c *btc.Configuration) *BaricoinParser {
	return &BaricoinParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main Baricoin network,
// and the test Baricoin network
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}
