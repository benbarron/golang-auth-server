package database

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Uid        	uuid.UUID       `gorm:"type:char(36);primaryKey" json:"uid"`
	CreatedAt 	time.Time		`gorm:"" json:"createdAt"`
	UpdatedAt 	time.Time		`gorm:"" json:"updatedAt"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"deletedAt"`
	Username 	string 			`gorm:"not null;unique" json:"username"`
	Password 	string 			`gorm:"not null" json:"password"`
	TokenStep	int				`gorm:"not null" json:"tokenStep`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Uid = uuid.NewV4()
	return nil
}

func (u *User) HashPassword() error {
	b := []byte(u.Password)
	hash, err := bcrypt.GenerateFromPassword(b, bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	hashedPassword := []byte(u.Password)
	plainTextPassword := []byte(password)
	if e := bcrypt.CompareHashAndPassword(hashedPassword, plainTextPassword); e == nil {
		return true
	}
	return false
}