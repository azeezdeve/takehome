package internal

import (
	"context"
	"log"
	"sync"
	"trustwallet/internal/repository"
	"trustwallet/pkg/blockchains/eth"
)

type (
	ETHParser struct {
		Transaction map[string][]Transaction
		addressRepo repository.IAddress
		ctx         context.Context
		mu          sync.Mutex
	}

	Transaction struct {
		ID string
	}

	Parser interface {
		// last parsed block
		GetCurrentBlock() int
		// add address to observer
		Subscribe(address string) bool
		// list of inbound or outbound transactions for an address
		GetTransactions(address string) []Transaction
	}
)

type Config func(e *ETHParser) error

func NewParser(cfgs ...Config) (*ETHParser, error) {
	conf := &ETHParser{
		Transaction: make(map[string][]Transaction),
	}

	for _, cfg := range cfgs {
		err := cfg(conf)
		if err != nil {

		}
	}
	return conf, nil
}

func (p ETHParser) GetCurrentBlock() int {
	blc := eth.New()
	resp, err := blc.GetCurrentBlock(p.ctx)
	if err != nil {
		log.Printf("error getting current block %+v", err)
		return 0
	}
	return resp
}

func (p ETHParser) Subscribe(address string) bool {
	if err := p.addressRepo.AddSubscribers(p.ctx, address); err != nil {
		return false
	}
	return true
}

func (p ETHParser) GetTransactions(address string) []Transaction {
	p.mu.Lock()
	defer p.mu.Unlock()
	transaction := p.Transaction[address]
	return transaction
}

func WithAddressRepository(addressRepo repository.IAddress) Config {
	return func(e *ETHParser) error {
		e.addressRepo = addressRepo
		return nil
	}
}

func WithCtx(ctx context.Context) Config {
	return func(e *ETHParser) error {
		e.ctx = ctx
		return nil
	}
}

func AddTransaction(address string, transaction ...Transaction) Config {
	return func(e *ETHParser) error {
		e.Transaction[address] = append(e.Transaction[address], transaction...)
		return nil
	}
}

var _ Parser = ETHParser{}
