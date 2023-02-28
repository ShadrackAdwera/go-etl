package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateSession(t *testing.T) Session {

	uuid, err := uuid.NewRandom()
	user := CreateUser(t)

	require.NoError(t, err)

	_, refreshToken, err := testMaker.CreateToken(user.Username, user.ID, user.Email, time.Minute)

	require.NoError(t, err)

	sess, err := testQuery.CreateSession(context.Background(), CreateSessionParams{
		ID:           uuid,
		Username:     user.Username,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    time.Now().Add(time.Minute),
	})

	require.NoError(t, err)
	require.NotEmpty(t, sess)
	require.NotNil(t, sess.ID)
	return sess
}

func TestGetSession(t *testing.T) {
	sess := CreateSession(t)

	foundSess, err := testQuery.GetSession(context.Background(), sess.ID)

	require.NoError(t, err)
	require.NotEmpty(t, foundSess)
	require.Equal(t, sess.ID, foundSess.ID)
}
