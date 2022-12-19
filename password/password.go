package password

import (
	"math/rand"
	"strings"
	"time"
)

func CrackPass(pass string) (bool, int) {
	chars := "abcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*()"
	cracked := false
	tries := 0
	length := len(pass)
	var charsList = strings.Split(chars, "")

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(charsList), func(i, j int) { charsList[i], charsList[j] = charsList[j], charsList[i] })
	testPass := strings.Join(charsList[0:length], "")
	print(testPass)
	if testPass == pass {
		cracked = true
	}
	tries++

	return cracked, tries
}
