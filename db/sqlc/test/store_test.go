package sqlc

import (
	"context"
	"project/e-commerce/db/sqlc"
	"project/e-commerce/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOrderTx(t *testing.T) {
	testStore := sqlc.NewStore(testDB)

	// create products 
	products := make([]sqlc.Product, 0)
	goodItem := 5
	goodPrice := 100.0
	for i := 0; i < goodItem; i++ {
		products = append(products, CreateRandomProduct(t, goodPrice, 100, 0))
	}

	// create a user and the cart
	count := int64(20)
	user := CreateRandomUser(t)
	cartId := CreateRandomCart(t, user)
	requestOrderItems := make([]sqlc.RequestOrderItem, 0)
	for i := 0; i < goodItem; i++ {
		requestOrderItems = append(requestOrderItems, sqlc.RequestOrderItem{
			CartID: cartId,
			ProductID: products[i].ID,
			Count:  count,
		})
	}

	// test order transactions
	result, err := testStore.OrderTx(context.Background(), sqlc.OrderParams{
		UserID: user.ID,
		Item:   requestOrderItems,
	})

	// check transaction
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Greater(t, len(result.ID), 1)
	totalPrice := utils.FloatToMoney(float64(count) * goodPrice * float64(goodItem))
	require.Equal(t, totalPrice, result.TotalAmount)
	require.Equal(t, count*int64(goodItem), result.TotalCount)
	require.False(t, result.IsPaied)
	require.WithinDuration(t, result.CreatedAt, time.Now(), time.Minute)
}

func TestOrderTxConcurrency(t *testing.T) {
	testStore := sqlc.NewStore(testDB)

	products := make([]sqlc.Product, 0)
	goodItem := 5
	for i := 0; i < goodItem; i++ {
		products = append(products, CreateRandomProduct(t, 100.0, 100, 0))
	}

	// run n concorrent transfer transcations
	n := 5
	count := int64(20)
	errs := make(chan error, n)
	results := make(chan *sqlc.OrderResult, n)

	for i := 0; i < n; i++ {
		go func() {
			// create a user and the cart
			user := CreateRandomUser(t)
			cartId := CreateRandomCart(t, user)
			requestOrderItems := make([]sqlc.RequestOrderItem, 0)

			for i := 0; i < goodItem; i++ {
				requestOrderItems = append(requestOrderItems, sqlc.RequestOrderItem{
					CartID: cartId,
					ProductID: products[i].ID,
					Count:  count,
				})
			}
			result, err := testStore.OrderTx(context.Background(), sqlc.OrderParams{
				UserID: user.ID,
				Item:   requestOrderItems,
			})
			results <- result
			errs <- err
		}()
	}

	// check transaction
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
		require.Greater(t, len(result.ID), 1)
	}
}
