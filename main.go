package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/big"
	"serv/config"
	"serv/service"
)

func main() {
	conf, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	serv, err := service.NewService(conf)
	if err != nil {
		log.Fatal(err)
	}

	if err := serv.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	if err := run(serv); err != nil {
		log.Fatal(err)
	}
}

func run(serv service.Service) error {
	app := fiber.New()

	app.Get("/contract/:addr", func(ctx *fiber.Ctx) error {
		addr := ctx.Params("addr")
		if addr == "" {
			return ctx.JSON("incorrect address")
		}

		address := common.HexToAddress(addr)

		balance, err := serv.GetBalance(ctx.Context(), address)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err)
		}

		return ctx.JSON(fmt.Sprintf("balance: %s", balance.String()))
	})

	app.Post("/contract", func(ctx *fiber.Ctx) error {
		address, err := serv.Deploy(ctx.Context())
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err)
		}

		return ctx.JSON(fmt.Sprintf("contract address: %s", address.String()))
	})

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

	return app.Listen(":8000")
}

type mintInput struct {
	Amount string `json:"amount"`
}

type sendInput struct {
	To     string `json:"to"`
	Amount string `json:"amount"`
}
