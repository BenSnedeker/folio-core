package folio

import "testing"

func TestFolioCreation(t *testing.T) {
	f := Folio{
		ID:            "F001",
		SchemaVersion: "1.0",
		Name:          "Test",
	}

	if f.Name != "Test" {
		t.Errorf("Expected name \"Test\", got %s", f.Name)
	}
}
