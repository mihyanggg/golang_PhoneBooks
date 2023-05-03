package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

// add data
func add_data() {
	data = append(data, Entry{"Mary", "Doe", "010-1234-5678"})
	data = append(data, Entry{"Jhon", "Black", "010-5678-1234"})
	data = append(data, Entry{"Sara", "Rosa", "010-2435-6240"})
}

func main() {
	// add data
	add_data()

	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <argument>\n", exe)
		return
	}

	// choice.. command
	switch arguments[1] {

	// search command
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)
	// list command
	case "list":
		list()
	// command not found
	default:
		fmt.Println("Not a valid option")
	}
}
