package ether

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetEtherBalance(address string) (*big.Float, error) {

	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		return nil, err
		// log.Fatal(err)
	}
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
		// log.Fatal(err)
	}
	ether := WeiToETH(new(big.Float).SetInt(balance))
	return ether, nil
}
