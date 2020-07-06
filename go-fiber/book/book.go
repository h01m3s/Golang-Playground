package book

import (
	db "../db"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// Book type
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// GetBooks func
func GetBooks(c *fiber.Ctx) {
	db := db.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

// GetBook func
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DBConn
	var book Book
	db.Find(&book, id)

	if book.Title == "" {
		c.Status(400).Send("No book found with given ID")
		return
	}
	c.JSON(book)
}

// NewBook func
func NewBook(c *fiber.Ctx) {
	db := db.DBConn

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(400).Send(err)
		return
	}

	db.Create(&book)
	c.Status(201).JSON(book)
}

// DeleteBook func
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DBConn
	var book Book
	db.First(&book, id)

	if book.Title == "" {
		c.Status(400).Send("No book found with given ID")
		return
	}
	db.Delete(&book)
	c.Send("Book successfully deleted")

}
