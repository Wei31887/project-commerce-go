package sqlc

import (
	"context"
	"project/e-commerce/db/sqlc"
	"project/e-commerce/utils"
	"testing"

	"github.com/stretchr/testify/require"
)


func CreateRandomCategory(t *testing.T) sqlc.ProductCategory {

	arg := sqlc.CreateProductCategoryParams{
		Name: utils.RandomString(5),
		Sort: int32(utils.RandomInt(0, 6)),
	}

	goodType, err := testQueries.CreateProductCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, goodType)

	require.Equal(t, goodType.Name, arg.Name)
	require.Equal(t, goodType.Sort, arg.Sort)
	require.NotZero(t, goodType.ID)

	return goodType
}

func TestCreateProductCategory(t *testing.T) {
	CreateRandomCategory(t)
}