package xk6modules

import (
	"time"

	"github.com/zitadel/oidc/v3/pkg/client"
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
