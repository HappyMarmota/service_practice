package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"awesomeProject1/hello"
)

func main() {
	r := gin.Default()
	rpcClient := initGrpc("hello-grpc.default.svc.cluster.local:8080")
	r.GET("/ping", func(c *gin.Context) {
		var b dataStruct
		rpcRsp, err := rpcClient.SayHi(context.Background())
		c.Bind(&b)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	go r.Run() // listen and serve on 0.0.0.0:8080
}

type dataStruct struct {
	message string `form:"message"`
}

func initGrpc(addr string) hello.HelloerClient {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()
	return hello.NewHelloerClient(conn) // need generic

}
