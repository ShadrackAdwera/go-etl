package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/ShadrackAdwera/go-etl/utils"
	"github.com/stretchr/testify/require"
)

func CreateUser(t *testing.T) User {

	username := utils.RandomString(12)
	email := fmt.Sprintf("%s@mail.com", username)
	hashedPw, err := utils.HashPassword(utils.RandomString(8))

	require.NoError(t, err)

	newUserParams := CreateUserParams{
		Username: username,
		Email:    email,
		Password: hashedPw,
	}

	user, err := testQuery.CreateUser(context.Background(), newUserParams)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Username, username)
	require.Equal(t, user.Email, email)
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestFindUser(t *testing.T) {
	user := CreateUser(t)

	foundUser, err := testQuery.FindUserByEmail(context.Background(), user.Email)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.ID, foundUser.ID)
	require.Equal(t, user.Email, foundUser.Email)
	require.Equal(t, user.Username, foundUser.Username)
}
