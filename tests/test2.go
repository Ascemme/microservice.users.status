package main

import (
	"fmt"
	"time"
)

func main() {
	sa := make(chan string)
	go but(sa)
	go sag(sa)

	time.Sleep(time.Second * 10)

}

func but(sa chan string) {
	var name string

	for {
		fmt.Scanf("%s\n", &name)
		sa <- name
	}

}
func sag(sa chan string) {
	for s := range sa {
		fmt.Println(s)
	}

}
