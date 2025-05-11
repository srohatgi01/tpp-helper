package handlers

import (
	"encoding/json"
	"os"
)

// ReadJSONFile reads a JSON file from the given path and unmarshals it into the provided interface.
func ReadJSONFile(path string, v interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(v)
}
