package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestParsePk(t *testing.T) {
	privateKey, err := crypto.HexToECDSA("d796d40644dbf2b3dfaecfec166afb5f9c498317bec97b6e7a996df7845e1f85")
	if err != nil {
		t.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("nok")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println(fromAddress.String())
}

func TestUmnarshal(t *testing.T) {
	privateKey, err := crypto.HexToECDSA("74781f38a17f12cbf6855d4c53cc85b411f52e8ed1161326dedc3d49f75024da")
	if err != nil {
		t.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("nok")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println(fromAddress.String())

}

func TestGenAddress(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("err parse")
	}

	addr := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println(addr.String())

}

func TestGetCode(t *testing.T) {
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, "https://polygon-mumbai.blockpi.network/v1/rpc/public")
	if err != nil {
		t.Fatal(err)
	}

	code, err := client.CodeAt(ctx, common.HexToAddress("0xdc2AEc3E46cbCe19e2A15D621127746488eC8d58"), nil)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(code)

}

func TestConnectGanache(t *testing.T) {
	providerURL := "http://0.0.0.0:8545"

	client, err := ethclient.DialContext(context.Background(), providerURL)
	if err != nil {
		t.Fatal(err)
	}

	cid, err := client.ChainID(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(cid)
}
