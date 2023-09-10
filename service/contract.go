package service

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"serv/coin"
)

type contractInstance struct {
	coin    *coin.Coin
	address common.Address
}

func (c *contractInstance) deploy(ctx context.Context, cli *ethclient.Client, o *transactionOptions) (*common.Address, error) {
	auth, err := c.auth(o)
	if err != nil {
		return nil, err
	}

	address, _, instance, err := coin.DeployCoin(auth, cli)
	if err != nil {
		return nil, err
	}

	c.address = address
	c.coin = instance

	return &address, nil
}

func (c *contractInstance) load(addr string, cli *ethclient.Client) error {
	address := common.HexToAddress(addr)
	instance, err := coin.NewCoin(address, cli)
	if err != nil {
		return err
	}

	c.address = address
	c.coin = instance

	return nil
}

func (c *contractInstance) mintTokens(o *transactionOptions, count *big.Int, minter common.Address) error {
	auth, err := c.auth(o)
	if err != nil {
		return err
	}

	if _, err := c.coin.Mint(auth, minter, count); err != nil {
		return err
	}

	return nil
}

func (c *contractInstance) sendTokens(o *transactionOptions, to common.Address, count *big.Int) error {
	auth, err := c.auth(o)
	if err != nil {
		return err
	}

	if _, err := c.coin.Send(auth, to, count); err != nil {
		return err
	}

	return nil
}

type transactionOptions struct {
	pkKey    *ecdsa.PrivateKey
	chainID  *big.Int
	nonce    uint64
	value    *big.Int
	gasLimit uint64
	gasPrice *big.Int
}

func (c *contractInstance) auth(o *transactionOptions) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(o.pkKey, o.chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = new(big.Int).SetUint64(o.nonce)
	auth.Value = o.value       // in wei
	auth.GasLimit = o.gasLimit // in units
	auth.GasPrice = o.gasPrice

	return auth, nil

}
