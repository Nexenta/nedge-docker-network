package main

import (
	"github.com/Nexenta/nedge-docker-network/ndnet/ndnetapi"
	"github.com/Nexenta/nedge-docker-network/ndnet/ndnetcli"
	"os"
)

const (
	VERSION = "0.0.1"
)

var (
	client *ndnetapi.Client
)

func main() {
	ncli := ndnetcli.NewCli(VERSION)
	ncli.Run(os.Args)
}

