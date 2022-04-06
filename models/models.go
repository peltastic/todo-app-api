package models

import "time"

type User struct {
	Id       int    `json:"id" gorm:"autoIncrement"`
	Name     string `json:"name"`
	Email    string `gorm:"unique"`
	Password []byte `json:"-"`
}

type Todo struct {
	UserID      int    `json:"userid"`
	Todo        string `json:"todo"`
	TodoID      int    `json:"todoid"`
	IsCompleted bool   `json:"iscompleted"`
	CreatedAt   time.Time
}
