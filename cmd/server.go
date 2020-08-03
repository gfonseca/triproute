package main

import (
	"fmt"
	"log"
	"os"
	"triproute/pkg/controller"
	"triproute/pkg/repository"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

var dbRoutes repository.Repository

func main() {
	err := godotenv.Load()

	var APPPort string = os.Getenv("PORT")
	var DBFile string = os.Getenv("DBFILE")

	dbRoutes, err := repository.NewRepository(DBFile)

	if err != nil {
		panic(fmt.Sprintf("Failed to start databse: %s", err))
	}

	app := fiber.New()
	app.Get("/findroute/", func(c *fiber.Ctx) {
		start := c.Query("start")
		end := c.Query("end")
		request := controller.RequestBestRoute{Start: start, End: end}
		response, err := controller.GetBestRoute(request, dbRoutes)

		log.Printf("Get: findroute/ %s", request)

		if err != nil {
			errResponse := controller.ErrorResponse{Msg: fmt.Sprintf("%s", err)}
			log.Printf("[Error]: findroute/ %s", errResponse)
			c.SendStatus(400)
			c.JSON(errResponse)
			return
		}

		c.SendStatus(200)
		c.JSON(response)
	})

	app.Post("/createroute/", func(c *fiber.Ctx) {
		requestRoute := new(controller.RequestInsertRoute)

		log.Printf("POST: createroute/ %s", requestRoute)

		if err := c.BodyParser(requestRoute); err != nil {
			log.Println(err)
		}

		err := controller.InsertNewRoute(*requestRoute, dbRoutes)

		if err != nil {
			errResponse := controller.ErrorResponse{Msg: fmt.Sprintf("%s", err)}
			log.Printf("[Error]: findroute/ %s", errResponse)
			c.SendStatus(400)
			c.JSON(errResponse)
			return
		}

		c.SendStatus(200)
		c.JSON(requestRoute)
	})

	app.Listen(APPPort)
}
