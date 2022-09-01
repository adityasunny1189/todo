package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Content string `json:"content"`
	Status  bool   `json:"status"`
}
