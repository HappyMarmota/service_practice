package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"awesomeProject1/grpcCommon/helloPb"
)

type server struct {
	helloPb.UnimplementedHelloerServer
}

var (
	port = flag.Int("port", 8080, "The server port")
)

func (s *server) SayHi(ctx context.Context, req *helloPb.HiReq) (*helloPb.HiRsp, error) {
	rsp := req.Note + req.Note
	log.Println(req)
	return &helloPb.HiRsp{
		Note: rsp,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}

	s := grpc.NewServer()
	helloPb.RegisterHelloerServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
