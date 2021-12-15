package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pp5ere/hexagonal/ports/cli"
	"github.com/pp5ere/hexagonal/application"
	mock_application "github.com/pp5ere/hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T)  {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := application.Product{
		Id: "abc",
		Name: "Product Test",
		Price: 32.1,
		Status: "enabled",
	}

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(product.Id).AnyTimes()
	productMock.EXPECT().GetStatus().Return(product.Status).AnyTimes()
	productMock.EXPECT().GetPrice().Return(product.Price).AnyTimes()
	productMock.EXPECT().GetName().Return(product.Name).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(product.Name, product.Price).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(product.Id).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	productExpected := fmt.Sprintf("The Product Id: %s Name: %s Price: %f Status: %s has been created",product.Id, product.Name, product.Price, product.Status)
	productResult, err := cli.Run(service, "create", "", product.Name, product.Price)
	require.Nil(t, err)
	require.Equal(t, productResult, productExpected)

	productExpected = fmt.Sprintf("The Product Id: %s Name: %s Price: %f Status: %s has been enabled",product.Id, product.Name, product.Price, product.Status)
	productResult, err = cli.Run(service, "enable", product.Id, "", 0)
	require.Nil(t, err)
	require.Equal(t, productResult, productExpected)

	productExpected = fmt.Sprintf("The Product Id: %s Name: %s Price: %f Status: %s has been disabled",product.Id, product.Name, product.Price, product.Status)
	productResult, err = cli.Run(service, "disable", product.Id, "", 0)
	require.Nil(t, err)
	require.Equal(t, productResult, productExpected)

	productExpected = fmt.Sprintf("Product Id: %s Name: %s Price: %f Status: %s",product.Id, product.Name, product.Price, product.Status)
	productResult, err = cli.Run(service, "get", product.Id, "", 0)
	require.Nil(t, err)
	require.Equal(t, productResult, productExpected)
}