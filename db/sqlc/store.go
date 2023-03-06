package sqlc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// Store interface for order transactions
type Store interface {
	Querier
	OrderTx(context.Context, OrderParams) (*OrderResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}

// for executing transactions
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	qu := New(tx)
	err = fn(qu)
	if err != nil {
		// if err is not nil, rollback the transaction
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type RequestOrderItem struct {
	CartID int64 `json:"cart_id" binding:"required"`
	ProductID int64 `json:"product_id" binding:"required"`
	Count  int64 `json:"count" binding:"required"`
}

type OrderParams struct {
	UserID int64              `json:"user_id" binding:"required"`
	Item   []RequestOrderItem `json:"request_order_item"`
}

type OrderResult struct {
	ID          uuid.UUID `json:"id"`
	UserID      int64     `json:"user_id"`
	TotalAmount string    `json:"total_amount"`
	TotalCount  int64     `json:"total_count"`
	IsPaied     bool      `json:"is_paied"`
	CreatedAt   time.Time `json:"created_at"`
}

// OrderTx: Create a order from one cart with each order item created
// first create the order, then create each order item from the cart and balance the product status
// finally update the order status.
func (store *SQLStore) OrderTx(ctx context.Context, arg OrderParams) (*OrderResult, error) {
	var result OrderResult

	err := store.execTx(ctx, func(q *Queries) error {
		// create the order
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		req := CreateOrderParams{
			ID:      id,
			UserID:  arg.UserID,
			IsPaied: false,
		}
		order, err := store.CreateOrder(ctx, req)
		if err != nil {
			return err
		}

		var totalAmount float64 = 0
		var totalCount int64 = 0
		for _, item := range arg.Item {
			itemReq := CreateOrderItemParams{
				OrderID: order.ID,
				ProductID:  item.ProductID,
				Count:   item.Count,
			}

			// create the order items
			_, err := store.CreateOrderItem(ctx, itemReq)
			if err != nil {
				return err
			}

			// balance product status
			goodReq := UpdateProductStockSellParams{
				ID:        item.ProductID,
				Count:     -item.Count,
				SellCount: item.Count,
				UpdatedAt: time.Now(),
			}
			product, err := store.UpdateProductStockSell(ctx, goodReq)
			if err != nil || product.Stock < 0 {
				if product.Stock < 0 {
					return errors.New("run out of stock")
				}
				return err
			}

			// calculate total price and total item
			price, _ := strconv.ParseFloat(product.Price, 64)
			totalAmount += price * float64(item.Count)
			totalCount += item.Count
		}

		// update the order status
		resAmount := strconv.FormatFloat(totalAmount, 'E', -1, 64)
		updateOrderReq := UpdateOrderInfoParams{
			ID:          order.ID,
			TotalAmount: resAmount,
			TotalCount:  totalCount,
		}

		order, err = store.UpdateOrderInfo(ctx, updateOrderReq)
		if err != nil {
			return err
		}

		result = OrderResult(order)
		return nil
	})

	return &result, err
}
