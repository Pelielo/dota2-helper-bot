package actions

import (
	"math/rand"
	"strconv"
)

func RollNumber() string {
	return strconv.Itoa(rand.Intn(101))
}
