package models

import (
	"github.com/jinzhu/gorm"
	"github.com/nwochaadim/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model

	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func (b *Book) UpdateBook() *Book {
	db.Save(&b)

	return b
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book

	db := db.Where("ID=?", id).Find(&book)

	return &book, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
