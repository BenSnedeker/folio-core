package io

import (
	"os"
	"testing"
	"time"

	"github.com/BenSnedeker/folio-core/pkg/folio"
)

func TestSaveLoad(t *testing.T) {
	now := time.Now()

	f := folio.Folio{
		ID:            "F001",
		SchemaVersion: "1.0",
		Name:          "Test Folio",
		Created:       now,
		LastModified:  now,
	}

	// Save
	err := Save(&f, "test.folio")
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Load
	loaded, err := Load("test.folio")
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	// Verify basic fields carried
	if f.ID != loaded.ID {
		t.Errorf("ID mismatch: got %s, want %s", loaded.ID, f.ID)
	}
	if f.Name != loaded.Name {
		t.Errorf("Name mismatch: got %s, want %s", loaded.Name, f.Name)
	}
	if f.Created != now {
		t.Errorf("Created mismatch: got %s, want %s", loaded.Created, now)
	}
	if f.LastModified != now {
		t.Errorf("LastModified mismatch: got %s, want %s", loaded.LastModified, now)
	}

	// Clean up test file
	_ = os.Remove("test.folio")
}
