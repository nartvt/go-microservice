package model

import (
	"time"
)

type Product struct {
	Id        int
	Name      string
	Image     string
	Price     float32
	CreatedBy int
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Active    bool
}

type ProductEntity struct {
	Id        int
	Name      string
	Image     string
	CreatedBy int
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Active    bool
}
