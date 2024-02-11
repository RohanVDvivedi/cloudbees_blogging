package main

import (
	"fmt"
)

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

import (
	"cloudbees_blogging/pb"
)

var port int32 = 6959

func main() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), opts...)
	if err != nil {
		fmt.Println("Could not connect to client ", err);
		return
	}
	defer conn.Close()
	client := pb.NewBloggingServiceClient(conn)

	rq1 := pb.CreateParams{
		Title: "T1",
		Content: "C1",
		Author: "A1",
		PublicationDate: "P1",
		Tags: []string{"Tg1", "Tg2"},
	}
	rp1, _ := client.Create(context.Background(), &rq1)
	fmt.Println(rq1, " -> ", rp1);
}