package sqlc

import (
	"context"
	"project/e-commerce/db/sqlc"
	"project/e-commerce/utils"
	"testing"

	"github.com/stretchr/testify/require"
)


func CreateRandomSeller(t *testing.T) (sqlc.Seller) {
	arg := sqlc.CreateSellerParams{
		SellerName: utils.RandomOwner(),
		BankAccount: utils.RandomString(5),
	}
	
	seller, err := testQueries.CreateSeller(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, seller)
	require.Equal(t,seller.SellerName, arg.SellerName)
	require.Equal(t,seller.BankAccount, arg.BankAccount)
	return seller
}

func TestCreateSeller(t *testing.T) {
	CreateRandomSeller(t)
}