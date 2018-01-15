package repo

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/dgraph-io/dgraph/client"
	"github.com/dgraph-io/dgraph/protos/api"
	"google.golang.org/grpc"
)

//  Main
//-----------------------------------------------
var (
	testDgraph *TestDgraph
)

func TestMain(m *testing.M) {
	testDgraph = MustCreateTestDgraph()
	defer testDgraph.Cleanup()

	// Start testing
	code := m.Run()
	os.Exit(code)
}

//  TestDgraph
//-----------------------------------------------
type TestDgraph struct {
	GrpcPool
}

func MustCreateTestDgraph() *TestDgraph {
	var addr string
	if v, ok := os.LookupEnv("TEST_DGRAPH_GRPC_ADDR"); ok {
		addr = v
	} else if v, ok := os.LookupEnv("TEST_DGRAPH_GRPC_PORT"); ok {
		addr = fmt.Sprintf("localhost:%s", v)
	} else {
		addr = "localhost:9922" // default setting of the-dgraph-test container
	}

	return &TestDgraph{
		GrpcPool: NewGrpcPool(addr),
	}
}

func (d *TestDgraph) getClient(conn *grpc.ClientConn) *client.Dgraph {
	return client.NewDgraphClient(api.NewDgraphClient(conn))
}

func (d *TestDgraph) MustAlter(t *testing.T, schema string) {
	conn, err := d.Get()
	if err != nil {
		t.Fatalf("Failed to connect to Dgraph: %v", err)
	}
	defer conn.Close()
	if err := d.getClient(conn).Alter(context.Background(), &api.Operation{Schema: schema}); err != nil {
		t.Fatalf("Failed to alter test Dgraph: %v", err)
	}
}

func (d *TestDgraph) MustCleanup(t *testing.T) {
	if err := d.Cleanup(); err != nil {
		t.Fatalf("Failed to cleanup test Dgraph: %v", err)
	}
}

func (d *TestDgraph) Cleanup() error {
	conn, err := d.Get()
	if err != nil {
		return err
	}
	defer conn.Close()
	return d.getClient(conn).Alter(context.Background(), &api.Operation{DropAll: true})
}
