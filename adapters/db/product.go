package db

import (
	"database/sql"
	"github.com/pp5ere/hexagonal/application"
)

type ProductDb struct{
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	product := application.Product{}
	stmt, err := p.db.Prepare("select * from products where id = ?"); if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Status);if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb)Save(product application.ProductInterface)(application.ProductInterface, error){
	rows := 0
	p.db.QueryRow("select id from products where id = ?", product.GetId()).Scan(rows)
	if rows == 0{
		return p.insertProduct(product)
	}
	return p.updateProduct(product)
}

func (p * ProductDb) insertProduct (product application.ProductInterface)(application.ProductInterface, error){	
	smtm, err := p.db.Prepare("insert into products values(?,?,?,?)");if err != nil {
		return nil, err
	}
	_, err = smtm.Exec(product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus()); if err != nil {
		return nil, err
	}
	err = smtm.Close();if err != nil {
		return nil, err
	}
	return product, nil
}

func (p * ProductDb) updateProduct (product application.ProductInterface)(application.ProductInterface, error){		
	_, err := p.db.Exec("update products set(name=?, price=?, status=?) where id=?", 
						product.GetName(), product.GetPrice(), product.GetStatus(), product.GetId());if err != nil {
		return nil, err
	}
	return product, nil
}