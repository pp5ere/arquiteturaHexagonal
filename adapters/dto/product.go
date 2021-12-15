package dto

import (
	"github.com/pp5ere/hexagonal/application"
)

type Product struct{
	ID 		string 	`json:"id"`
	Name 	string	`json:"name"`
	Price 	float64	`json:"price"`
	Status 	string	`json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.Id = p.ID
	}
	
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status
	_, err := product.IsValid();if err != nil {
		return nil, err
	}
	return product, nil
}