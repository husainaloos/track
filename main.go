package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("empty arguments")
	}

	switch os.Args[1] {
	case "create":
		create(os.Args[2:])
	case "end":
		end()
	default:
		log.Fatalf("unrecognized argument: %s", os.Args[1])
	}
}

func create(args []string) {
	log.Println("creating a task")
}

func end() {
	log.Println("ending a task")
}
