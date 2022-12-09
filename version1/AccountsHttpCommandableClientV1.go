package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cclients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type AccountsHttpCommandableClientV1 struct {
	*cclients.CommandableHttpClient
}

func NewAccountsHttpCommandableClientV1() *AccountsHttpCommandableClientV1 {
	return &AccountsHttpCommandableClientV1{
		CommandableHttpClient: cclients.NewCommandableHttpClient("v1/accounts"),
	}
}

func (c *AccountsHttpCommandableClientV1) GetAccounts(ctx context.Context, correlationId string, filter *data.FilterParams,
	paging *cdata.PagingParams) (result cdata.DataPage[*AccountV1], err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(ctx, "get_accounts", correlationId, params)
	if err != nil {
		return *cdata.NewEmptyDataPage[*AccountV1](), err
	}

	return cclients.HandleHttpResponse[cdata.DataPage[*AccountV1]](res, correlationId)
}

func (c *AccountsHttpCommandableClientV1) GetAccountById(ctx context.Context, correlationId string, id string) (result *AccountV1, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"account_id", id,
	)

	res, err := c.CallCommand(ctx, "get_account_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*AccountV1](res, correlationId)
}

func (c *AccountsHttpCommandableClientV1) GetAccountByLogin(ctx context.Context, correlationId string, login string) (result *AccountV1, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"login", login,
	)

	res, err := c.CallCommand(ctx, "get_account_by_login", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*AccountV1](res, correlationId)
}

func (c *AccountsHttpCommandableClientV1) GetAccountByIdOrLogin(ctx context.Context, correlationId string, idOrLogin string) (result *AccountV1, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"id_or_login", idOrLogin,
	)

	res, err := c.CallCommand(ctx, "get_account_by_id_or_login", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*AccountV1](res, correlationId)
}

func (c *AccountsHttpCommandableClientV1) CreateAccount(ctx context.Context, correlationId string, account *AccountV1) (result *AccountV1, err error) {

	if account.Id == "" {
		account.Id = cdata.IdGenerator.NextLong()
	}

	params := cdata.NewAnyValueMapFromTuples(
		"account", account,
	)

	res, err := c.CallCommand(ctx, "create_account", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*AccountV1](res, correlationId)
}

func (c *AccountsHttpCommandableClientV1) UpdateAccount(ctx context.Context, correlationId string, account *AccountV1) (result *AccountV1, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"account", account,
	)

	res, err := c.CallCommand(ctx, "update_account", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*AccountV1](res, correlationId)
}

func (c *AccountsHttpCommandableClientV1) DeleteAccountById(ctx context.Context, correlationId string, id string) (result *AccountV1, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"account_id", id,
	)

	res, err := c.CallCommand(ctx, "delete_account_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*AccountV1](res, correlationId)
}
