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

type AccountsMockClientV1 struct {
	accounts    []AccountV1
	maxPageSize int
	proto       reflect.Type
}

func NewAccountsMockClientV1(accounts []AccountV1) *AccountsMockClientV1 {

	c := AccountsMockClientV1{
		accounts:    make([]AccountV1, 0),
		maxPageSize: 100,
		proto:       reflect.TypeOf(AccountV1{}),
	}
	if accounts != nil {
		c.accounts = append(c.accounts, accounts...)
	}
	return &c
}

func NewEmptyAccountsMockClientV1() *AccountsMockClientV1 {
	return NewAccountsMockClientV1(nil)
}

func (c *AccountsMockClientV1) GetAccounts(ctx context.Context, correlationId string, filter *data.FilterParams,
	paging *cdata.PagingParams) (result cdata.DataPage[*AccountV1], err error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*AccountV1, 0)
	for _, v := range c.accounts {
		item := v
		if filterFunc(item) {
			items = append(items, &item)
		}
	}
	return *cdata.NewDataPage(items, len(c.accounts)), nil
}

func (c *AccountsMockClientV1) GetAccountById(ctx context.Context, correlationId string, id string) (result *AccountV1, err error) {

	for _, v := range c.accounts {
		if v.Id == id {
			buf := v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *AccountsMockClientV1) GetAccountByLogin(ctx context.Context, correlationId string, login string) (result *AccountV1, err error) {
	for _, v := range c.accounts {
		if v.Login == login {
			buf := v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *AccountsMockClientV1) GetAccountByIdOrLogin(ctx context.Context, correlationId string, idOrLogin string) (result *AccountV1, err error) {
	for _, v := range c.accounts {
		if v.Id == idOrLogin || v.Login == idOrLogin {
			buf := v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *AccountsMockClientV1) CreateAccount(ctx context.Context, correlationId string, account *AccountV1) (result *AccountV1, err error) {
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

func (c *AccountsMockClientV1) UpdateAccount(ctx context.Context, correlationId string, account *AccountV1) (result *AccountV1, err error) {

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

func (c *AccountsMockClientV1) DeleteAccountById(ctx context.Context, correlationId string, id string) (result *AccountV1, err error) {
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

func (c *AccountsMockClientV1) matchString(value string, search string) bool {
	if value == "" && search == "" {
		return true
	}
	if value == "" || search == "" {
		return false
	}
	return strings.Contains(strings.ToLower(value), strings.ToLower(search))
}

func (c *AccountsMockClientV1) matchSearch(item AccountV1, search string) bool {
	search = strings.ToLower(search)
	return c.matchString(item.Name, search)
}

func (c *AccountsMockClientV1) composeFilter(filter *cdata.FilterParams) func(item AccountV1) bool {
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

	ids := make([]string, 0)

	if idsStr := filter.GetAsString("ids"); idsStr != "" {
		ids = strings.Split(idsStr, ",")
	}

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
		if fromCreateTimeOK && item.CreateTime.Unix() >= fromCreateTime.Unix() {
			return false
		}
		if toCreateTimeOk && item.CreateTime.Unix() < toCreateTime.Unix() {
			return false
		}
		if !deleted && item.Deleted {
			return false
		}
		if len(ids) > 0 && !contains(ids, item.Id) {
			return false
		}
		return true
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
