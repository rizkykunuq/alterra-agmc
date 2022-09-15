package database

import (
	"submission/day2/config"
	"submission/day2/models"

	"gorm.io/gorm/clause"
)

func GetBooks() (interface{}, error) {
	var books []models.Books

	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func GetBookById(id int) (interface{}, error) {
	var book *models.Books

	if e := config.DB.Where("id = ?", id).Take(&book).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func CreateBook(book *models.Books) (interface{}, error) {
	if e := config.DB.Select("Title", "Author", "Year").Create(book).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func UpdateBookById(book *models.Books) (interface{}, error) {
	if e := config.DB.Model(&book).Clauses(clause.Returning{}).Where("id = ?", book.Id).Select("Title", "Author", "Year").Updates(book).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func DeleteBookById(id int) (interface{}, error) {
	var book *models.Books

	if e := config.DB.Where("id = ?", id).Delete(&book).Error; e != nil {
		return nil, e
	}
	return book, nil
}
