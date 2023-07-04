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
	rpcClient := initGrpc("hello-grpc.default.svc.cluster.local:8080")
	r.GET("/ping", func(c *gin.Context) {
		var b dataStruct
		rpcRsp, err := rpcClient.SayHi(nil, &helloPb.HiReq{Note: "hello"})
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
	defer conn.Close()
	return helloPb.NewHelloerClient(conn) // need generic

}
