package entity

import (
	"encoding/json"
	"os"
)

type Storage struct {
	Path string
}

func (s *Storage) Load(v interface{}) error {
	f, err := os.Open(s.Path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(v)
}

func (s *Storage) Save(v interface{}) error {
	f, err := os.Create(s.Path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(v)
}
