package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"neophora/cli"
	"os"
	"strings"
)

func main() {
	address := os.ExpandEnv("${DBS_ADDRESS}")
	addresses := strings.Split(address, " ")
	client := &cli.T{
		Addresses: addresses,
		TryTimes:  3,
	}

	key := []byte(os.Args[1])
	var ret []byte

	if err := client.Call("DB.Get", key, &ret); err != nil {
		log.Fatalln(err)
	}
	switch os.ExpandEnv("${ENC}") {
	case "hex":
		fmt.Println(hex.EncodeToString(ret))
	case "str":
		fmt.Println(string(ret))
	default:
		fmt.Println(ret)
	}
}
