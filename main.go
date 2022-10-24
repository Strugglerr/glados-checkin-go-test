package main

import (
	"checkin/lib"
	"os"
)

func main() {
	host := "https://glados.rocks"

	checkInUrl := "/api/user/checkin"
	checkInStatusUrl := "/api/user/status"

	cookie := os.Getenv("COOKIE")
	// use env or git action secrets

	msg1 := lib.CheckIn(host+checkInUrl, cookie)
	msg2 := lib.CheckStatus(host+checkInStatusUrl, cookie)
	msg := msg1 + "\n" + msg2
	lib.Push(msg)
	// push message by pushplus(https://www.pushplus.plus/)
}
