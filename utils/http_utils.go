package utils

import (
	"fmt"
	"strconv"
)

func ParseQueryParamToInt(param string) (int, error) {
	if param == "" {
		return -1, fmt.Errorf("")
	}

	num, err := strconv.Atoi(param)
	if err != nil {
		return -1, err
	}

	return num, nil
}
