package main

import (
	"fmt"
	"os"

	"rpc"
	"thrift"
)

var (
	transportFactory = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory  = thrift.NewTBinaryProtocolFactoryDefault()
)

const (
	remoteAddr = "127.0.0.1:8000"
)

func main() {
	defer os.Exit(1)

	socket, err := thrift.NewTSocket(remoteAddr)
	if err != nil {
		fmt.Println("thrift.NewTSocket: ", err)
		return
	}
	defer socket.Close()

	transport := transportFactory.GetTransport(socket)

	client := rpc.NewRpcServiceClientFactory(transport, protocolFactory)

	// open
	if err := socket.Open(); err != nil {
		fmt.Println("socket.Open: ", err)
		return
	}

	res, err := client.FindAll(1, "123456", map[string]string{"name": "viney", "email": "viney.chow@gmail.com"})
	if err != nil {
		fmt.Println("client.FindAll:", err)
		return
	}

	fmt.Println("res: ", res)
}
