package repository

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/models"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	Database Database
}

type RepositoryIFace interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(user models.User) (models.User, error)
	UserExists(email, username string) bool
}

func (r *Repository) CreateUser(user models.User) (models.User, error) {
	if err := r.Database.DB.Create(&user).Error; err != nil {
		return models.User{}, fmt.Errorf("%w:%w", ErrCreatingUser, err)
	}

	return user, nil
}

func (r *Repository) UserExists(email, username string) bool {
	var user models.User

	if err := r.Database.DB.Where("email = ? OR username = ?", email, username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
	}

	return true
}

func (r *Repository) GetUser(user models.User) (models.User, error) {
	var databaseUser models.User

	b, _ := json.Marshal(user)
	fmt.Println(string(b))
	if err := r.Database.DB.Where("id = ? OR username = ? OR email = ?", user.ID, user.Username, user.Email).First(&databaseUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, fmt.Errorf("%w:%w", ErrGettingUser, err)
		}
		return models.User{}, fmt.Errorf("%w:%w", ErrUnknown, err)
	}
	return databaseUser, nil
}
