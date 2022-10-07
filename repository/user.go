package repository

import (
	"github.com/kyh0703/stock-server/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Select(query string) []models.User
	SelectById(query string, id int64) models.User
	SelectByName(query string, name string) models.User
}

type userMysqlRepository struct {
	DB *gorm.DB
}

type NewUserRepository(db *gorm.DB) UserRepository {
	return &userMysqlRepository{DB: db}
}

func (m *userMysqlRepository) Select(query string) []models.Book {
	result := []models.Book{}
	m.DB.Raw(query).Scan(&result)
	return result
}

func (m *userMysqlRepository) SelectById(query string, id int64) models.Book {
	result := models.Book{}
	m.DB.Raw(query, id).Scan(&result)
	return result
}

func (m *userMysqlRepository) SelectByName(query string, name string) models.Book {
	result := models.Book{}
	m.DB.Raw(query, name).Scan(&result)
	return result
}