package main

import (
	"password-cracker/password"
	"testing"
)

func TestHash_superman(t *testing.T) {
	var cracked, tries = password.CrackPass("password")
	if cracked {
		t.Logf("success | number of tries: %v", tries)
	} else {
		t.Errorf("failure | could not crack password")
	}
}
