package repo

import (
	"context"
	"fmt"
	"log"
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
	defer testDgraph.MustClose()

	// Start testing
	code := m.Run()
	os.Exit(code)
}

//  TestDgraph
//-----------------------------------------------
type TestDgraph struct {
	conn   *grpc.ClientConn
	client *client.Dgraph
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
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Faild to connect to Dgraph server: %v", err)
	}

	return &TestDgraph{
		conn:   conn,
		client: client.NewDgraphClient(api.NewDgraphClient(conn)),
	}
}

func (d *TestDgraph) GetConn() *grpc.ClientConn {
	return d.conn
}

func (d *TestDgraph) MustAlter(t *testing.T, schema string) {
	if err := d.client.Alter(context.Background(), &api.Operation{Schema: schema}); err != nil {
		t.Fatalf("Failed to alter test Dgraph: %v", err)
	}
}

func (d *TestDgraph) MustCleanup(t *testing.T) {
	if err := d.Cleanup(); err != nil {
		t.Fatalf("Failed to cleanup test Dgraph: %v", err)
	}
}

func (d *TestDgraph) Cleanup() error {
	return d.client.Alter(context.Background(), &api.Operation{DropAll: true})
}

func (d *TestDgraph) MustClose() {
	var errs []error
	if err := d.Cleanup(); err != nil {
		errs = append(errs, err)
	}
	if err := d.conn.Close(); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		log.Fatalf("Failed to close connection with test Dgraph: %v", errs)
	}
}
