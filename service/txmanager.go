package service

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

func awaitTx(txHash common.Hash, cli *ethclient.Client, fn func(txHash common.Hash)) {
	tctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour*8))

	ticker := time.NewTicker(2 * time.Second)

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
