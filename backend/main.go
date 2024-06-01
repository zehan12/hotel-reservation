package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zehan12/hotel-reservation/backend/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUri  = "mongodb://localhost:27017"
	dbName = "hotel-reservation"
)

func main() {

	// database setup
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatalf("failed to connect with db: %v", err)
	}
	defer client.Disconnect(context.TODO())

	// verify mongoDB connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	client.Database(dbName)

	// create fiber app
	app := fiber.New()

	port := flag.String("port", ":3000", "The API server port")

	// group routes to v1
	apiv1 := app.Group("/api/v1")

	// index routes
	app.Get("/", handleHome)
	apiv1.Get("/", handleIndex)

	// user routes
	userRoutes := apiv1.Group("/user")
	userRoutes.Get("/", api.HandleGetUser)

	// server listing
	app.Listen(*port)
}

func handleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "server is running!"})
}

func handleIndex(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "welcome to hotel reservation index!"})
}
