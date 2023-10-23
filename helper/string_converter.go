package helper

import (
	"fmt"
	"strconv"
)

func StringToInt(s string) int {
	if s != "" {
		sConverted, err := strconv.Atoi(s)
		PanicIfError(err)
		return sConverted
	} else {
		return 0
	}
}

func MapStringToStrings(input map[string]interface{}) (map[string]string, error) {
	output := make(map[string]string)
	for key, value := range input {
		strValue, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("unable to convert %v to string", value)
		}
		output[key] = strValue
	}
	return output, nil
}
