package ether

import "math/big"

var APIKeyToken = "3QSCE6H3DJ8RVENY5BCMD21A3GHIEXB9ZJ"

func WeiToETH(wei *big.Float) *big.Float {

	return wei.Quo(wei, big.NewFloat(1e18))
}
