package io

import (
	"encoding/json"
	"os"

	"github.com/BenSnedeker/folio-core/pkg/folio"
)

// Save writes the Folio as a JSON file
func Save(f *folio.Folio, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(f)
}

// Load reads a JSON as a Folio
func Load(path string) (*folio.Folio, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var f folio.Folio
	dec := json.NewDecoder(file)
	if err := dec.Decode(&f); err != nil {
		return nil, err
	}
	return &f, nil
}
