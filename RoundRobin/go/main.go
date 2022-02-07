package main

import (
	"net/url"

	roundrobin "github.com/hlts2/round-robin"
)

var i = 0
var servers = []string{"127.0.0.1:8000", "127.0.0.1:8001", "127.0.0.1:8003"}

// Balance returns one of the servers based using round-robin algorithm
func Balance() string {
	server := servers[i]
	i++

	// it means that we reached the end of servers
	// and we need to reset the counter and start
	// from the beginning
	if i >= len(servers) {
		i = 0
	}
	return server
}

func main() {
	// requests loop
	// for i := 0; i < 20; i++ {
	// 	fmt.Println(Balance())
	// }

	rr, _ := roundrobin.New(
		&url.URL{Host: "192.168.33.10"},
		&url.URL{Host: "192.168.33.11"},
		&url.URL{Host: "192.168.33.12"},
		&url.URL{Host: "192.168.33.13"},
	)

	rr.Next() // {Host: "192.168.33.10"}
	rr.Next() // {Host: "192.168.33.11"}
	rr.Next() // {Host: "192.168.33.12"}
	rr.Next() // {Host: "192.168.33.13"}
	rr.Next() // {Host: "192.168.33.10"}
	rr.Next() // {Host: "192.168.33.11"}
}
