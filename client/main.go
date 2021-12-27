package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	"grpc/client/cat"
)

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("GRPC Dialing Error %s", err)
	}

	catClient := cat.NewClient(conn)

	res, err := catClient.GetCategory(3)
	if err != nil {
		log.Fatalf("Response Error to GetCategory %s", err)
	}

	fmt.Println(res)

	if err := catClient.GetCategories(); err != nil {
		log.Fatalf("Response Error to GetCategories %s", err)
	}

	if err := catClient.GetCat(); err != nil {
		log.Fatalf("Response Error to GetCategories %s", err)
	}
}
