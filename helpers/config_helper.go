package helpers

import (
	"encoding/json"
	"fmt"
	"os"
)

const config string = "files/days.json"

func ReadConfig() map[string]string {
	file, err := os.Open(config)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return make(map[string]string)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	var data map[string]string
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON file: %v\n", err)
		return make(map[string]string)
	}
	return data
}
