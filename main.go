package main

import (
	"fmt"
	"os"
)

func main() {
	client, err := NewClient(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	ch, err := client.GetChannel("z9cz74g3ninjdyxdi8ryt5qz7c")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", ch)
}

