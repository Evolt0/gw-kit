package base

import (
	"context"
	"fmt"
	"github.com/Evolt0/def-kit/proto/book"
	"github.com/Evolt0/def-kit/proto/hello"
	"github.com/Evolt0/gw-kit/pkg/controller/base"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

func Routes(root *gin.Engine) {
	base := root.Group("/base")
	base.GET("/hello", Hello)
	base.GET("/book/:id", Book)
}
func Hello(ctx *gin.Context) {
	conn, err := grpc.Dial(":666", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//创建endpoint，并指明grpc调用的接口名和方法名
	client := base.GetEndpoint(conn, "hello.HelloFun", "HelloWorld")
	resp, err := client.Endpoint()(context.Background(), &hello.HelloReq{})

	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": resp.(*hello.HelloResp).GetResp(),
	})
}

func Book(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusBadRequest)})
		return
	}
	conn, err := grpc.Dial(":666", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//创建endpoint，并指明grpc调用的接口名和方法名
	client := base.GetBookEndpoint(conn, "book.BookFun", "GetBookInfoByID")
	resp, err := client.Endpoint()(context.Background(), &book.BookRequest{Id: int32(id)})
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"name": resp.(*book.BookResponse).GetName(),
	})
}
