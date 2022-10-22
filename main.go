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

	lib.CheckIn(host+checkInUrl, cookie)
	lib.CheckStatus(host+checkInStatusUrl, cookie)
}