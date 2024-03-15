package model

import "context"

// This interface will implemented by the respective blockchain
type IBlock interface {
	GetCurrentBlock(ctx context.Context) (int, error)
}
