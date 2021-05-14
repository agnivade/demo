package main

import (
	"fmt"
	// "time"

	"github.com/agnivade/demo"
)

func main() {
	srv, err := demo.NewServer()
	if err != nil {
		fmt.Println(err)
		return
	}

	go srv.Start()

	// time.Sleep(20 * time.Second)

	// srv.Big()
	select {}
}
