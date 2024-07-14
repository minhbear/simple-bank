package db

import (
	"context"
	"simple_bank/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, account.Owner, args.Owner)
	require.Equal(t, account.Currency, args.Currency)
	require.Equal(t, account.Balance, args.Balance)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	get_account, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, get_account)
	require.Equal(t, get_account.ID, account.ID)
	require.Equal(t, get_account.Owner, account.Owner)
	require.Equal(t, get_account.Balance, account.Balance)
	require.Equal(t, get_account.Currency, account.Currency)
	require.WithinDuration(t, get_account.CreatedAt, account.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	args := UpdateAccountParams{
		ID:      account.ID,
		Balance: utils.RandomMoney(),
	}
	account_update, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account_update)
	require.Equal(t, account.ID, account_update.ID)
	require.Equal(t, account.Owner, account_update.Owner)
	require.Equal(t, account.Currency, account_update.Currency)
	require.Equal(t, args.Balance, account_update.Balance)
	require.WithinDuration(t, account_update.CreatedAt, account.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}