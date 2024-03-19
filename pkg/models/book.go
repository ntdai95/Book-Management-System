package models

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/ntdai95/Book-Management-System/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	MYSQL_DB_BOOK_MANAGEMENT_SYSTEM := os.Getenv("MYSQL_DB_BOOK_MANAGEMENT_SYSTEM")
	if MYSQL_DB_BOOK_MANAGEMENT_SYSTEM == "" {
		log.Fatal("MYSQL_DB_BOOK_MANAGEMENT_SYSTEM is not found in the .env file")
	}

	config.Connect(MYSQL_DB_BOOK_MANAGEMENT_SYSTEM)
	db = config.GetDB()
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

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
