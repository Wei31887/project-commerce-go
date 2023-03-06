package sqlc

import (
	"context"
	"math/rand"
	"project/e-commerce/db/sqlc"
	"project/e-commerce/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomProduct(t *testing.T, price float64, stock, sell int64) sqlc.Product {
	goodType := CreateRandomCategory(t)

	arg := sqlc.CreateProductParams{
		Name:        utils.RandomString(5),
		CategoryID:  goodType.ID,
		Stock:       stock,
		Sell:        sell,
		Price:       utils.FloatToMoney(price),
		OnSell:      utils.FloatToMoney(rand.Float64()),
		Description: utils.RandomString(20),
	}

	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, product.Name, arg.Name)
	require.Equal(t, product.Stock, arg.Stock)
	require.Equal(t, product.CategoryID, arg.CategoryID)
	require.Equal(t, product.Sell, arg.Sell)
	require.Equal(t, product.OnSell, arg.OnSell)
	// require.Equal(t, product.Price, arg.Price)
	require.Equal(t, product.Description, arg.Description)
	require.NotZero(t, product.ID)
	return product
}

func TestCreateProduct(t *testing.T) {
	CreateRandomProduct(t, utils.RandomFloat(1, 100), utils.RandomInt(0, 10), utils.RandomInt(0, 10))
}
