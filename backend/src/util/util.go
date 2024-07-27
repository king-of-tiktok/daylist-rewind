package util

import (
	"crypto/md5"
	"daylist-rewind-backend/src/global"
	"encoding/hex"
	"fmt"
	"log/slog"
	"math/rand"
	"reflect"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ShuffleArray shuffles an array in place.
func ShuffleArray(slice interface{}) {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		panic("Shuffle: not a slice")
	}
	n := rv.Len()
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		tmp := reflect.ValueOf(rv.Index(i).Interface())
		rv.Index(i).Set(rv.Index(j))
		rv.Index(j).Set(tmp)
	}
}

// FormatTime formats a time.Time object as a string in the format "2006-01-02 15:04:05.000Z".
func FormatTime(t time.Time) string {
	utcTime := t.UTC()
	layout := "2006-01-02 15:04:05.000Z"
	return utcTime.Format(layout)
}

// GetMD5Hash returns the MD5 hash of a string.
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func RandomString(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func ValidateJWT() bool {
	// global.AdminToken = ""

	type TokenClaims struct {
		Exp  int64  `json:"exp"`
		ID   string `json:"id"`
		Type string `json:"type"`
	}

	// claims := jwt.MapClaims{}
	token, err := jwt.Parse(global.AdminToken, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key for validation
		return []byte("secret"), nil
	})
	if err != nil {
		if err.Error() != "token signature is invalid: signature is invalid" {
			slog.Error("Error parsing token: ",
				"error", err.Error(),
			)
			return false
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		slog.Info("Token: ", "token", claims)

		// Map claims to struct
		tokenClaims := TokenClaims{
			Exp: int64(claims["exp"].(float64)),
			// ID:   claims["id"].(string),
			// Type: claims["type"].(string),
		}

		exp := time.Unix(tokenClaims.Exp, 0)
		now := time.Now()
		return !exp.Before(now)
	} else {
		slog.Error("Invalid token claims")
		return false
	}
}
