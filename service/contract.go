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
	coin         *coin.Coin
	address      common.Address
	resourseCoin *big.Int
	balances     map[common.Address]*big.Int

	cli *ethclient.Client
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

	awaitTx(tx.Hash(), c.cli, func(txHash common.Hash) {
		c.address = address
		c.coin = instance
		c.setupCoin()
		logrus.Info("tx success: ", txHash.String())
	})
	return &address, nil
}

func (c *contractInstance) load(addr common.Address) error {
	instance, err := coin.NewCoin(addr, c.cli)
	if err != nil {
		return err
	}

	c.address = addr
	c.coin = instance
	c.setupCoin()

	return nil
}

func (c *contractInstance) setupCoin() {
	c.balances = make(map[common.Address]*big.Int)
	c.resourseCoin = big.NewInt(0)
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
		if c.resourseCoin == nil {
			c.resourseCoin = new(big.Int).Add(big.NewInt(0), count)
		}
		c.resourseCoin = new(big.Int).Add(c.resourseCoin, count)
		logrus.Info("tx success: ", txHash.String())
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
		bal, ok := c.balances[to]
		if !ok {
			c.balances[to] = new(big.Int).Add(big.NewInt(0), count)
			return
		}

		c.balances[to] = new(big.Int).Add(bal, count)
		logrus.Info("tx success: ", txHash.String())
	})

	logrus.Info("Tx hash: ", tx.Hash().String())

	return nil
}

func (c *contractInstance) getBalance(whom common.Address) (*big.Int, error) {
	if bal, ok := c.balances[whom]; ok {
		return bal, nil
	}

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
