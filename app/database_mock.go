// database_mock.go
package app

import (
	"database/sql"
	"github.com/stretchr/testify/mock"
)

// MockDatabase is a mock implementation of the Database interface.
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Begin() (*sql.Tx, error) {
	args := m.Called()
	return args.Get(0).(*sql.Tx), args.Error(1)
	// ... other methods for the mock
}

// MockResult is a mock implementation of the Result interface.
type MockResult struct {
	mock.Mock
}

func (mr *MockResult) LastInsertId() (int64, error) {
	result := mr.Called()
	return result.Get(0).(int64), result.Error(1)
}

func (mr *MockResult) RowsAffected() (int64, error) {
	result := mr.Called()
	return result.Get(0).(int64), result.Error(1)
}

func (m *MockDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	// Use the .Called method to record the function call and its arguments.
	// You can customize the return values for testing different scenarios.
	argsList := append([]interface{}{query}, args...)
	result := m.Called(argsList...)
	return result.Get(0).(sql.Result), result.Error(1)
}

func (m *MockDatabase) QueryRow(query string, args ...interface{}) *sql.Row {
	argsList := append([]interface{}{query}, args...)
	result := m.Called(argsList...)
	return result.Get(0).(*sql.Row)
}

func (m *MockDatabase) Commit() error {
	result := m.Called()
	return result.Error(0)
}

func (m *MockDatabase) Rollback() error {
	result := m.Called()
	return result.Error(0)
}

func (r *RealDatabase) Begin() (*sql.Tx, error) {
	return r.Begin()
	// Implement other database methods
}
