package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"awesomeProject1/grpcCommon/helloPb"
)

func main() {
	r := gin.Default()
	// addr : <svc-name>.<namespace>.<cluster....
	rpcClient := initGrpc("hello-grpc:80")

	r.GET("/ping", func(c *gin.Context) {
		var b dataStruct

		rpcRsp, err := rpcClient.SayHi(c, &helloPb.HiReq{Note: "hello"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		c.Bind(&b)
		c.JSON(200, gin.H{
			"message": rpcRsp.Note,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

type dataStruct struct {
	message string `form:"message"`
}

func initGrpc(addr string) helloPb.HelloerClient {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	// defer conn.Close() // 可以选择是长链接还是短连接
	return helloPb.NewHelloerClient(conn) // need generic

}
