package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-accounts-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type AccountsCommandableHttpClientV1Test struct {
	client  *version1.AccountsCommandableHttpClientV1
	fixture *AccountsClientFixtureV1
}

func newAccountsCommandableHttpClientV1Test() *AccountsCommandableHttpClientV1Test {
	return &AccountsCommandableHttpClientV1Test{}
}

func (c *AccountsCommandableHttpClientV1Test) setup(t *testing.T) *AccountsClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewAccountsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewAccountsClientFixtureV1(c.client)

	return c.fixture
}

func (c *AccountsCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestHttpCrudOperations(t *testing.T) {
	c := newAccountsCommandableHttpClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
