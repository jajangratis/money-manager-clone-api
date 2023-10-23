package controller

import (
	"github.com/jajangratis/money-manager-clone-api/model/domain"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"context"
	"github.com/jajangratis/money-manager-clone-api/helper"
	"github.com/jajangratis/money-manager-clone-api/model/web"
	_ "github.com/jajangratis/money-manager-clone-api/service"
	"github.com/stretchr/testify/mock"
)

// MockMstTransactionMethodService is a mock implementation of the service.MstTransactionMethodService interface.
type MockMstTransactionMethodService struct {
	mock.Mock
}

func (m *MockMstTransactionMethodService) Delete(ctx context.Context, id domain.MstTransactionMethod) {
	//TODO implement me
	panic("implement me")
}

func (m *MockMstTransactionMethodService) Create(ctx context.Context, request web.MstTransactionMethodRequest) web.MstTransactionMethodResponse {
	args := m.Called(ctx)
	return args.Get(0).(web.MstTransactionMethodResponse)
}

func (m *MockMstTransactionMethodService) FindAll(ctx context.Context) []web.MstTransactionMethodResponse {
	args := m.Called(ctx)
	return args.Get(0).([]web.MstTransactionMethodResponse)
}

func TestMstTransactionMethodController_Create(t *testing.T) {
	// Create a mock service and controller
	mockService := new(MockMstTransactionMethodService)
	controller := NewMstTransactionMethodController(mockService)
	requestBody := strings.NewReader(`{"method_name": "test", "method_id": "test"}`)
	// Create a request and response recorder for testing
	req, err := http.NewRequest("POST", "/api/mst-transaction-method", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Define the expected behavior of the mock service
	expectedData := web.MstTransactionMethodResponse{MethodId: "test", MethodName: "test"}
	mockService.On("Create", mock.Anything).Return(expectedData)

	// Call the FindAll method
	controller.Create(rr, req, nil)

	// Check the response
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Parse the response body (assuming it's JSON) for further validation
	var response = web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   rr.Body,
	}
	helper.WriteToResponseBody(rr, &rr.Body)

	// Validate the response content
	if response.Status != "Ok" {
		t.Errorf("Expected status 'Ok', but got '%s'", response.Status)
	}

	//// Assert the expected data
	//if len(response.Data) != len(expectedData) {
	//	t.Errorf("Expected %d data items, but got %d", len(expectedData), len(response.Data))
	//}

	// You can add more specific assertions as needed
}

func TestMstTransactionMethodController_FindAll(t *testing.T) {
	// Create a mock service and controller
	mockService := new(MockMstTransactionMethodService)
	controller := NewMstTransactionMethodController(mockService)

	// Create a request and response recorder for testing
	req, err := http.NewRequest("GET", "/api/mst-transaction-method", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Define the expected behavior of the mock service
	expectedData := []web.MstTransactionMethodResponse{{Id: "ExampleData", MethodId: "test", MethodName: "test"}}
	mockService.On("FindAll", mock.Anything).Return(expectedData)

	// Call the FindAll method
	controller.FindAll(rr, req, nil)

	// Check the response
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Parse the response body (assuming it's JSON) for further validation
	var response = web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   rr.Body,
	}
	helper.WriteToResponseBody(rr, &rr.Body)

	// Validate the response content
	if response.Status != "Ok" {
		t.Errorf("Expected status 'Ok', but got '%s'", response.Status)
	}

	//// Assert the expected data
	//if len(response.Data) != len(expectedData) {
	//	t.Errorf("Expected %d data items, but got %d", len(expectedData), len(response.Data))
	//}

	// You can add more specific assertions as needed
}
