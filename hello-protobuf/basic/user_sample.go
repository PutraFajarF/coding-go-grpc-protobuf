package basic

import (
	"fmt"
	"hello-protobuf/protogen/basic"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {
	addr := basic.Address{
		Street:  "jalan beo",
		City:    "jakarta",
		Country: "indonesia",
	}
	u := basic.User{
		Id:       10,
		Username: "fajarf",
		IsActive: true,
		Password: []byte("fajaraja"),
		Emails:   []string{"fajarf@gmail.com", "fajf@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Adress:   &addr,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	p := basic.User{
		Id:       2,
		Username: "endi",
		IsActive: true,
		Password: []byte("endipass"),
		Emails:   []string{"endi@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}

func JsonToProtoUser() {
	json := `{
		"id": 98,
		"username": "budi",
		"is_active": true,
		"password": "budipass",
		"emails":[
			"budi@gmail.com",
			"budias@gmail.com"
		],
		"gender": "GENDER_MALE"
	}`

	var p basic.User

	err := protojson.Unmarshal([]byte(json), &p)
	if err != nil {
		fmt.Println(err)
	}

	log.Println(&p)
}
