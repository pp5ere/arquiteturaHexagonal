package application

type ProductServiceInterface interface{
	Get(id string) (ProductInterface, error)
	Create(name string, price float64)(ProductInterface, error)
	Enable(p ProductInterface)(ProductInterface, error)
	Disable(p ProductInterface)(ProductInterface, error)
}

type ProductReader interface{
	Get(id string)(ProductInterface, error)
}

type ProductWriter interface{
	Save(p ProductInterface)(ProductInterface, error)
}

type ProductPersistenceInterface interface{
	ProductReader
	ProductWriter
}

type ProductService struct{
	Persistence ProductPersistenceInterface
}

func NewProductService (persistense ProductPersistenceInterface) *ProductService{
	return &ProductService{Persistence: persistense}
} 

func (s *ProductService) Get(id string) (ProductInterface, error) {
	return s.Persistence.Get(id)
}

func (s *ProductService) Create(name string, price float64)(ProductInterface, error) {
	p := NewProduct()
	p.Name = name
	p.Price = price
	_, err := p.IsValid();if err != nil {
		return nil, err
	}
	return s.Persistence.Save(p)
}

func (s *ProductService) Enable(p ProductInterface)(ProductInterface, error) {
	err := p.Enable(); if err != nil {
		return nil, err
	}
	return s.Persistence.Save(p)
}

func (s *ProductService) Disable(p ProductInterface)(ProductInterface, error) {
	err := p.Disable(); if err != nil {
		return nil, err
	}
	return s.Persistence.Save(p)
}