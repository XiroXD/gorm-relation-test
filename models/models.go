package models

type Post struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	Title   string
	Content string
	Author  Author `gorm:"foreignKey:PostID;references:ID"`
}

type Author struct {
	Name   string
	Email  string
	PostID uint
}
