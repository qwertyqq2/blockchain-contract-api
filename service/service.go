package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math/big"
	"serv/config"
	"time"
)

type Service interface {
	Connect(ctx context.Context) error
	Deploy(ctx context.Context) (common.Address, error)
	MintTokens(ctx context.Context, count *big.Int) error
	SendTokens(ctx context.Context, to common.Address, count *big.Int) error
	GetBalance(ctx context.Context, address common.Address) (*big.Int, error)
}

type impl struct {
	providerUrl string
	pk          *ecdsa.PrivateKey

	contractAddress string
	contract        contractInstance
	cli             *ethclient.Client

	address common.Address
}

func NewService(conf config.Conf) (Service, error) {
	serv := &impl{}

	if conf.ProviderUrl == "" {
		return nil, fmt.Errorf("undefined provider")
	}

	privateKey, err := crypto.HexToECDSA(conf.PkKey)
	if err != nil {
		return nil, fmt.Errorf("cant parse pk")
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("err parse pub key")
	}

	serv.address = crypto.PubkeyToAddress(*publicKeyECDSA)
	serv.pk = privateKey
	serv.contractAddress = conf.ContractAddress
	serv.providerUrl = conf.ProviderUrl

	return serv, nil
}

func (s *impl) Connect(ctx context.Context) error {
	if s.providerUrl == "" {
		return fmt.Errorf("undef provider")
	}

	client, err := ethclient.DialContext(ctx, s.providerUrl)
	if err != nil {
		return fmt.Errorf("err conn: %w", err)
	}

	s.cli = client

	return s.setup(ctx)
}

func (s *impl) setup(ctx context.Context) error {
	if s.contractAddress != "" {
		logrus.Info("Setup contract...")
		contract := contractInstance{cli: s.cli}
		addr := common.HexToAddress(s.contractAddress)
		if err := contract.load(addr); err != nil {
			return fmt.Errorf("err load contract: %w", err)
		}
		s.contract = contract
	}

	logrus.Info("Service is ready")
	logrus.Info("Contract address: ", s.contractAddress)
	logrus.Info("Minter address: ", s.address.String())

	return nil
}

func (s *impl) Deploy(ctx context.Context) (common.Address, error) {
	contract := contractInstance{cli: s.cli}

	auth, err := s.currentAuth(ctx)
	if err != nil {
		logrus.Error(fmt.Errorf("err tx oprions: %w", err))
		return common.Address{}, fmt.Errorf("err tx oprions: %w", err)
	}

	logrus.Info("Deploy...")
	addr, err := contract.deploy(ctx, auth)
	if err != nil {
		logrus.Errorf("err deploy: %s", err.Error())
		return common.Address{}, fmt.Errorf("err deploy: %w", err)
	}

	s.contractAddress = addr.String()
	s.contract = contract

	logrus.Info("New contract address: ", addr.String())

	return s.contract.address, nil
}

func (s *impl) MintTokens(ctx context.Context, count *big.Int) error {
	auth, err := s.currentAuth(ctx)
	if err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Mint tokens...")

	if err := s.contract.mintTokens(auth, count); err != nil {
		logrus.Error(fmt.Errorf("err mint tokens: %w", err))
		return fmt.Errorf("err mint tokens: %w", err)
	}

	logrus.Info(fmt.Sprintf("Mint %s tokens to %s", count.String(), s.contractAddress))

	return nil

}

func (s *impl) SendTokens(ctx context.Context, to common.Address, count *big.Int) error {
	auth, err := s.currentAuth(ctx)
	if err != nil {
		logrus.Error(fmt.Errorf("err send tokens: %w", err))
		return fmt.Errorf("err send tokens: %w", err)
	}

	if err := s.contract.sendTokens(auth, to, count); err != nil {
		logrus.Error(fmt.Errorf("err send tokens: %w", err))
		return fmt.Errorf("err send tokens: %w", err)
	}

	logrus.Info(fmt.Sprintf("Send %s tokens to receiver with address %s", count.String(), to.String()))

	return nil
}

func (s *impl) GetBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	balance, err := s.contract.getBalance(address)
	if err != nil {
		logrus.Error(fmt.Errorf("err get balance: %w", err))
		return nil, fmt.Errorf("err get balance: %w", err)
	}

	return balance, nil
}

func (s *impl) currentAuth(ctx context.Context) (*transactionOptions, error) {
	nonce, err := s.cli.PendingNonceAt(ctx, s.address)
	if err != nil {
		return nil, err
	}

	chainID, err := s.cli.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	return &transactionOptions{
		pkKey:    s.pk,
		chainID:  chainID,
		nonce:    nonce,
		gasPrice: big.NewInt(1000000000),
		value:    big.NewInt(0),
		gasLimit: uint64(3000000),
	}, nil

}

func awaitTx(txHash common.Hash, cli *ethclient.Client, fn func(txHash common.Hash)) {
	tctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour*8))

	ticker := time.NewTicker(10 * time.Second)

	go func() {
		defer cancel()
		for {
			select {
			case <-tctx.Done():
				return
			case <-ticker.C:
				_, pending, err := cli.TransactionByHash(tctx, txHash)
				if err != nil {
					return
				}
				if !pending {
					fn(txHash)
					return
				}
			}
		}
	}()
}
