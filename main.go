package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"math/big"
	"os"
	"serv/config"
	"serv/service"
	"strconv"
	"strings"
)

func must[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}

	return a
}

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}

func run() error {
	serv, err := service.NewService(config.Conf{
		PkKey:           os.Getenv("pk_key"),
		ContractAddress: os.Getenv("contract_address"),
		ProviderHost:    os.Getenv("provider_host"),
		ProviderPort:    must(strconv.Atoi(os.Getenv("provider_port"))),
	})
	if err != nil {
		return err
	}

	if err := serv.Connect(context.Background()); err != nil {
		return err
	}

	app := fiber.New()

	app.Use(fiberlogger.New(fiberlogger.Config{
		Next: func(c *fiber.Ctx) bool {
			if strings.Contains(c.Path(), "app_ping") {
				return true
			}

			return false
		},
	}))

	// возвращает баланс пользователя по адресу
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

	// создает новый контракт и привязывает его к сервису
	app.Post("/contract", func(ctx *fiber.Ctx) error {
		address, err := serv.Deploy(ctx.Context())
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err)
		}

		return ctx.JSON(fmt.Sprintf("contract address: %s", address.String()))
	})

	// выпускает новые токены
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

	// переводит токены пользователю по адресу
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
