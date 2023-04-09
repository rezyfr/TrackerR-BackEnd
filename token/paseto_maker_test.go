package token

import (
	"testing"
	"time"

	"github.com/rezyfr/Trackerr-BackEnd/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	email := util.RandomEmail()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(email, util.RandomInt(1, 100), duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, email, payload.Email)
	require.True(t, issuedAt.Before(payload.ExpiredAt))
	require.True(t, expiredAt.After(payload.IssuedAt))
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpirePasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, payload, err := maker.CreateToken(util.RandomEmail(), util.RandomInt(1, 100), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err = maker.VerifyToken(token)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
