package main

import (
	"fmt"
	"os"

	"temporal101/scratch/greeting"
	"temporal101/scratch/models"
)

func main() {

	name := os.Args[1]

	greeting, _ := greeting.GreetSomeone(nil, models.TranslateRequest{
		Name: name,
	})
	fmt.Println(greeting)

}
