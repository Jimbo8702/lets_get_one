package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"time"

	"github.com/Jimbo8702/lets_get_one/types"
)

func GenerateToken(userID int, ttl time.Duration) (*types.Token, error) {
	token := &types.Token{
		UserID: userID,
		Expires: time.Now().Add(ttl),
	}
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	token.PlainText = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(token.PlainText))
	token.Hash = hash[:]
	return token, nil
}

func UserTokenValid(user *types.User) (bool, error) {
	if user.Token.PlainText == "" {
		return false, errors.New("no matching token found")
	}
	if user.Token.Expires.Before(time.Now()) {
		return false, errors.New("expired token")
	}
	return true, nil
}