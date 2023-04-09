package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/rezyfr/Trackerr-BackEnd/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	user, err := testQueries.CreateUser(context.Background(), CreateUserParams{
		Email:    util.RandomEmail(),
		FullName: util.RandomString(5),
	})
	require.NoError(t, err)
	arg := CreateCategoryParams{
		UserID: user.ID,
		Type:   Transactiontype(util.RandomType()),
		Icon:   util.RandomString(5),
		Name:   util.RandomString(5),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Type, category.Type)
	require.Equal(t, arg.Icon, category.Icon)
	require.Equal(t, arg.UserID, category.UserID)
	require.Equal(t, arg.Name, category.Name)

	return category
}

func TestCreateCategory(t *testing.T) {
	user, err := testQueries.CreateUser(context.Background(), CreateUserParams{
		Email:    util.RandomEmail(),
		FullName: util.RandomString(5),
	})
	require.NoError(t, err)
	arg := CreateCategoryParams{
		UserID: user.ID,
		Type:   Transactiontype(util.RandomType()),
		Icon:   util.RandomString(5),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Type, category.Type)
	require.Equal(t, arg.Icon, category.Icon)
	require.Equal(t, arg.UserID, category.UserID)
}

func TestListCategories(t *testing.T) {
	var cat Category
	for i := 0; i < 3; i++ {
		cat = createRandomCategory(t)
	}

	categorys, err := testQueries.ListCategories(context.Background(), cat.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, categorys)

	for _, category := range categorys {
		require.NotEmpty(t, category)
	}
}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomCategory(t)

	arg := UpdateCategoryParams{
		ID:   category1.ID,
		Type: Transactiontype(util.RandomType()),
		Icon: util.RandomString(5),
	}

	category2, err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, arg.Type, category2.Type)
	require.Equal(t, arg.Icon, category2.Icon)
	require.Equal(t, arg.ID, category2.ID)
}

func TestDeleteCategory(t *testing.T) {
	category1 := createRandomCategory(t)

	err := testQueries.DeleteCategory(context.Background(), category1.ID)
	require.NoError(t, err)

	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.Error(t, err)
	require.Empty(t, category2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
