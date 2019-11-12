package entity

import (
	"encoding/json"
	"os"
)

// Storage persists data.
type Storage struct {
	Path string
}

// Load reads data from JSON file.
func (s *Storage) Load(v interface{}) error {
	f, err := os.Open(s.Path)
	if err != nil {
		return err
	}
	defer f.Close()
	// Read data from JSON file.
	return json.NewDecoder(f).Decode(v)
}

// Save writes data into JSON file.
func (s *Storage) Save(v interface{}) error {
	f, err := os.Create(s.Path)
	if err != nil {
		return err
	}
	defer f.Close()
	// Write data into JSON file.
	return json.NewEncoder(f).Encode(v)
}
