package utility

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"pemuda-peduli/src/common/constants"
	"time"
)

// GeneratePass generate password dengan sha1 dari password dan salt yang diberikan
func GeneratePass(salt, password string) string {
	hash := sha1.New()
	io.WriteString(hash, salt+password)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// GenerateSalt generate salt string berdasarkan length yang diberikan
func GenerateSalt(length int) string {
	const charset = constants.SaltCharset

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
