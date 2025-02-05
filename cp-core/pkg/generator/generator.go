package generator

import (
	"crypto/rand"
	"fmt"
	"github.com/rs/xid"
	"math/big"
	"time"
)

const (
	letters         = "abcdefghijklmnopqrstuvwxyz"
	timestampFormat = "20060102150405"
	suffixLength    = 5
)

// GenerateRandomString 生成固定长度的随机字符串
func GenerateRandomString(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result[i] = letters[index.Int64()]
	}
	return string(result), nil
}

func GenerateUserID() string {
	id := xid.New()
	return id.String()
}

func GenerateSID(base string) (string, error) {
	suffix, err := GenerateRandomString(suffixLength)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s-%s-%s",
		base,
		time.Now().UTC().Format(timestampFormat),
		suffix,
	), nil
}
