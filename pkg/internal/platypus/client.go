package platypus

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/plaid/plaid-go/plaid"
	"github.com/sirupsen/logrus"
)

type (
	Platypus interface {
		NewClientFromItemId(ctx context.Context, itemId string)
		NewClientFromLink(ctx context.Context, accountId uint64, linkId uint64)
	}

	Client interface {
		GetAccount(ctx context.Context, accountIds ...string) ([]BankAccount, error)
	}
)

var (
	_ Client = &PlaidClient{}
)

type PlaidClient struct {
	accountId   uint64
	linkId      uint64
	accessToken string
	log         *logrus.Entry
	client      *plaid.APIClient
}

func (p *PlaidClient) GetAccount(ctx context.Context, accountIds ...string) ([]BankAccount, error) {
	span := sentry.StartSpan(ctx, "Plaid - GetAccount")
	defer span.Finish()

	span.Data = map[string]interface{}{
		"accountIds": "ALL_BANK_ACCOUNTS",
	}

	if len(accountIds) > 0 {
		span.Data["accountIds"] = accountIds
	}

	return nil, nil
}
