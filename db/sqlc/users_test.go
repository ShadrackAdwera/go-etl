package db

import (
	"context"
	"database/sql"
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

func TestFindUserByEmail(t *testing.T) {
	user := CreateUser(t)

	foundUser, err := testQuery.FindUserByEmail(context.Background(), user.Email)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.ID, foundUser.ID)
	require.Equal(t, user.Email, foundUser.Email)
	require.Equal(t, user.Username, foundUser.Username)
}

func TestFindUserById(t *testing.T) {
	user := CreateUser(t)

	foundUser, err := testQuery.GetUser(context.Background(), user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.ID, foundUser.ID)
	require.Equal(t, user.Email, foundUser.Email)
	require.Equal(t, user.Username, foundUser.Username)
}

func TestUpdateUsernameOnly(t *testing.T) {
	user := CreateUser(t)
	updatedUsername := utils.RandomString(6)

	updatedUser, err := testQuery.UpdateUser(context.Background(), UpdateUserParams{
		Username: sql.NullString{
			String: updatedUsername,
			Valid:  true,
		},
		ID: user.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.ID, updatedUser.ID)
	require.Equal(t, user.Email, updatedUser.Email)
	require.Equal(t, updatedUser.Username, updatedUsername)
	require.NotEqual(t, user.Username, updatedUser.Username)
}

func TestUpdatePasswordOnly(t *testing.T) {
	user := CreateUser(t)
	hashedPw, err := utils.HashPassword(utils.RandomString(8))
	require.NoError(t, err)

	updatedUser, err := testQuery.UpdateUser(context.Background(), UpdateUserParams{
		Password: sql.NullString{
			String: hashedPw,
			Valid:  true,
		},
		ID: user.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.ID, updatedUser.ID)
	require.Equal(t, user.Email, updatedUser.Email)
	require.Equal(t, updatedUser.Username, user.Username)
	require.NotEqual(t, user.Password, updatedUser.Password)
	require.Equal(t, updatedUser.Password, hashedPw)
}

func TestUpdateAllFields(t *testing.T) {
	user := CreateUser(t)
	updatedUsername := utils.RandomString(6)
	email := fmt.Sprintf("%s@mail.com", updatedUsername)
	hashedPw, err := utils.HashPassword(utils.RandomString(8))
	require.NoError(t, err)

	updatedUser, err := testQuery.UpdateUser(context.Background(), UpdateUserParams{
		Username: sql.NullString{
			String: updatedUsername,
			Valid:  true,
		},
		Email: sql.NullString{
			String: email,
			Valid:  true,
		},
		Password: sql.NullString{
			String: hashedPw,
			Valid:  true,
		},
		ID: user.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.ID, updatedUser.ID)
	require.NotEqual(t, user.Email, updatedUser.Email)
	require.NotEqual(t, updatedUser.Username, user.Username)
	require.NotEqual(t, user.Password, updatedUser.Password)
	require.Equal(t, updatedUser.Username, updatedUsername)
	require.Equal(t, updatedUser.Email, email)
	require.Equal(t, updatedUser.Password, hashedPw)
}
