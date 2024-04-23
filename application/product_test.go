package application_test

import (
	"github.com/murilloc/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {

	product := application.Product{}
	product.Name = "Product Test"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than 0 to enable the product", err.Error())

}

func TestProduct_Disable(t *testing.T) {

	product := application.Product{}
	product.Name = "Product Test"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be 0 to disable the product", err.Error())

}

func TestProduct_IsValid(t *testing.T) {

	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Product Test"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid status"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater than 0", err.Error())

	product.Price = 10
	_, err = product.IsValid()
	require.Nil(t, err)

}
