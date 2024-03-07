package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/deepmap/oapi-codegen/v2/examples/test/petstore"
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return
	}

	request, err := petstore.NewAddPetRequest("https://s3.example.com", petstore.AddPetJSONRequestBody{
		Name: "Meow",
	})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Println("Before signing")

	for k, v := range request.Header {
		fmt.Printf("%s: %s, ", k, v)
	}

	fmt.Println()
	fmt.Println()

	sigV4RequestEditor := NewSigV4RequestEditor(cfg)

	err = sigV4RequestEditor(context.Background(), request)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Println("After signing")

	for k, v := range request.Header {
		fmt.Printf("%s: %s, ", k, v)
	}

	fmt.Println()

	/*
		client, err := petstore.NewClient("example.com", petstore.WithRequestEditorFn(sigV4RequestEditor))
		if err != nil {
			return
		}

		pet, err := client.AddPet(context.Background(), petstore.AddPetJSONRequestBody{
			Name: "Meow",
		})
		if err != nil {
			return
		}

		fmt.Println(pet)
	*/
}
