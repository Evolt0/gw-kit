package base

import (
	"context"
	"github.com/Evolt0/def-kit/proto/book"
	"github.com/Evolt0/def-kit/proto/hello"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

//创建一个grpc客户端
//第二个参数是接口名称 .proto文件中的 BookFun
//第三个参数是方法名称 .proto文件中的 GetBookInfoByID
func GetEndpoint(cc *grpc.ClientConn, serviceName string, method string) *kitgrpc.Client {
	//参数grpcReply是grpc返回值，实际上是取类型
	return kitgrpc.NewClient(cc, serviceName, method, EncodeRequestHello, DecodeResponseHello, hello.HelloResp{})
}

//编码
func EncodeRequestHello(ctx context.Context, inter interface{}) (request interface{}, err error) {
	return inter, nil
}

//解码
func DecodeResponseHello(ctx context.Context, inter interface{}) (response interface{}, err error) {
	return inter, nil
}

func GetBookEndpoint(cc *grpc.ClientConn, serviceName string, method string) *kitgrpc.Client {
	//参数grpcReply是grpc返回值，实际上是取类型
	return kitgrpc.NewClient(cc, serviceName, method, EncodeRequestBook, DecodeResponseBook, book.BookResponse{})
}

//编码
func EncodeRequestBook(ctx context.Context, inter interface{}) (request interface{}, err error) {
	return inter, nil
}

//解码
func DecodeResponseBook(ctx context.Context, inter interface{}) (response interface{}, err error) {
	return inter, nil
}