package main

import (
	"fmt"
	"strings"

	"github.com/rfyiamcool/smooth_wrr"
)

func main() {
	peers := []*smoothwrr.Peer{
		smoothwrr.CreatePeer("a", 5),
		smoothwrr.CreatePeer("b", 1),
		smoothwrr.CreatePeer("c", 1),
	}

	pool := smoothwrr.NewPool(peers)
	for index := 0; index < 5; index++ {
		fmt.Println("Loop: ", index)

		result := []string{}
		for i := 0; i < 7; i++ {
			peer := pool.Get()
			result = append(result, peer)
		}

		resultStr := strings.Join(result, ",")
		fmt.Println(resultStr)
	}
}
