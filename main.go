package main

import (
	"context"
	"fmt"
	"trustwallet/internal"
	"trustwallet/internal/repository"
)

func main() {
	ctx := context.Background()
	addressRepo := repository.NewMemoDB()

	addrCfg := internal.WithAddressRepository(addressRepo)
	ctxCfg := internal.WithCtx(ctx)

	address := "0xtrustwallet"

	includeTransactions := []internal.Transaction{
		{ID: "1"},
		{ID: "2"},
		{ID: "3"},
	}

	transact := internal.AddTransaction(address, includeTransactions...)

	parser, err := internal.NewParser(addrCfg, ctxCfg, transact)
	if err != nil {
		panic(err)
	}
	//add transaction

	subscribed := parser.Subscribe(address)
	fmt.Println("is subscribe ", subscribed)

	transactions := parser.GetTransactions(address)
	fmt.Println("list of transactions ", transactions)

	currentBlk := parser.GetCurrentBlock()
	fmt.Println("current block on blockchain ", currentBlk)
}
