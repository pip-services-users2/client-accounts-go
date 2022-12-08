package version1

import (
	"context"
	"reflect"
	"strings"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cerr "github.com/pip-services3-gox/pip-services3-commons-gox/errors"
	mdata "github.com/pip-services3-gox/pip-services3-data-gox/persistence"
)

type AccountsMemoryClientV1 struct {
	accounts    []AccountV1
	maxPageSize int
	proto       reflect.Type
}

func NewAccountsMemoryClientV1(accounts []AccountV1) *AccountsMemoryClientV1 {

	c := AccountsMemoryClientV1{
		accounts:    make([]AccountV1, 0),
		maxPageSize: 100,
		proto:       reflect.TypeOf(AccountV1{}),
	}
	if accounts != nil {
		c.accounts = append(c.accounts, accounts...)
	}
	return &c
}

func NewEmptyAccountsMemoryClientV1() *AccountsMemoryClientV1 {
	return NewAccountsMemoryClientV1(nil)
}

func (c *AccountsMemoryClientV1) GetAccounts(ctx context.Context, correlationId string, filter data.FilterParams,
	paging cdata.PagingParams) (result cdata.DataPage[*AccountV1], err error) {

	items := make([]*AccountV1, 0)
	for _, v := range c.accounts {
		item := v
		items = append(items, &item)
	}
	return *cdata.NewDataPage(items, len(c.accounts)), nil
}

func (c *AccountsMemoryClientV1) GetAccountById(ctx context.Context, correlationId string, id string) (result *AccountV1, err error) {

	for _, v := range c.accounts {
		if v.Id == id {
			buf := v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *AccountsMemoryClientV1) GetAccountByLogin(ctx context.Context, correlationId string, login string) (result *AccountV1, err error) {
	for _, v := range c.accounts {
		if v.Login == login {
			buf := v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *AccountsMemoryClientV1) GetAccountByIdOrLogin(ctx context.Context, correlationId string, idOrLogin string) (result *AccountV1, err error) {
	for _, v := range c.accounts {
		if v.Id == idOrLogin || v.Login == idOrLogin {
			buf := v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *AccountsMemoryClientV1) CreateAccount(ctx context.Context, correlationId string, account *AccountV1) (result *AccountV1, err error) {
	if account == nil {
		return nil, nil
	}

	var index = -1
	for i, v := range c.accounts {
		if v.Login == account.Login {
			index = i
			break
		}
	}

	if index >= 0 {
		err := cerr.NewBadRequestError(correlationId, "ACCOUNT_ALREADY_EXIST", "Account "+account.Login+" already exists")
		return nil, err
	}

	newItem := mdata.CloneObject(account, c.proto)
	mdata.GenerateObjectId(&newItem)
	item, _ := newItem.(AccountV1)

	c.accounts = append(c.accounts, item)

	return &item, nil
}

func (c *AccountsMemoryClientV1) UpdateAccount(ctx context.Context, correlationId string, account *AccountV1) (result *AccountV1, err error) {

	if account == nil {
		return nil, nil
	}

	var index = -1
	for i, v := range c.accounts {
		if v.Id == account.Id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	newItem := mdata.CloneObject(account, c.proto)
	item, _ := newItem.(AccountV1)
	c.accounts[index] = item
	return &item, nil

}

func (c *AccountsMemoryClientV1) DeleteAccountById(ctx context.Context, correlationId string, id string) (result *AccountV1, err error) {
	var index = -1
	for i, v := range c.accounts {
		if v.Id == id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}
	c.accounts[index].Deleted = true
	var item = c.accounts[index]
	return &item, nil
}

func (c *AccountsMemoryClientV1) matchString(value string, search string) bool {
	if value == "" && search == "" {
		return true
	}
	if value == "" || search == "" {
		return false
	}
	return strings.Contains(strings.ToLower(value), strings.ToLower(search))
}

func (c *AccountsMemoryClientV1) matchSearch(item AccountV1, search string) bool {
	search = strings.ToLower(search)
	return c.matchString(item.Name, search)
}

func (c *AccountsMemoryClientV1) composeFilter(filter *cdata.FilterParams) func(item AccountV1) bool {
	if filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	search := filter.GetAsString("search")
	id := filter.GetAsString("id")
	name := filter.GetAsString("name")
	login := filter.GetAsString("login")
	active, activeOk := filter.GetAsNullableBoolean("active")
	fromCreateTime, fromCreateTimeOK := filter.GetAsNullableDateTime("from_create_time")
	toCreateTime, toCreateTimeOk := filter.GetAsNullableDateTime("to_create_time")
	deleted := filter.GetAsBooleanWithDefault("deleted", false)

	return func(item AccountV1) bool {
		if search != "" && !c.matchSearch(item, search) {
			return false
		}
		if id != "" && id != item.Id {
			return false
		}
		if name != "" && name != item.Name {
			return false
		}
		if login != "" && login != item.Login {
			return false
		}
		if activeOk && active != item.Active {
			return false
		}
		if fromCreateTimeOK && item.CreateTime.Nanosecond() >= fromCreateTime.Nanosecond() {
			return false
		}
		if toCreateTimeOk && item.CreateTime.Nanosecond() < toCreateTime.Nanosecond() {
			return false
		}
		if !deleted && item.Deleted {
			return false
		}
		return true
	}
}
