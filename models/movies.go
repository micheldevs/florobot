package models

import "time"

type Movie struct {
	ID           uint      `gorm:"unique;primaryKey;autoIncrement"`
	ExtId        string    `gorm:"ext_id"`
	ExtUrl       string    `gorm:"ext_url"`
	Title        string    `gorm:"title"`
	PosterUrl    string    `gorm:"poster_url"`
	Genres       string    `gorm:"genres"`
	Director     string    `gorm:"director"`
	Actors       string    `gorm:"actors"`
	Synopsis     string    `gorm:"synopsis"`
	Duration     string    `gorm:"duration"`
	PremiereDate time.Time `gorm:"premiere_date"`
	CreatedAt    time.Time
}
