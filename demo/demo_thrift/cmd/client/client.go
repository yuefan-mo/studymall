package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"

	"github.com/yuefan-mo/studymall/demo/demo_thrift/kitex_gen/api"
	"github.com/yuefan-mo/studymall/demo/demo_thrift/kitex_gen/api/echo"
)

func main() {
	cli, err := echo.NewClient(
		"demo_thrift",
		client.WithHostPorts("localhost:8888"),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "demo_thrift_client",
		}),
	)
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}

	res, err := cli.Echo(context.Background(), &api.Request{
		Message: "hello",
	})
	if err != nil {
		log.Printf("调用 Echo 失败: %v", err)
		return
	}
	fmt.Printf("响应结果: %v\n", res)
}
