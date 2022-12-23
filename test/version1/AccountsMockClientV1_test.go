package test_version1

import (
	"testing"

	"github.com/pip-services-users2/client-accounts-go/version1"
)

type accountsMockClientV1Test struct {
	client  *version1.AccountsMockClientV1
	fixture *AccountsClientFixtureV1
}

func newAccountsMockClientV1Test() *accountsMockClientV1Test {
	return &accountsMockClientV1Test{}
}

func (c *accountsMockClientV1Test) setup(t *testing.T) *AccountsClientFixtureV1 {

	c.client = version1.NewEmptyAccountsMockClientV1()
	c.fixture = NewAccountsClientFixtureV1(c.client)
	return c.fixture
}

func (c *accountsMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
	c.fixture = nil
}

func TestMockCrudOperations(t *testing.T) {
	c := newAccountsMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
