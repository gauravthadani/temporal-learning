package main

import (
	"fmt"
	"os"

	"temporal101/scratch/greeting"
)

func main() {

	name := os.Args[1]

	greeting, _ := greeting.GreetSomeone(nil, name)
	fmt.Println(greeting)

}
