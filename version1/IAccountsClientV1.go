package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IAccountsClientV1 interface {
	GetAccounts(ctx context.Context, correlationId string, filter data.FilterParams,
		paging data.PagingParams) (result data.DataPage[*AccountV1], err error)

	GetAccountById(ctx context.Context, correlationId string, id string) (result *AccountV1, err error)

	GetAccountByLogin(ctx context.Context, correlationId string, login string) (result *AccountV1, err error)

	GetAccountByIdOrLogin(ctx context.Context, correlationId string, idOrLogin string) (result *AccountV1, err error)

	CreateAccount(ctx context.Context, correlationId string, account *AccountV1) (result *AccountV1, err error)

	UpdateAccount(ctx context.Context, correlationId string, account *AccountV1) (result *AccountV1, err error)

	DeleteAccountById(ctx context.Context, correlationId string, id string) (result *AccountV1, err error)
}
