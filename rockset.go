package main

import (
	"fmt"
	"log"

	optional "github.com/antihax/optional"
	"github.com/rockset/rockset-go-client"
	models "github.com/rockset/rockset-go-client/lib/go"
)

func main() {

	// Client function is superseded by NewClient func
	
	client, err := rockset.NewClient(rockset.WithAPIKey("APIkey"))
	if err != nil {
		log.Fatal(err)
	}

	request := &models.ExecuteOpts{
		Body: optional.NewInterface(models.ExecuteQueryLambdaRequest{
			Parameters: []models.QueryParameter{
				{
					Name:  "country",
					Type_: "string",
					Value: "China",
				},
			},
		}),
	}

	res, _, err := client.QueryLambdas.Execute("commons", "SumStats", 4, request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query result: %v\n", res.Results[0])

	fmt.Println("Press Enter Key to exit")
	fmt.Scanln() // wait for Enter Key

}
