package main

import (
	"context"
	"log"
	"mohhefni/grpc-golang/common/config"
	"mohhefni/grpc-golang/common/model"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

type UsersServer struct {
	model.UnimplementedUsersServer
}

func (UsersServer) Register(_ context.Context, param *model.User) (*empty.Empty, error) {
	user := param

	localStorage.List = append(localStorage.List, user)

	log.Println("Registering user", user.String())

	return new(empty.Empty), nil
}

func (UsersServer) List(context.Context, *empty.Empty) (*model.UserList, error) {
	return localStorage, nil
}

func main() {
	srv := grpc.NewServer()
	var userSrv UsersServer
	model.RegisterUsersServer(srv, userSrv)

	log.Println("Starting RPC server at", config.ServiceUserPort)
	l, err := net.Listen("tcp", config.ServiceUserPort)

	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.ServiceUserPort, err)
	}
	log.Fatal(srv.Serve(l))
}
