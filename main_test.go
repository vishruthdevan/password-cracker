package main

import (
	"password-cracker/password"
	"testing"
)

func TestHash(t *testing.T) {
	password.SetSeed()
	count := 1000000
	for i := 0; i < count; i++ {
		var cracked, testPass = password.CrackPass("qwerty")
		if cracked {
			t.Logf("success | number of tries: %v", i+1)
			t.Logf("generated password: %v", testPass)
			break
		} else {
			t.Errorf("failure | could not crack password after %v tries", i+1)
			t.Errorf("generated password: %v", testPass)
		}
	}
}
