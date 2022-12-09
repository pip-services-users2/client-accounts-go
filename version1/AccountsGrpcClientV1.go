package version1

import (
	"context"

	"github.com/pip-services-users2/client-accounts-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type AccountGrpcClientV1 struct {
	*clients.GrpcClient
}

func NewAccountGrpcClientV1() *AccountGrpcClientV1 {
	return &AccountGrpcClientV1{
		GrpcClient: clients.NewGrpcClient("accounts_v1.Accounts"),
	}
}

func (c *AccountGrpcClientV1) GetAccounts(ctx context.Context, correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result data.DataPage[*AccountV1], err error) {
	timing := c.Instrument(ctx, correlationId, "accounts_v1.get_account_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.AccountPageRequest{
		CorrelationId: correlationId,
	}

	if filter != nil {
		req.Filter = filter.Value()
	}

	if paging != nil {
		req.Paging = &protos.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  (int32)(paging.GetTake(100)),
			Total: paging.Total,
		}
	}

	reply := new(protos.AccountPageReply)
	err = c.CallWithContext(ctx, "get_accounts", correlationId, req, reply)
	if err != nil {
		return *cdata.NewEmptyDataPage[*AccountV1](), err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return *cdata.NewEmptyDataPage[*AccountV1](), err
	}

	result = toAccountPage(reply.Page)

	return result, nil
}

func (c *AccountGrpcClientV1) GetAccountById(ctx context.Context, correlationId string, id string) (result *AccountV1, err error) {
	timing := c.Instrument(ctx, correlationId, "accounts_v1.get_account_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.AccountIdRequest{
		CorrelationId: correlationId,
		AccountId:     id,
	}

	reply := new(protos.AccountObjectReply)
	err = c.CallWithContext(ctx, "get_account_by_id", correlationId, req, reply)

	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toAccount(reply.Account)

	return result, nil
}

func (c *AccountGrpcClientV1) GetAccountByLogin(ctx context.Context, correlationId string, login string) (result *AccountV1, err error) {
	timing := c.Instrument(ctx, correlationId, "accounts_v1.get_account_by_login")
	defer timing.EndTiming(ctx, err)

	req := &protos.AccountLoginRequest{
		CorrelationId: correlationId,
		Login:         login,
	}

	reply := new(protos.AccountObjectReply)
	err = c.CallWithContext(ctx, "get_account_by_login", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toAccount(reply.Account)

	return result, nil
}

func (c *AccountGrpcClientV1) GetAccountByIdOrLogin(ctx context.Context, correlationId string, idOrLogin string) (result *AccountV1, err error) {
	timing := c.Instrument(ctx, correlationId, "accounts_v1.get_account_by_id_or_login")
	defer timing.EndTiming(ctx, err)

	req := &protos.AccountLoginRequest{
		CorrelationId: correlationId,
		Login:         idOrLogin,
	}

	reply := new(protos.AccountObjectReply)
	err = c.Call("get_account_by_id_or_login", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toAccount(reply.Account)

	return result, nil
}

func (c *AccountGrpcClientV1) CreateAccount(ctx context.Context, correlationId string, account *AccountV1) (result *AccountV1, err error) {
	timing := c.Instrument(ctx, correlationId, "accounts_v1.create_account")
	defer timing.EndTiming(ctx, err)

	req := &protos.AccountObjectRequest{
		CorrelationId: correlationId,
		Account:       fromAccount(account),
	}

	reply := new(protos.AccountObjectReply)
	err = c.CallWithContext(ctx, "create_account", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toAccount(reply.Account)

	return result, nil
}

func (c *AccountGrpcClientV1) UpdateAccount(ctx context.Context, correlationId string, account *AccountV1) (result *AccountV1, err error) {
	timing := c.Instrument(ctx, correlationId, "accounts_v1.update_account")
	defer timing.EndTiming(ctx, err)

	req := &protos.AccountObjectRequest{
		CorrelationId: correlationId,
		Account:       fromAccount(account),
	}

	reply := new(protos.AccountObjectReply)
	err = c.CallWithContext(ctx, "update_account", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toAccount(reply.Account)

	return result, nil
}

func (c *AccountGrpcClientV1) DeleteAccountById(ctx context.Context, correlationId string, id string) (result *AccountV1, err error) {
	timing := c.Instrument(ctx, correlationId, "accounts_v1.delete_account_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.AccountIdRequest{
		CorrelationId: correlationId,
		AccountId:     id,
	}

	reply := new(protos.AccountObjectReply)
	err = c.CallWithContext(ctx, "delete_account_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toAccount(reply.Account)

	return result, nil
}
