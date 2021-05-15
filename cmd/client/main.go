package main

import (
	"fmt"
	"os"
	"time"

	"github.com/agnivade/demo"
)

func main() {
	client, err := demo.NewClient(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Register()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 2; i++ {
		ch, err := client.GetChannel("z9cz74g3ninjdyxdi8ryt5qz7c")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", ch)

		time.Sleep(500 * time.Millisecond)
	}

	client.Callback("hello world")

	time.Sleep(time.Second)
}

// type Server struct {.. s *Suite ..}
// NewServer(network, addr string) *Server
//
// (*Suite) GetChannel(..)
// (*Suite) GetTeam(..)
// (*Suite) SaveConfig(..)

// NewClient(network, addr string, hooks Interface) *Client
// (*Client) GetChannel(..)
// (*Client) GetTeam(..)
// (*Client) SaveConfig(..)
// (*Client) ...

// type Hooks interface {
// OnConfigChange()
// OnClusterEvent()
// ...
// }