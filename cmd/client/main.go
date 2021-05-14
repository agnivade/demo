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

	// res, err := client.Register()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(res)
}
