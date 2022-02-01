package main

import (
	"context"
	pb "crud/guser"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
)

type server struct {
	DB *pgxpool.Pool
	pb.UnimplementedUserServiceServer
}

func (s *server) AddUser(ctx context.Context, in *pb.User) (*wrappers.StringValue, error) {
	var err error
	in.Uid = uuid.NewString()
	in.Passwd, err = HashPassword(in.Passwd)
	if err != nil {
		fmt.Println(err)
	}

	if err := InsertRow(ctx, s.DB, in); err != nil {
		return nil, err
	}

	return &wrappers.StringValue{Value: in.Uid}, nil
}

func (s *server) DleteUser(ctx context.Context, in *wrappers.Int64Value) (*wrappers.StringValue, error) {
	if err := DeleteRow(ctx, s.DB, in.Value); err != nil {
		return nil, err
	}

	return &wrappers.StringValue{Value: "Successful"}, nil
}

func (s *server) AllUsers(ctx context.Context, in *pb.Empty) (*pb.UsersList, error) {

	list, err := AllRows(ctx, s.DB)
	if err != nil {
		return nil, err
	}
	return &pb.UsersList{Mas: list}, nil
}

func main() {
	host := fmt.Sprintf("%s:%s", conf.GrpcHost, conf.GrpcPort)
	l, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Println("[ERROR] tcp listen")
	}

	g := grpc.NewServer()

	db, err := connection(conf.CountConnect)
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterUserServiceServer(g, &server{DB: db})

	if err = g.Serve(l); err != nil {
		fmt.Printf("[Error] gServe")
	}

}
