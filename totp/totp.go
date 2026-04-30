package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"time"
)

func Generate(secret string) (string, error) {
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	if err != nil {
		return "", err
	}

	timeStep := time.Now().Unix() / 30

	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], uint64(timeStep))

	h := hmac.New(sha1.New, key)
	h.Write(buf[:])
	hash := h.Sum(nil)

	// reduce to 6 digits
	code := truncate(hash) % 1_000_000

	return fmt.Sprintf("%06d", code), nil
}

func truncate(hash []byte) int {
	// determine dynamic offset (last 4 bits of hash)
	offset := int(hash[len(hash)-1] & 0x0F)

	// extract 4 bytes starting at offset
	b0 := int(hash[offset]) & 0x7F // drop sign bit
	b1 := int(hash[offset+1]) & 0xFF
	b2 := int(hash[offset+2]) & 0xFF
	b3 := int(hash[offset+3]) & 0xFF

	// combine bytes into a 31-bit integer
	return (b0 << 24) |
		(b1 << 16) |
		(b2 << 8) |
		b3
}
