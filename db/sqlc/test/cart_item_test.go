package sqlc

import (
	"context"
	"project/e-commerce/db/sqlc"
	"project/e-commerce/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomCartItem(t *testing.T, product sqlc.Product, count, cartId int64) sqlc.CartItem {
	arg := sqlc.CreateCartItemParams{
		CartID: cartId,
		Count:  count,
		ProductID: product.ID,
	}

	cartItem, err := testQueries.CreateCartItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cartItem)

	require.Equal(t, cartItem.CartID, arg.CartID)
	require.Equal(t, cartItem.Count, arg.Count)
	require.Equal(t, cartItem.ProductID, arg.ProductID)

	return cartItem
}

func TestCreateCartItem(t *testing.T) {
	user := CreateRandomUser(t)
	product := CreateRandomProduct(t, utils.RandomFloat(1, 100), utils.RandomInt(1, 10), utils.RandomInt(1, 10))
	cartId := CreateRandomCart(t, user)
	CreateRandomCartItem(t, product, utils.RandomInt(1, product.Stock), cartId)
}
