package main

import (
	"context"
	pb "crud/guser"
	"crud/internal/config"
	"crud/internal/pkg/client/postgresql"
	"crud/internal/workdb"
	wr "crud/internal/workdb/db"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	pg workdb.Ways
}

func (s *server) AddUser(ctx context.Context, in *pb.User) (*wrappers.StringValue, error) {
	var wrk workdb.UserData
	var err error
	wrk.Name = in.Name
	wrk.Uid = uuid.NewString()
	wrk.Passwd, err = HashPassword(in.Passwd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wrk)
	if err := s.pg.InsertRow(ctx, &wrk); err != nil {
		return nil, err
	}

	return &wrappers.StringValue{Value: wrk.Uid}, nil
}

func (s *server) DleteUser(ctx context.Context, in *wrappers.Int64Value) (*wrappers.StringValue, error) {
	if err := s.pg.DeleteRow(ctx, in.Value); err != nil {
		return nil, err
	}

	return &wrappers.StringValue{Value: "Successful"}, nil
}

func (s *server) AllUsers(ctx context.Context, in *pb.Empty) (*pb.UsersList, error) {
	var ul = []*pb.User{}
	list, err := s.pg.AllRows(ctx)
	if err != nil {
		return nil, err
	}

	for _, l := range list {
		ul = append(ul, &pb.User{Num: l.Num, Uid: l.Uid, Name: l.Name, Passwd: l.Passwd})
	}

	return &pb.UsersList{Mas: ul}, nil
}

func main() {
	conf := config.GetConfig()
	host := fmt.Sprintf("%s:%s", conf.GrpcHost, conf.GrpcPort)
	l, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Println("[ERROR] tcp listen")
	}

	g := grpc.NewServer()

	db, err := postgresql.NewConnection(conf.CountConnect, conf)
	if err != nil {
		log.Fatal(err)
	}

	newpool := wr.NewpoolPGX(db)

	pb.RegisterUserServiceServer(g, &server{pg: newpool})

	if err = g.Serve(l); err != nil {
		fmt.Printf("[Error] gServe")
	}

}
