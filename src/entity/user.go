package entity

import "time"

type User struct {
	ID             string    `json:"uuid,omitempty" bson:"_id,omitempty"`
	Email          string    `json:"email" bson:"email"`
	Password       []byte    `bson:"password"`
	Roles          []string  `bson:"roles"`
	FirstName      string    `json:"firstName,omitempty" bson:"first_name,omitempty"`
	LastName       string    `json:"lastName,omitempty" bson:"last_name,omitempty"`
	CreatedAt      time.Time `json:"dateOfCreate,omitempty" bson:"created_at"`
	UpdatedAt      time.Time `json:"dateOfUpdate,omitempty" bson:"updated_at,omitempty"`
}