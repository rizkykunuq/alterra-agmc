package models

import (
	"time"
)

type Books struct {
	Id        int       `param:"id" json:"id" form:"id"`
	Title     string    `json:"title" form:"title"`
	Author    string    `json:"author" form:"author"`
	Year      int       `json:"year" form:"year"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
