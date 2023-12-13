package db

import (
	"context"
	"testing"
	"time"

	"github.com/nehaal10/simeplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandOwner(),
		Balance:  util.RandomBalnce(),
		Currency: util.RandomCurrency(),
	}

	acc, err := testQuery.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

	return acc
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	account1, err := testQuery.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account1)

	require.Equal(t, account.ID, account1.ID)
	require.Equal(t, account.Balance, account1.Balance)
	require.Equal(t, account.Owner, account1.Owner)
	require.Equal(t, account.Currency, account1.Currency)

	require.WithinDuration(t, account.CreatedAt, account1.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomBalnce(),
	}

	acc, err := testQuery.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, account1.ID, acc.ID)
	require.Equal(t, account1.Owner, acc.Owner)
	require.Equal(t, account1.Currency, acc.Currency)
	require.WithinDuration(t, account1.CreatedAt, acc.CreatedAt, time.Second)
	require.Equal(t, arg.Balance, acc.Balance)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQuery.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	acc, err := testQuery.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, acc)
}

func TestDisplayAccount(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomAccount(t)
	}

	arg := DispAcoountParams{
		Offset: 5,
		Limit:  5,
	}

	acc, err := testQuery.DispAcoount(context.Background(), arg)

	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		require.NotEmpty(t, acc[i])
	}
}
