package test_client

import (
	"fmt"
)

func main() {
	var opts []grpc.DialOption
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		fmt.Println("Could not connect to client");
	}
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)
}