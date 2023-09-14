package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math/big"
	"serv/coin"
)

type contractInstance struct {
	coin    *coin.Coin
	address common.Address
	cli     *ethclient.Client
}

func (c *contractInstance) deploy(ctx context.Context, o *transactionOptions) (*common.Address, error) {
	auth, err := c.auth(o)
	if err != nil {
		return nil, err
	}

	address, tx, instance, err := coin.DeployCoin(auth, c.cli)
	if err != nil {
		return nil, err
	}

	logrus.Info("Tx hash: ", tx.Hash().String())

	c.address = address
	c.coin = instance

	return &address, nil
}

func (c *contractInstance) load(addr common.Address, cli *ethclient.Client) error {
	instance, err := coin.NewCoin(addr, cli)
	if err != nil {
		return err
	}

	c.address = addr
	c.coin = instance
	c.cli = cli

	return nil
}

func (c *contractInstance) mintTokens(o *transactionOptions, count *big.Int) error {
	auth, err := c.auth(o)
	if err != nil {
		return err
	}

	tx, err := c.coin.Mint(auth, count)
	if err != nil {
		return err
	}

	awaitTx(tx.Hash(), c.cli, func(txHash common.Hash) {
		fmt.Printf("tx success %s", txHash.String())
	})

	logrus.Info("Tx hash: ", tx.Hash().String())

	return nil
}

func (c *contractInstance) sendTokens(o *transactionOptions, to common.Address, count *big.Int) error {
	auth, err := c.auth(o)
	if err != nil {
		return err
	}

	tx, err := c.coin.Send(auth, to, count)
	if err != nil {
		return err
	}

	awaitTx(tx.Hash(), c.cli, func(txHash common.Hash) {
		fmt.Printf("tx success %s", txHash.String())
	})

	logrus.Info("Tx hash: ", tx.Hash().String())

	return nil
}

func (c *contractInstance) getBalance(whom common.Address) (*big.Int, error) {
	val, err := c.coin.GetBalance(&bind.CallOpts{
		Pending: true,
	}, whom)
	if err != nil {
		return nil, fmt.Errorf("err get balance: %w", err)
	}

	return val, nil
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
