package xk6modules

import (
	"time"

	"github.com/zitadel/oidc/v3/pkg/client"
	"github.com/zitadel/oidc/v3/pkg/crypto"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

// JwtFromKey creates a jwt token from the provided key and audience
// The function name in k6 is jwtFromKey
func (*Mod) JwtFromKey(key []byte, audience string) (string, error) {
	keyData, err := client.ConfigFromKeyFileData(key)
	if err != nil {
		return "", err
	}
	signer, err := client.NewSignerFromPrivateKeyByte([]byte(keyData.Key), keyData.KeyID)
	if err != nil {
		return "", err
	}
	return client.SignedJWTProfileAssertion(keyData.ClientID, []string{audience}, time.Hour, signer)
}

type JWTProfileRequest struct {
	Audience   oidc.Audience
	Expiration time.Duration
	Key        []byte
}

func (*Mod) SignJWTProfileAssertion(userID string, keyID string, request JWTProfileRequest) (string, error) {
	signer, err := client.NewSignerFromPrivateKeyByte(request.Key, string(keyID))
	if err != nil {
		return "", err
	}

	iat := time.Now()
	exp := iat.Add(5 * time.Second)
	return crypto.Sign(&oidc.JWTTokenRequest{
		Issuer:    userID,
		Subject:   userID,
		Audience:  request.Audience,
		ExpiresAt: oidc.FromTime(exp),
		IssuedAt:  oidc.FromTime(iat),
	}, signer)
}
