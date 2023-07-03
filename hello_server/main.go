package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"awesomeProject1/hello"
)

type server struct {
	hello.UnimplementedHelloerServer
}

var (
	port = flag.Int("port", 8080, "The Server Port")
)

func (s *server) SayHi(ctx context.Context, req *hello.HiReq) (*hello.HiRsp, error) {
	rsp := req.Note + req.Note
	log.Println(req)
	return &hello.HiRsp{
		Note: rsp,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}

	s := grpc.NewServer()
	hello.RegisterHelloerServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
