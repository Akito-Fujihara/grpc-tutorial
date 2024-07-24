package main

import (
	"context"
	"fmt"

	"github.com/Akito-Fujihara/grpc-tutorial/grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	callListFiles(client)
}

func callListFiles(client pb.FileServiceClient) {
	req := &pb.ListFilesRequest{}
	res, err := client.ListFiles(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println("ListFiles Response: ", res)
}
