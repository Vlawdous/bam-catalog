package entity

import "github.com/google/uuid"

type Book struct {
	ID   uuid.UUID `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	Name string
}
