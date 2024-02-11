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
	fmt.Println(rq1, " -> ", rp1)

	rq2 := pb.CreateParams{
		Title: "T2",
		Content: "C2",
		Author: "A2",
		PublicationDate: "P2",
		Tags: []string{"Tg3", "Tg4"},
	}
	rp2, _ := client.Create(context.Background(), &rq2)
	fmt.Println(rq2, " -> ", rp2)

	rq3 := pb.ReadParams{
		PostID: 2,
	}
	rp3, _ := client.Read(context.Background(), &rq3)
	fmt.Println(rq3, " -> ", rp3);

	rq4 := pb.ReadParams{
		PostID: 1,
	}
	rp4, _ := client.Read(context.Background(), &rq4)
	fmt.Println(rq4, " -> ", rp4);

	rq5 := pb.ReadParams{
		PostID: 3,
	}
	rp5, _ := client.Read(context.Background(), &rq5)
	fmt.Println(rq5, " -> ", rp5);
}