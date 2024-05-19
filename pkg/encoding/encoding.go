package encoding

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

// EncodeCursor encodes createdAt and UUID into a base64 string
func EncodeCursor(t time.Time, id uuid.UUID) string {
	key := fmt.Sprintf("%s,%s", t.Format(time.RFC3339Nano), id.String())
	return base64.StdEncoding.EncodeToString([]byte(key))
}

// DecodeCursor decodes a base64 string into createdAt and UUID
func DecodeCursor(encodedCursor string) (time.Time, uuid.UUID, error) {
	decoded, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return time.Time{}, uuid.Nil, fmt.Errorf("invalid base64 string: %w", err)
	}

	arrStr := strings.Split(string(decoded), ",")
	if len(arrStr) != 2 {
		return time.Time{}, uuid.Nil, fmt.Errorf("invalid cursor format: %w", err)
	}

	createdAt, err := time.Parse(time.RFC3339Nano, arrStr[0])
	if err != nil {
		return time.Time{}, uuid.Nil, fmt.Errorf("invalid time format: %w", err)
	}
	id, err := uuid.Parse(arrStr[1])
	if err != nil {
		return time.Time{}, uuid.Nil, fmt.Errorf("invalid UUID format: %w", err)
	}
	return createdAt, id, nil
}
