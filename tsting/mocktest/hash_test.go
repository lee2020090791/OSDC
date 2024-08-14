package mocktest

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CalculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
func TestCalculateHash(t *testing.T) {
	data1 := "Hello, world"
	data2 := "Hello, world"

	hash1 := CalculateHash(data1)
	hash2 := CalculateHash(data2)

	assert.NotEqual(t, hash1, hash2)
}
