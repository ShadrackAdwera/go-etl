package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/ShadrackAdwera/go-etl/utils"
	"github.com/stretchr/testify/require"
)

func CreateFile(t *testing.T, user User) File {
	fileUrl := fmt.Sprintf("http://%s.fileurl.com", utils.RandomString(6))

	file, err := testQuery.CreateFile(context.Background(), CreateFileParams{
		FileUrl:     fileUrl,
		CreatedByID: user.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, file)
	require.NotZero(t, file.ID)
	require.Equal(t, file.FileUrl, fileUrl)
	require.Equal(t, file.CreatedByID, user.ID)

	return file
}

func TestGetFiles(t *testing.T) {
	user := CreateUser(t)
	n := 5

	for i := 0; i < n; i++ {
		CreateFile(t, user)
	}

	files, err := testQuery.GetFiles(context.Background(), user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, files)
	require.Equal(t, len(files), n)
}
