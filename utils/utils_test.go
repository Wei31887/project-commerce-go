package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFloatToMoney(t *testing.T) {
	money := 100.10
	moneyString := FloatToMoney(money)
	require.Equal(t, "100.10", moneyString)
}