package encoding

import (
	"encoding/base64"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestEncodeCursor(t *testing.T) {
	createdAt := time.Now()
	id := uuid.New()

	encoded := EncodeCursor(createdAt, id)
	if encoded == "" {
		t.Errorf("Expected non-empty encoded string")
	}
}

func TestDecodeCursor(t *testing.T) {
	createdAt := time.Now().Truncate(time.Second)
	id := uuid.New()

	encoded := EncodeCursor(createdAt, id)
	decodedCreatedAt, decodedID, err := DecodeCursor(encoded)
	if err != nil {
		t.Fatalf("DecodeCursor returned error: %v", err)
	}

	if !decodedCreatedAt.Equal(createdAt) {
		t.Errorf("Expected createdAt %v, got %v", createdAt, decodedCreatedAt)
	}

	if decodedID != id {
		t.Errorf("Expected id %v, got %v", id, decodedID)
	}
}

func TestDecodedCursorInvalid(t *testing.T) {
	invalidEncoded := "invalid_base64_string"

	_, _, err := DecodeCursor(invalidEncoded)
	if err == nil {
		t.Fatalf("Expected error for invalid base64 string")
	}

	validEncoded := base64.StdEncoding.EncodeToString([]byte("invalid|format"))
	_, _, err = DecodeCursor(validEncoded)
	if err == nil {
		t.Fatalf("Expected error for invalid cursor format")
	}

	validEncoded = base64.StdEncoding.EncodeToString([]byte("2024-05-16T17:24:16Z|invalid-uuid"))
	_, _, err = DecodeCursor(validEncoded)
	if err == nil {
		t.Fatalf("Expected error for invalid UUID format")
	}
}
