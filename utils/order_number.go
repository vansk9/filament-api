package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateOrderNumber() string {
	timestamp := time.Now().Unix()
	random := rand.Intn(1000)
	return fmt.Sprintf("ORD-%d-%03d", timestamp, random)
}
