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

	fmt.Println("----TESTS STARTED----")

	rq1 := pb.CreateParams{
		Title: "T1",
		Content: "C1",
		Author: "A1",
		PublicationDate: "P1",
		Tags: []string{"Tg1", "Tg2"},
	}
	rp1, _ := client.Create(context.Background(), &rq1)
	fmt.Println(rq1, " -> ", rp1)
	if(rp1.PostID == 1){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	rq2 := pb.CreateParams{
		Title: "T2",
		Content: "C2",
		Author: "A2",
		PublicationDate: "P2",
		Tags: []string{"Tg3", "Tg4"},
	}
	rp2, _ := client.Create(context.Background(), &rq2)
	fmt.Println(rq2, " -> ", rp2)
	if(rp2.PostID == 2){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	rq3 := pb.ReadParams{
		PostID: 2,
	}
	rp3, _ := client.Read(context.Background(), &rq3)
	fmt.Println(rq3, " -> ", rp3)
	if(rp3.PostID == 2){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	rq4 := pb.ReadParams{
		PostID: 1,
	}
	rp4, _ := client.Read(context.Background(), &rq4)
	fmt.Println(rq4, " -> ", rp4)
	if(rp4.PostID == 1){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	rq5 := pb.ReadParams{
		PostID: 3,
	}
	rp5, _ := client.Read(context.Background(), &rq5)
	fmt.Println(rq5, " -> ", rp5)
	if(rp5.Error == "BLOG NOT FOUND"){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	rq6 := pb.UpdateParams{
		PostID: 2,
		Title: "T2_2",
		Content: "C2_2",
		Author: "A2_2",
		PublicationDate: "P2_2",
		Tags: []string{"Tg3_2", "Tg4_2"},
	}
	rp6, _ := client.Update(context.Background(), &rq6)
	fmt.Println(rq6, " -> ", rp6)
	if(rp6.Error != "BLOG NOT FOUND" && rp6.PostID == 2){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	rq7 := pb.ReadParams{
		PostID: 1,
	}
	rp7, _ := client.Read(context.Background(), &rq7)
	fmt.Println(rq7, " -> ", rp7)
	if(rp7.PostID == 1){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	rq8 := pb.ReadParams{
		PostID: 2,
	}
	rp8, _ := client.Read(context.Background(), &rq8)
	fmt.Println(rq8, " -> ", rp8)
	if(rp8.PostID == 2){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	rq9 := pb.UpdateParams{
		PostID: 3,
		Title: "T2_2",
		Content: "C2_2",
		Author: "A2_2",
		PublicationDate: "P2_2",
		Tags: []string{"Tg3_2", "Tg4_2"},
	}
	rp9, _ := client.Update(context.Background(), &rq9)
	fmt.Println(rq9, " -> ", rp9)
	if(rp9.Error == "BLOG NOT FOUND"){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	rq10 := pb.CreateParams{
		Title: "T3",
		Content: "C3",
		Author: "A3",
		PublicationDate: "P3",
		Tags: []string{"Tg5", "Tg6"},
	}
	rp10, _ := client.Create(context.Background(), &rq10)
	fmt.Println(rq10, " -> ", rp10)
	if(rp10.PostID == 3){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	rq11 := pb.DeleteParams{
		PostID: 2,
	}
	rp11, _ := client.Delete(context.Background(), &rq11)
	fmt.Println(rq11, " -> ", rp11)
	if(rp11.Error == ""){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	rq12 := pb.DeleteParams{
		PostID: 4,
	}
	rp12, _ := client.Delete(context.Background(), &rq12)
	fmt.Println(rq12, " -> ", rp12)
	if(rp12.Error == "BLOG NOT FOUND"){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	rq13 := pb.ReadParams{
		PostID: 1,
	}
	rp13, _ := client.Read(context.Background(), &rq13)
	fmt.Println(rq13, " -> ", rp13)
	if(rp13.PostID == 1){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	rq14 := pb.ReadParams{
		PostID: 2,
	}
	rp14, _ := client.Read(context.Background(), &rq14)
	fmt.Println(rq14, " -> ", rp14)
	if(rp14.Error == "BLOG NOT FOUND"){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	rq15 := pb.ReadParams{
		PostID: 3,
	}
	rp15, _ := client.Read(context.Background(), &rq15)
	fmt.Println(rq15, " -> ", rp15)
	if(rp15.PostID == 3){
		fmt.Println("PASS")
	}else{
		fmt.Println("FAIL")
		return
	}

	fmt.Println("----ALL TESTS PASSED----")
}