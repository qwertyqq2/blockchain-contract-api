package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/big"
	"serv/service"
)

func main() {
	serv := service.NewService(
		service.WithProvider("https://goerli.infura.io/v3/779b28f2780741df94f2ebac83bd49d1"),
	)

	if err := serv.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Post("/contract/mint", func(ctx *fiber.Ctx) error {
		var mint mintInput
		if err := ctx.BodyParser(&mint); err != nil {
			return ctx.JSON("err parse")
		}

		amount, ok := new(big.Int).SetString(mint.Amount, 10)
		if !ok {
			return ctx.JSON("err parse ampunt")
		}

		if err := serv.MintTokens(ctx.Context(), amount); err != nil {
			return ctx.JSON(err.Error())
		}

		return ctx.JSON("ok")

	})

	app.Post("/contract/send", func(ctx *fiber.Ctx) error {
		var send sendInput
		if err := ctx.BodyParser(&send); err != nil {
			return ctx.JSON("err parse")
		}

		amount, ok := new(big.Int).SetString(send.Amount, 10)
		if !ok {
			return ctx.JSON("err parse amount")
		}

		to := common.HexToAddress(send.To)

		if err := serv.SendTokens(ctx.Context(), to, amount); err != nil {
			return ctx.JSON(err.Error())
		}

		return ctx.JSON("ok")
	})

	app.Listen(":8000")
}

type mintInput struct {
	Amount string `json:"amount"`
}

type sendInput struct {
	To     string `json:"to"`
	Amount string `json:"amount"`
}
