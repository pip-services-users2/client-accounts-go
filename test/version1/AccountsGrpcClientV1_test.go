package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-accounts-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type accountsGrpcCommandableClientV1Test struct {
	client  *version1.AccountGrpcClientV1
	fixture *AccountsClientFixtureV1
}

func newAccountsGrpcCommandableClientV1Test() *accountsGrpcCommandableClientV1Test {
	return &accountsGrpcCommandableClientV1Test{}
}

func (c *accountsGrpcCommandableClientV1Test) setup(t *testing.T) *AccountsClientFixtureV1 {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewAccountGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewAccountsClientFixtureV1(c.client)

	return c.fixture
}

func (c *accountsGrpcCommandableClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestGrpcCrudOperations(t *testing.T) {
	c := newAccountsGrpcCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
