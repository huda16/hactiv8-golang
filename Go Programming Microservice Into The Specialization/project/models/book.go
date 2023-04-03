package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	NameBook  string    `gorm:"not null;type:varchar(191)" json:"name_book"`
	Author    string    `gorm:"not null;type:varchar(191)" json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Book) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Book Before Create()")

	if len(p.NameBook) < 1 {
		err = errors.New("book name is required")
	}

	if len(p.Author) < 1 {
		err = errors.New("book author is required")
	}

	if len(p.NameBook) > 191 {
		err = errors.New("book name is too long")
	}

	if len(p.Author) > 191 {
		err = errors.New("book author is too long")
	}

	return
}
