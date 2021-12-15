package application

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)


//run all tests : go test ./...

type ProductInterface interface{
	IsValid()(bool, error)
	Enable()error
	Disable()error
	GetId()string
	GetName()string
	GetPrice()float64
	GetStatus() string
}

const(
	DISABLED = "disabled"
	ENABLED = "enabled"
)

func init()  {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Product struct{
	Id 		string  `valid:"uuidv4"`
	Name 	string  `valid:"required"`
	Price 	float64 `valid:"float,optional"`
	Status 	string  `valid:"required"`
}

func NewProduct() *Product {
	return &Product{
		Id: uuid.NewV4().String(), 
		Status: DISABLED,
	}
}

func (p *Product) IsValid()(bool, error){
	if p.Status == ""{
		p.Status = DISABLED
	}
	if p.Status != ENABLED && p.Status != DISABLED{
		return false, errors.New("the status must be enabled or disabled")
	}
	if p.Price < 0 {
		return false, errors.New("the price must be greater or equal zero")
	}

	_, err := govalidator.ValidateStruct(p); if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable()error{
	if p.Price > 0 {
		p.Status = ENABLED
		return nil	
	}
	return errors.New("The price must be graater than zero to enable a product")
}

func (p *Product) Disable()error{
	if p.Price == 0{
		p.Status = DISABLED
		return nil	
	}
	return errors.New("The price must be equal zero to have the product disabled")
}

func (p *Product) GetId()string{
	return p.Id
}

func (p *Product) GetName()string{
	return p.Name
}

func (p *Product) GetPrice()float64{
	return p.Price
}

func (p *Product) GetStatus()string{
	return p.Status
}