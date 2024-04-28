package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/murilloc/go-hexagonal/adapters/cli"
	mock_application "github.com/murilloc/go-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productId := "abc"
	productStatus := "enabled"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
		productId,
		productName,
		productPrice,
		productStatus)

	result, err := cli.Run(service, "create", "", productName, productPrice)
	if err != nil {
		t.Error("Error was not expected")
	}

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s has been enabled", productName)
	result, err = cli.Run(service, "enable", productId, "", 0)

	resultExpected = fmt.Sprintf("Product ID %s has been disabled", productName)
	result, err = cli.Run(service, "disabled", productId, "", 0)

	resultExpected = fmt.Sprintf("Product ID %s with the name %s has the price %f and status %s",
		productId,
		productName,
		productPrice,
		productStatus)

	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}
