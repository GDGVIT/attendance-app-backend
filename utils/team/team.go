package team

import (
	"math/rand"
	"time"
)

// generate random invite code for team
const letterBytes = "abcdefghijklmnopqrstuvwxyz"
const digitBytes = "0123456789"

func generateRandomString(length int, characters string) string {
	rand.Seed(time.Now().UnixNano())
	codeBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		codeBytes[i] = characters[rand.Intn(len(characters))]
	}
	return string(codeBytes)
}

func GenerateInviteCode() string {
	code := generateRandomString(3, letterBytes) + "-" + generateRandomString(4, letterBytes) + "-" + generateRandomString(3, digitBytes)
	return code
}
