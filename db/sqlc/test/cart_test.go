package sqlc

import (
	"context"
	"project/e-commerce/db/sqlc"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomCart(t *testing.T, user sqlc.User) (cartId int64) {
	cartId, err := testQueries.CreateCart(context.Background(), user.ID)
	require.NoError(t, err)

	require.Equal(t, cartId, user.ID)
	return cartId
}

func TestCreateCart(t *testing.T) {
	user := CreateRandomUser(t)
	CreateRandomCart(t, user)
}