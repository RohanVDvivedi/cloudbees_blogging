package main

import (
	"fmt"
)

import (
	"google.golang.org/grpc"
)

import (
	"cloudbees_blogging/pb"
)

var port int32 = 6959

func main() {
	var opts []grpc.DialOption
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), opts...)
	if err != nil {
		fmt.Println("Could not connect to client");
	}
	defer conn.Close()
	client := pb.NewBloggingServiceClient(conn)
}