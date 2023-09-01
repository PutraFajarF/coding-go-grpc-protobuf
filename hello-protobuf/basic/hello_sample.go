package basic

import (
	"hello-proto/protogen/basic"
	"log"
)

func BasicHello() {
	h := basic.Hello{
		Name: "Putra F",
	}

	log.Println(&h)
}
