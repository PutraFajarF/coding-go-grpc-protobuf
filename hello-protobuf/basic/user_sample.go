package basic

import (
	"hello-protobuf/protogen/basic"
	"log"
)

func BasicUser() {
	u := basic.User{
		Id:       10,
		Username: "fajarf",
		IsActive: true,
		Password: []byte("fajaraja"),
		Emails:   []string{"fajarf@gmail.com", "fajf@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}

	log.Println(&u)
}
