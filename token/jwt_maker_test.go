package token

import (
	"project/e-commerce/utils"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker := NewJWTMaker(utils.RandomString(32))
	duration := time.Minute

	issuedTime := time.Now()
	expiredTime := issuedTime.Add(duration)

	token, payload, err := maker.CreateToken(utils.RandomInt(1, 100), duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	require.WithinDuration(t, payload.ExpiresAt, expiredTime, time.Second)
	require.WithinDuration(t, payload.IssuedAt, issuedTime, time.Second)
}

func TestExpiredJWT(t *testing.T) {
	maker := NewJWTMaker(utils.RandomString(32))

	token, payload, err := maker.CreateToken(utils.RandomInt(1, 100), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrExpiredToken)
	require.Nil(t, payload)
}

func TestInvalidJWT(t *testing.T) {
	payload, err := NewPayload(utils.RandomInt(1, 100), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker := NewJWTMaker(utils.RandomString(32))

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
