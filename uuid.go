// package uuid zbase32 encodes UUIDs
package uuid

import (
	"github.com/google/uuid"
	"github.com/tv42/zbase32"
)

// Encode z-base-32 encodes a UUID.
func Encode(uuid uuid.UUID) string {
	return zbase32.EncodeToString(uuid[:])
}

// Decode decodes a UUID from a string created by Encode.
func Decode(s string) (uuid.UUID, bool) {
	b, err := zbase32.DecodeString(s)
	if err != nil {
		return uuid.UUID{}, false
	}
	u, err := uuid.FromBytes(b)
	if err != nil {
		return uuid.UUID{}, false
	}
	return u, true
}
