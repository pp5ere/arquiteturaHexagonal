package cli

import (
	"fmt"

	"github.com/pp5ere/hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64) (string, error) {
	switch action {
	case "create":
		product, err := service.Create(productName, productPrice);if err != nil {
			return "", err
		}
		return fmt.Sprintf("The Product Id: %s Name: %s Price: %f Status: %s has been created",product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus()), nil
	
	case "enable":
		productInterface, err := service.Get(productId); if err != nil {
			return "", err
		}
		product, err := service.Enable(productInterface);if err != nil {
			return "", err
		}
		return fmt.Sprintf("The Product Id: %s Name: %s Price: %f Status: %s has been enabled",product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus()), nil
	case "disable":
		productInterface, err := service.Get(productId); if err != nil {
			return "", err
		}
		product, err := service.Disable(productInterface);if err != nil {
			return "", err
		}
		return fmt.Sprintf("The Product Id: %s Name: %s Price: %f Status: %s has been disabled",product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus()), nil
	default:
		product, err := service.Get(productId); if err != nil {
			return "", err
		}
		return fmt.Sprintf("Product Id: %s Name: %s Price: %f Status: %s",product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus()), nil
	}
}