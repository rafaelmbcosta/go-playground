package book

import (
	"myproject/fiber_gorm_webservice/database"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book

	db.Find(&books)
	c.JSON(books)
}

// curl -X GET http://localhost:3000/api/v1/book/1
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

// curl  -X POST -H "Content-Type: application/json" --data "{\"title\": \"Angels and Demons\", \"author\": \"Dan Brown\", \"rating\": 5}"  http://localhost:3000/api/v1/book/
func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	// var book Book
	// book.Title = "1982"
	// book.Author = "Pablo Neruda"
	// book.Rating = 5
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book was found")
	}
	db.Delete(&book)
	c.Send("Book successfuly deleted")
}

func UpdateBook(c *fiber.Ctx) {
	c.Send("Update book")
}
