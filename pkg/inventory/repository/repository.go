package repository

import "github.com/codered-by-ec-council/Hands-on-Network-Programming-with-Go/pkg/models"

type Vendor interface {
	New(name string) error
	GetByID(id int) (models.Vendor, error)
	GetByName(name string) (models.Vendor, error)
	GetAll() ([]models.Vendor, error)
	Update(vendor models.Vendor) error
	Delete(id int) error
}
