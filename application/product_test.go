package application_test

import (
	"testing"
	uuid "github.com/satori/go.uuid"
	"github.com/pp5ere/hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestProductEnabled(t *testing.T)  {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10
	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be graater than zero to enable a product", err.Error())
}

func TestProductDisabled(t *testing.T)  {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0
	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be equal zero to have the product disabled", err.Error())
}

func TestProductIsValid(t *testing.T)  {
	product := application.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -1
	_, err =  product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())

}