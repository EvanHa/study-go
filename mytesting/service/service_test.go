// +build integration

package service_test

import (
	"mytesting/service"
	"mytesting/service/mocks"
	"testing"
)

// type smsServiceMock struct {
// 	mock.Mock
// }

// func (m *smsServiceMock) SendChargeNotification(value int) bool {
// 	fmt.Println("Mocked charge notification")
// 	args := m.Called(value)
// 	return args.Bool(0)
// }

func TestChargeCustomer(t *testing.T) {
	smsService := new(mocks.MessageService)
	smsService.On("SendChargeNotification", 100).Return(true)

	myService := service.MyService{smsService}

	myService.ChargeCustomer(100)

	smsService.AssertExpectations(t)
}
