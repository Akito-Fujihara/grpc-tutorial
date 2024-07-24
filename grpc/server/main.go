package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/Akito-Fujihara/grpc-tutorial/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedFileServiceServer
}

func (*server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	fmt.Println("ListFiles function was invoked with request: ", req)

	dir := "/Users/fujiwaraakihito/github/grpc-tutorial/grpc/storage"

	paths, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, path := range paths {
		if !path.IsDir() {
			filenames = append(filenames, path.Name())
		}
	}

	res := &pb.ListFilesResponse{
		Filenames: filenames,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterFileServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
