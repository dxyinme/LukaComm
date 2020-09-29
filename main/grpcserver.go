package main

import (
	"context"
	pb "github.com/dxyinme/LukaComm/TestPb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type grpcTestServer struct {

}

func (g *grpcTestServer) GrpcCall(ctx context.Context,in *pb.TestMsg) (*pb.TestMsgReply, error) {
	return &pb.TestMsgReply{Ans:"echo" + in.Name}, nil
}

func main(){
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterTestServerServer(s, &grpcTestServer{})
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}