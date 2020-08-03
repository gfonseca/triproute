package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"triproute/pkg/controller"
	"triproute/pkg/repository"
)

const dbFile = "./input-file.txt"

var dbRoutes repository.Repository

func main() {
	err := godotenv.Load()

	var DBFile string = os.Getenv("DBFILE")

	dbRoutes, err := repository.NewRepository(DBFile)

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n\nTripRoute find the best route...")

	for {
		fmt.Print("Please enter the route: ")
		input, _ := reader.ReadString('\n')
		inputs := strings.Split(strings.TrimSpace(input), "-")

		if len(inputs) != 2 {
			fmt.Println("Invalid route format. enter the route: ORG-DST")
			continue
		}
		req := controller.RequestBestRoute{Start: strings.TrimSpace(inputs[0]), End: strings.TrimSpace(inputs[1])}
		response, err := controller.GetBestRoute(req, dbRoutes)

		if err != nil {
			fmt.Printf("[Error] %s \n", err)
			continue
		}

		fmt.Println("best route: ", response.Route)
	}
}
