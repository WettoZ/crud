package main

import (
	"context"
	pb "crud/guser"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	// user := &pb.User{Name: "qqqq", Passwd: "1111"}
	// res, err := c.AddUser(ctx, user)
	// if err != nil {
	// 	fmt.Println("Error AddUser ", err)
	// }
	// fmt.Println(res)

	// userlist, err := c.AllUsers(ctx, &pb.Empty{})
	// if err != nil {
	// 	fmt.Println("Error AllUser ", err)
	// }

	// for _, us := range userlist.Mas {
	// 	fmt.Println(us)
	// }

	str, err := c.DleteUser(ctx, &wrapperspb.Int64Value{Value: 5})
	if err != nil {
		fmt.Println("Error DeleteUser ", err)
	}
	fmt.Println(str)

	// userlist, err = c.AllUsers(ctx, &pb.Empty{})
	// if err != nil {
	// 	fmt.Println("Error AllUser ", err)
	// }

	// for _, us := range userlist.Mas {
	// 	fmt.Println(us)
	// }
}
