package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type Service interface {
	Connect(ctx context.Context) error
	MintTokens(ctx context.Context, count *big.Int) error
	SendTokens(ctx context.Context, to common.Address, count *big.Int) error
}

type Option func(s *impl)

func WithProvider(url string) Option {
	return func(s *impl) {
		s.providerUrl = url
	}
}

func WithPath(path string) Option {
	return func(s *impl) {
		s.repoPath = path
	}
}

func WithContractAddress(addr string) Option {
	return func(s *impl) {
		s.contractAddress = addr
	}
}

type impl struct {
	providerUrl string
	repoPath    string

	pk              *ecdsa.PrivateKey
	contractAddress string
	contract        contractInstance
	cli             *ethclient.Client

	address common.Address
}

func NewService(opts ...Option) Service {
	serv := &impl{}

	for _, o := range opts {
		o(serv)
	}

	return serv
}

func (s *impl) MintTokens(ctx context.Context, count *big.Int) error {
	auth, err := s.currentAuth(ctx)
	if err != nil {
		return err
	}

	return s.contract.mintTokens(auth, count, s.address)
}

func (s *impl) SendTokens(ctx context.Context, to common.Address, count *big.Int) error {
	auth, err := s.currentAuth(ctx)
	if err != nil {
		return err
	}

	return s.contract.sendTokens(auth, to, count)
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
	if s.repoPath == "" {
		log.Println("Generate private key...")
		pk, err := crypto.GenerateKey()
		if err != nil {
			return err
		}
		s.pk = pk
	} else {
		pk, err := loadPkKey(s.repoPath)
		if err != nil {
			return err
		}
		s.pk = pk
	}

	publicKey := s.pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("err parse pub key")
	}

	addr := crypto.PubkeyToAddress(*publicKeyECDSA)
	s.address = addr

	contract := &contractInstance{}

	if s.contractAddress == "" {
		auth, err := s.currentAuth(ctx)
		if err != nil {
			return err
		}

		log.Println("Deploy contract...")
		addr, err := contract.deploy(ctx, s.cli, auth)
		if err != nil {
			return err
		}

		s.contractAddress = addr.String()
	} else {
		if err := contract.load(s.contractAddress, s.cli); err != nil {
			return err
		}

	}
	return nil
}

func (s *impl) currentAuth(ctx context.Context) (*transactionOptions, error) {
	publicKey := s.pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("err parse pub key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.cli.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := s.cli.SuggestGasPrice(ctx)
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
		gasPrice: gasPrice,
		value:    big.NewInt(0),
		gasLimit: 3000000,
	}, nil

}

func loadPkKey(repoPath string) (*ecdsa.PrivateKey, error) {
	return nil, nil
}
