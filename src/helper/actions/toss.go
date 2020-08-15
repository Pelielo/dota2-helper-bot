package actions

import "math/rand"

func CoinToss() string {
	coin := []string{
		"heads",
		"tails",
	}

	return coin[rand.Intn(len(coin))]
}
