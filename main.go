package main

import (
	"fmt"
	"log"
	"time"

	"github.com/BenSnedeker/folio-core/internal/io"
	"github.com/BenSnedeker/folio-core/pkg/folio"
)

func main() {
	f := folio.Folio{
		ID:            "F001",
		SchemaVersion: "1.0",
		Name:          "Test Folio",
		Created:       time.Now(),
		LastModified:  time.Now(),
	}

	// Save
	err := io.Save(&f, "test.folio")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Saved Folio!")

	// Load
	loaded, err := io.Load("test.folio")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Loaded %+v\n", loaded)
}
