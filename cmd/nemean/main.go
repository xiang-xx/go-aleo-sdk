package main

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
	"github.com/xiang-xx/go-aleo-sdk/rpc"
)

func main() {
	app := cli.NewApp()
	app.Name = "nemean"
	app.Usage = "An unfairly private wallet for Aleo."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "rpc",
			Value: "",
			Usage: "the host:port of SnarkOS",
		},
	}

	app.Commands = []cli.Command{
		getBlockCommand,
		getBlockHashCommand,
		getBlockHashesCommand,
		getBlockHeaderCommand,
		getBlockHeightCommand,
		getBlocksCommand,
		getBlocksTransactionsCommand,
		getCiphertextCommand,
		getLedgerProofCommand,
		getTransactionCommand,
		getTransitionCommand,
		latestBlockCommand,
		latestBlockHeaderCommand,
		latestBlockHeightCommand,
		latestBlockTransactionsCommand,
		latestLedgerRootCommand,
		sendTransactionCommand,
		newAccountCommand,
		fromAccountCommand,
		newTransactionCommand,
		newRecordCommand,
		encryptRecordCommand,
		decryptRecordCommand,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type profile struct {
	host string
	port string
}

func getProfile(ctx *cli.Context) (*profile, error) {
	snarkos := strings.Split(ctx.GlobalString("rpc"), ":")
	if len(snarkos) != 2 {
		return nil, errors.New("invalid rpc")
	}
	return &profile{
		host: snarkos[0],
		port: snarkos[1],
	}, nil
}

func getClient(host, port string) (*rpc.Client, error) {
	return rpc.NewClient(&rpc.Config{
		User:     "",
		Password: "",
		Host:     host,
		Port:     port,
	})
}
