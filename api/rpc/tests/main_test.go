package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/appcelerator/amp/api/server"
	"golang.org/x/net/context"

	//tested packages
	"github.com/appcelerator/amp/api/rpc/logs"
	"github.com/appcelerator/amp/api/rpc/service"
	"github.com/appcelerator/amp/api/rpc/stack"
	"github.com/appcelerator/amp/api/rpc/stats"
	"github.com/appcelerator/amp/api/rpc/topic"
)

var (
	ctx           context.Context
	statsClient   stats.StatsClient
	stackClient   stack.StackServiceClient
	topicClient   topic.TopicClient
	serviceClient service.ServiceClient
	logsClient    logs.LogsClient
)

func TestMain(m *testing.M) {

	//init server and context
	config, conn := server.StartTestServer()
	ctx = context.Background()

	//init package clients
	statsClient = stats.NewStatsClient(conn)
	stackClient = stack.NewStackServiceClient(conn)
	topicClient = topic.NewTopicClient(conn)
	serviceClient = service.NewServiceClient(conn)
	logsClient = logs.NewLogsClient(conn)

	//Init packages if needed
	if err := logsInit(config); err != nil {
		fmt.Printf("logInit error: %v\n", err)
		os.Exit(1)
	}
	initEtcd()
	influxInit(config)

	//start tests
	ret := m.Run()

	//clean up pakages if needed
	logsEnd()
	influxEnd()

	//exit
	os.Exit(ret)
}
