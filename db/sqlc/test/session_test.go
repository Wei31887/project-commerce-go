package sqlc

import (
	"context"
	"project/e-commerce/db/sqlc"
	"project/e-commerce/utils"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateRandomSession(t *testing.T) *sqlc.Session {
	user := CreateRandomUser(t)
	uuid, err := uuid.NewRandom()
	require.NoError(t, err)
	arg := sqlc.CreateSessionParams{
		ID: uuid,
		UserID: user.ID,
		UserAgent: utils.RandomString(5),
		Ip: utils.RandomIp(),
		IsBlocked: false,
	}
	session, err := testQueries.CreateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session)
	return &session
}

func TestCreateSession(t *testing.T) {
	CreateRandomSession(t)
}