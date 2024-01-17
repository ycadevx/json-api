package main

import (
	"flag"
	"fmt"
	"log"
)

func addAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account created === ", acc.Number)

	return acc
}

// example data
func addAccounts(s Storage) {
	addAccount(s, "yca", "cakir", "ycadevx94")
}

func main() {
	// TO DO YCA
	add := flag.Bool("add", false, "add the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *add {
		fmt.Println("Adding sample data to database")
		addAccounts(store)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
