package main

import (
	"errors"
	"fmt"
	"os"

	"rpc"
	"thrift"
)

var (
	transportFactory    = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory     = thrift.NewTBinaryProtocolFactoryDefault()
	errUserOrPwdInvalid = errors.New("userid or password invalid")
)

const (
	remoteAddr = "192.168.1.241:8000"
)

type rpcService struct {
}

func (rs *rpcService) FindAll(userid int64, password string, param map[string]string) (res []string, err error) {
	if userid != 1 || password != "123456" || len(param) == 0 {
		err = errUserOrPwdInvalid
		return
	}

	for k, v := range param {
		res = append(res, k+":"+v)
	}

	return
}

func main() {
	defer os.Exit(1)

	serverSocket, err := thrift.NewTServerSocket(remoteAddr)
	if err != nil {
		fmt.Println("thrift.NewTServerSocket: ", err)
		return
	}
	defer serverSocket.Close()

	handler := &rpcService{}
	processor := rpc.NewRpcServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverSocket, transportFactory, protocolFactory)
	defer server.Stop()

	if err := server.Serve(); err != nil {
		fmt.Println("server.Server: ", err)
		return
	}
}
