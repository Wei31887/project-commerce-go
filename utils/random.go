package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandomFloat(min, max int64) float64 {
	return float64(RandomInt(min, max)) + rand.Float64()
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(k)])
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(5))
}

func RandomIp() string {
	ip := fmt.Sprintf(
		"%s.%s.%s.%s", 
		strconv.Itoa(rand.Intn(256)),
		strconv.Itoa(rand.Intn(256)),
		strconv.Itoa(rand.Intn(256)),
		strconv.Itoa(rand.Intn(256)),
	)
	return ip
}