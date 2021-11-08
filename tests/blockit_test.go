package tests

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Test_Local(t *testing.T) {
	addresses := make([]common.Address, 0)
	addresses = append(addresses, common.HexToAddress("0x1509027c6d99d1c63fd1940c592f93d09e3c9e45"))
	addresses = append(addresses, common.HexToAddress("0x0483058142ece6b426350242710a389ed96d076d"))
	addresses = append(addresses, common.HexToAddress("0x0b5b7aac9c554692b54b95e96f776f7a50611b80"))
	addresses = append(addresses, common.HexToAddress("0x0021c445131c74af616bb7a982d03b8423f18ca3"))
	client, _ := ethclient.Dial("http://192.168.56.210:8545")
	balances, _ := client.MultipleBalanceAt(context.Background(), addresses, big.NewInt(21))
	for address, balance := range balances {
		t.Log(fmt.Sprintf("%s:%d", address.String(), balance))
	}
}
