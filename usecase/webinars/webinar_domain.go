package webinars

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           int
	UserID       int
	CategoryID   int
	CategoryName string
	ImageUrl     string
	Name         string
	Description  string
	Price        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type Usecase interface {
	GetAll(ctx context.Context, name string) ([]Domain, error)
	GetByID(ctx context.Context, ID int) (Domain, error)
	GetByName(ctx context.Context, categoryName string) (Domain, error)
	Store(ctx context.Context, categoryDomain *Domain) (Domain, error)
	Update(ctx context.Context, newsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, categoriesDomain *Domain) (*Domain, error)
}

type Repository interface {
	GetAll(ctx context.Context, name string) ([]Domain, error)
	GetByID(ctx context.Context, ID int) (Domain, error)
	GetByName(ctx context.Context, categoryName string) (Domain, error)
	Store(ctx context.Context, categoriesDomain *Domain) (Domain, error)
	Update(ctx context.Context, categoriesDomain *Domain) (Domain, error)
	Delete(ctx context.Context, categoriesDomain *Domain) (Domain, error)
}