package main

import (
	"os"

	"github.com/xiang-xx/go-aleo-sd/account"
)

func main() {
	pk := os.Getenv("pk")
	privateKey, err := account.ParsePrivateKey(pk)
	if err != nil {
		println(err.Error())
		return
	}
	println(privateKey.String())
}
