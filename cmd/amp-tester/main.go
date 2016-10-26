package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// build vars
var (
	// Version is set with a linker flag (see Makefile)
	Version string
	// Build is set with a linker flag (see Makefile)
	Build   string
	modules = [...]string{"github.com/appcelerator/amp/api/rpc/service", "github.com/appcelerator/amp/data/storage/etcd", "github.com/appcelerator/amp/api/rpc/stack", "github.com/appcelerator/amp/data/influx", "github.com/appcelerator/amp/api/rpc/topic", "github.com/appcelerator/amp/api/rpc/logs"}
)

func main() {
	log.Printf("amp-tester (version: %s, build: %s)\n", Version, Build)

	var result []byte
	var args [3]string
	var err error
	args[0] = "test"
	args[1] = "-v"
	for _, module := range modules {
		args[2] = module
		fmt.Printf("Testing %s\n", module)
		cmd := exec.Command("go", args[:]...)
		result, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("output: %s\n", string(result[:]))
			fmt.Printf("test failed for %s\n", module)
			os.Exit(1)
		}
		fmt.Printf("output: %s\n", string(result[:]))
	}
	return
}
