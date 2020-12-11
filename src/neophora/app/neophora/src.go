package main

import (
	"log"
	"neophora/biz/api"
	"neophora/biz/data"
	"neophora/lib/joh"
	"net/http"
	"net/rpc"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func main() {
	address := os.ExpandEnv("0.0.0.0:${NEOPHORA_PORT}")
	log.Println("[LISTEN]", address)
	http.ListenAndServe(address, &joh.T{})
}

func init() {
	netowrk := os.ExpandEnv("${REDIS_NETWORK}")
	address := os.ExpandEnv("${REDIS_ADDRESS}")
	maxidle, err := strconv.Atoi(os.ExpandEnv("${REDIS_MAXIDLE}"))
	if err != nil {
		log.Fatalln(err)
	}
	db := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial(netowrk, address)
	}, maxidle)
	rpc.Register(&api.T{
		Data: &data.T{
			DB: db,
		},
	})
}
