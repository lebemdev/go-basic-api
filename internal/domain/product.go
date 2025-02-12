package domain

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"PrimaryKey"`
	CategoryId  uint      `json:"category_id"`
	Description string    `json:"description"`
	Img         string    `json:"img"`
	Price       float64   `json:"price"`
	Slug        string    `json:"slug" gorm:"index;unique;not null"`
	Stock       bool      `json:"stock"`
	Title       string    `json:"title"`
	UserId      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}
