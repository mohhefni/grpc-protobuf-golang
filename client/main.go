package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mohhefni/grpc-golang/common/config"
	"mohhefni/grpc-golang/common/model"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func serviceGarage() model.GaragesClient {
	port := config.ServiceGaragePort

	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.ServiceUserPort

	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "U001",
		Name:     "Moh. Hefni",
		Password: "Very Secret Password",
		Gender:   model.UserGender_MALE,
	}
	user2 := model.User{
		Id:       "U002",
		Name:     "Nur Syifa Fadila",
		Password: "Very Secret Password 2",
		Gender:   model.UserGender_FEMALE,
	}

	garage1 := model.Garage{
		Id:   "G001",
		Name: "Quel'thalas",
		Coordinate: &model.GarageCoordinate{
			Latitude: 45.123123123,
			Logitude: 54.1231313123,
		},
	}

	garage2 := model.Garage{
		Id:   "G002",
		Name: "Sumenep",
		Coordinate: &model.GarageCoordinate{
			Latitude: 46.123123123,
			Logitude: 56.1231313123,
		},
	}

	user := serviceUser()

	fmt.Printf("\n %s> user test\n", strings.Repeat("=", 10))

	// register user
	user.Register(context.Background(), &user1)
	user.Register(context.Background(), &user2)
	resp1, err := user.List(context.Background(), new(emptypb.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	jsonrResp1, _ := json.Marshal(resp1.List)
	log.Println(string(jsonrResp1))

	garage := serviceGarage()
	fmt.Printf("\n %s> garage test\n", strings.Repeat("=", 10))
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})

	reps2, err := garage.List(context.Background(), &model.GarageUserId{UserId: user1.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	jsonrreps2, _ := json.Marshal(reps2.List)
	log.Println(string(jsonrreps2))

	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user2.Id,
		Garage: &garage2,
	})

	reps3, err := garage.List(context.Background(), &model.GarageUserId{UserId: user2.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	jsonrreps3, _ := json.Marshal(reps3.List)
	log.Println(string(jsonrreps3))
}
