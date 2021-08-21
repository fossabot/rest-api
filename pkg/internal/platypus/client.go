package platypus

import "context"

type Platypus interface {
	NewClientFromItemId(ctx context.Context, itemId string)
	NewClientFromLink(ctx context.Context, accountId uint64, linkId uint64)
}

type Client interface {
	GetAccount()
}
