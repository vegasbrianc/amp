package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/appcelerator/amp/api/rpc/logs"
	"github.com/appcelerator/amp/api/rpc/service"
	"github.com/appcelerator/amp/api/rpc/stack"
	"github.com/appcelerator/amp/api/rpc/stats"
	"github.com/appcelerator/amp/api/rpc/topic"
)

var (
	amplifierEndpoint = flag.String("amplifier_endpoint", "localhost:8080", "endpoint of amplifier")
)

func run() (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = logs.RegisterLogsHandlerFromEndpoint(ctx, mux, *amplifierEndpoint, opts)
	if err != nil {
		return
	}
	err = service.RegisterServiceHandlerFromEndpoint(ctx, mux, *amplifierEndpoint, opts)
	if err != nil {
		return
	}
	err = stack.RegisterStackServiceHandlerFromEndpoint(ctx, mux, *amplifierEndpoint, opts)
	if err != nil {
		return
	}
	err = stats.RegisterStatsHandlerFromEndpoint(ctx, mux, *amplifierEndpoint, opts)
	if err != nil {
		return
	}
	err = topic.RegisterTopicHandlerFromEndpoint(ctx, mux, *amplifierEndpoint, opts)
	if err != nil {
		return
	}

	http.ListenAndServe(":3000", mux)
	return
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
