package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"

	"github.com/pip-services-users2/client-accounts-go/version1"
	"github.com/stretchr/testify/assert"
)

type AccountsClientFixtureV1 struct {
	Client version1.IAccountsClientV1

	ACCOUNT1 *version1.AccountV1
	ACCOUNT2 *version1.AccountV1
}

func NewAccountsClientFixtureV1(client version1.IAccountsClientV1) *AccountsClientFixtureV1 {
	ACCOUNT_ID1 := data.IdGenerator.NextLong()
	ACCOUNT_ID2 := data.IdGenerator.NextLong()
	return &AccountsClientFixtureV1{
		Client:   client,
		ACCOUNT1: version1.NewAccountV1(ACCOUNT_ID1, "Test Account "+ACCOUNT_ID1, ACCOUNT_ID1+"@conceptual.vision"),
		ACCOUNT2: version1.NewAccountV1(ACCOUNT_ID2, "Test Account "+ACCOUNT_ID2, ACCOUNT_ID2+"@conceptual.vision"),
	}
}

func (c *AccountsClientFixtureV1) clear() {
	page, _ := c.Client.GetAccounts(context.Background(), "", *data.NewEmptyFilterParams(), *data.NewEmptyPagingParams())

	for _, account := range page.Data {
		c.Client.DeleteAccountById(context.Background(), "", account.Id)
	}
}

func (c *AccountsClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create one account
	account, err := c.Client.CreateAccount(context.Background(), "", c.ACCOUNT1)
	assert.Nil(t, err)

	assert.NotNil(t, account)
	assert.Equal(t, account.Name, c.ACCOUNT1.Name)
	assert.Equal(t, account.Login, c.ACCOUNT1.Login)

	account1 := account

	// Create another account
	account, err = c.Client.CreateAccount(context.Background(), "", c.ACCOUNT2)
	assert.Nil(t, err)

	assert.NotNil(t, account)
	assert.Equal(t, account.Name, c.ACCOUNT2.Name)
	assert.Equal(t, account.Login, c.ACCOUNT2.Login)

	//account2 := account

	// Get all accounts
	page, err1 := c.Client.GetAccounts(context.Background(), "", *data.NewEmptyFilterParams(), *data.NewEmptyPagingParams())
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) >= 2)

	// Get account by login
	account, err = c.Client.GetAccountByIdOrLogin(context.Background(), "", c.ACCOUNT1.Login)
	assert.Nil(t, err)

	assert.NotNil(t, account)

	// Update the account
	account1.Name = "Updated Account 1"
	account, err = c.Client.UpdateAccount(context.Background(), "", account1)
	assert.Nil(t, err)

	assert.NotNil(t, account)
	assert.Equal(t, account.Name, "Updated Account 1")
	assert.Equal(t, account.Login, account1.Login)

	account1 = account

	// Delete account
	_, err = c.Client.DeleteAccountById(context.Background(), "", account1.Id)
	assert.Nil(t, err)

	// Try to get deleted account
	account, err = c.Client.GetAccountById(context.Background(), "", account1.Id)
	assert.Nil(t, err)

	assert.NotNil(t, account)
	assert.True(t, account.Deleted)
	//assert.Nil(t, account)
}
