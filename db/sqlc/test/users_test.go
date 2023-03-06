package sqlc

import (
	"context"
	"project/e-commerce/db/sqlc"
	"project/e-commerce/utils"
	"testing"

	"github.com/stretchr/testify/require"
)


func CreateRandomUser(t *testing.T) (sqlc.User) {
	hashpd, err := utils.HashedPassword(utils.RandomString(6))
	require.NoError(t, err)

	paras := sqlc.CreateUserParams{
		UserName: utils.RandomOwner(),
		HashedPassword: hashpd,
		Email: utils.RandomEmail(),
		FullName: utils.RandomOwner(),
	}
	
	user, err := testQueries.CreateUser(context.Background(), paras)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t,user.UserName, user.UserName)
	require.Equal(t,user.Email, user.Email)
	require.Equal(t,user.HashedPassword, user.HashedPassword)
	require.Equal(t,user.FullName, user.FullName)
	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}