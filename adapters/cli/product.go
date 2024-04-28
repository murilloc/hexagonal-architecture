package cli

import (
	"fmt"
	"github.com/murilloc/go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product ID %s with the name %s has been created with the price %f and status %s",
			product.GetId(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s has been enabled", res.GetId())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s has been disabled", res.GetId())
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has the price %f and status %s", res.GetId(), res.GetName(), res.GetPrice(), res.GetStatus())
	}

	return result, nil
}
