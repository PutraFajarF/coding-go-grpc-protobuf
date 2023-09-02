package basic

import (
	"hello-protobuf/protogen/basic"
	"log"
)

func BasicHello() {
	h := basic.Hello{
		Name: "Putra F",
	}

	log.Println(&h)
}
