package password

import (
	"math/rand"
	"strings"
	"time"
)

func SetSeed() {
	rand.Seed(time.Now().UnixNano())
}

func CrackPass(pass string) (bool, string) {
	chars := "abcdefghijklmnopqrstuvwxyz1234567890"
	cracked := false
	length := len(pass)
	var charsList = strings.Split(chars, "")

	rand.Shuffle(len(charsList), func(i, j int) { charsList[i], charsList[j] = charsList[j], charsList[i] })
	testPass := strings.Join(charsList[0:length], "")

	if testPass == pass {
		cracked = true
	}
	return cracked, testPass
}
