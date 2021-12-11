package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatalln("You should enter the sha1 certificate: XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX:XX")
	}
	sha1cert := os.Args[1]

	reg := regexp.MustCompile(`[^0-9A-Fa-f]+`)

	sha1HexStr := reg.ReplaceAllString(sha1cert, "")

	sha1HexBytes, err := hex.DecodeString(sha1HexStr)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(sha1HexBytes))
}
