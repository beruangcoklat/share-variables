package main

import (
	"fmt"
	"io"
	"net"

	sharedvariablespb "github.com/beruangcoklat/share-variables/proto"
	"google.golang.org/grpc"
)

type server struct{}

var streams []sharedvariablespb.ShareService_UpdateVariableServer

func (*server) UpdateVariable(stream sharedvariablespb.ShareService_UpdateVariableServer) error {
	streams = append(streams, stream)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			fmt.Println(err)
			return err
		}

		for _, s := range streams {
			err = s.Send(&sharedvariablespb.ShareResponse{
				Key:   req.GetKey(),
				Value: req.GetValue(),
			})

			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
}

func main() {
	streams = []sharedvariablespb.ShareService_UpdateVariableServer{}
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	s := grpc.NewServer()

	sharedvariablespb.RegisterShareServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
