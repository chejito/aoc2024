package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToArray(str string) []string {
	return strings.Fields(str)
}

func GetIntFromString(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error when converting:", err)
	}
	return num
}
