package main

import (
	"fmt"
	"myproject/fiber_gorm_webservice/book"
	"myproject/fiber_gorm_webservice/database"

	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"))
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connected!")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Migrated!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

// this is a tutorial from TutorialEdge on youtube.
func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
}
