package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/IlhamEndianto/Hacktiv8/session-3/intro/model"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	user1 := &model.User{
		Id:       "u001",
		Name:     "Sylvana Windrunner",
		Password: "f0r Th3 H0rD3",
		Gender:   model.UserGender_FEMALE,
	}
	userList := &model.UserList{
		List: []*model.User{
			user1,
		},
	}
	log.Println("userList", userList)
	garage1 := &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 53.22033123,
		},
	}
	garageList := &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}
	garageListByUser := &model.GarageListByUser{
		List: map[string]*model.GarageList{
			user1.Id: garageList,
		},
	}

	log.Println("garageListByUser", garageListByUser)
	fmt.Printf("# ==== Original\n 		%#v \n", user1)
	fmt.Printf("# ==== As String\n 		%v \n", user1.String())

	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	jsonString := buf.String()
	fmt.Printf("# ==== As JSON String\n 		%v \n", jsonString)

	buf2 := strings.NewReader(jsonString)
	protoObject := new(model.GarageList)
	err2 := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}

	fmt.Printf("# ==== As String\n 		&v \n", protoObject.String())
}
