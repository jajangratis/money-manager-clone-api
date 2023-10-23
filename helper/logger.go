package helper

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello Logger")
	fmt.Println("Hello Logger")
}

func LogInfo(data any) {
	logger := logrus.New()
	logger.Info(data)
}

func LogError(data any) {
	logger := logrus.New()
	logger.Error(data)
}
