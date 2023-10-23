package repository

import (
	"github.com/jajangratis/money-manager-clone-api/app"
	"testing"
)

// Returns a new instance of MstCategoryRepositoryImpl
func test_new_mst_category_repository_impl(t *testing.T) {
	repository := NewMstCategoryRepositoryImpl()
	if repository == nil {
		t.Errorf("Expected a new instance of MstCategoryRepositoryImpl, but got nil")
	}
}

// The returned instance should not be nil.
func TestNewMstCategoryRepositoryImplNotNil(t *testing.T) {
	repository := NewMstCategoryRepositoryImpl()
	if repository == nil {
		t.Errorf("Expected instance of MstCategoryRepositoryImpl, got nil")
	}
}

// The function should be called multiple times and return different instances each time.
func test_new_mst_category_repository_impl_multiple_calls(t *testing.T) {
	instances := make(map[*MstCategoryRepositoryImpl]bool)

	for i := 0; i < 10; i++ {
		repo := NewMstCategoryRepositoryImpl()
		instances[repo] = true
	}

	if len(instances) != 10 {
		t.Errorf("Expected 10 different instances, but got %d", len(instances))
	}
}

//TEST USING TESTIFY MOCK

type MyService struct {
	DB app.Database
}

//func TestYourFunction(t *testing.T) {
//	// Create an instance of the mock
//	mockDB := new(app.MockDatabase)
//	repository := NewMstCategoryRepositoryImpl()
//	ctx := context.TODO()
//
//	// Define the expected behavior of the mock
//	mockDB.On("FindAll").Return(&sql.Tx{}, nil)
//
//	// Pass the mock to your service
//	myService := MyService{DB: mockDB}
//
//	// Call your function
//	err := repository.FindAll(ctx, myService)
//
//	// Assert and verify
//	if err != nil {
//		t.Errorf("Expected no error, but got an error: %v", err)
//	}
//
//	// Verify that the expected methods were called
//	mockDB.AssertExpectations(t)
//}
