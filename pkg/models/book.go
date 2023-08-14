package models

import (
	"github.com/jinzhu/gorm"
	// "github.com/meghsha/g--bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func init() {
	config.connect()
	db = config.getDatabase()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Find(book)
	return book
}
