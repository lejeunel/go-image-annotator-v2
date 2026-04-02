package hasher

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256Hasher implements Hasher interface
type Sha256Hasher struct{}

// NewSha256Hasher returns a new instance of Sha256Hasher
func NewSha256Hasher() *Sha256Hasher {
	return &Sha256Hasher{}
}

// Hash computes the SHA-256 checksum of the input bytes and returns it as a hex string
func (h *Sha256Hasher) Hash(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}
