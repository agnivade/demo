package main

import (
	"fmt"
	"os"

	"github.com/agnivade/demo"
)

func main() {
	srv, err := demo.NewServer(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	go srv.Start()

	select {}
}
