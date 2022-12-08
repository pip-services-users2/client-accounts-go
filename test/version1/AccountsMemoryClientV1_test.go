package test_version1

import (
	"testing"

	"github.com/pip-services-users2/client-accounts-go/version1"
)

type accountsMemoryClientV1Test struct {
	client  *version1.AccountsMemoryClientV1
	fixture *AccountsClientFixtureV1
}

func newAccountsMemoryClientV1Test() *accountsMemoryClientV1Test {
	return &accountsMemoryClientV1Test{}
}

func (c *accountsMemoryClientV1Test) setup(t *testing.T) *AccountsClientFixtureV1 {

	c.client = version1.NewEmptyAccountsMemoryClientV1()
	c.fixture = NewAccountsClientFixtureV1(c.client)
	return c.fixture
}

func (c *accountsMemoryClientV1Test) teardown(t *testing.T) {
	c.client = nil
	c.fixture = nil
}

func TestMemoryCrudOperations(t *testing.T) {
	c := newAccountsMemoryClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
