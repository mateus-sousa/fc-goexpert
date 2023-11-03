package database

import (
	"github.com/mateus-sousa/goexpert/7-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Panda", "panda@email.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)
	var storedUser entity.User
	storedUser.ID = user.ID
	err = db.First(&storedUser).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, storedUser.ID)
	assert.Equal(t, user.Name, storedUser.Name)
	assert.Equal(t, user.Email, storedUser.Email)
	assert.NotNil(t, storedUser.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Panda", "panda@email.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)
	storedUser, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, storedUser.ID)
	assert.Equal(t, user.Name, storedUser.Name)
	assert.Equal(t, user.Email, storedUser.Email)
	assert.NotNil(t, storedUser.Password)
}
