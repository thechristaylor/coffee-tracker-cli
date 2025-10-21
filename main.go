package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type CoffeeLog struct {
	Id         int
	Timestamp  int64
	CoffeeType string
	Venue      string
	Size       string
}

var file = "data.json"

func init() {
	bs, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("error reading data.json file %w", err.Error())
	}
	seedData := string(bs)
	fmt.Printf("seedData: %v", seedData)
}

func lastId() (int, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("error reading data.json file: %w", err.Error())
	}
	fmt.Printf("data contains: %v", data)
	return 1, nil
}

func (c CoffeeLog) addCoffee() (string, error) {
	Id, err := lastId()
	if err != nil {
		fmt.Println("error trying to get the last coffee id: %w", err.Error())
	}
	fmt.Println("id: ", Id)
	return "", nil
}

func main() {

	cmd := os.Args[1]
	fmt.Println("os.Args[1]: ", cmd)

	if cmd == "" {
		fmt.Println("help message goes here! ")
	}

	if cmd == "add" {
		// create a new coffee log.
		flags := os.Args[2:]
		fmt.Println("Time to add a coffee")
		fmt.Printf("flags contains: %v", flags)
		addFlags := flag.NewFlagSet("add", flag.ExitOnError)
		addFlags.String("type", "americano", "The type of coffee (options: Espresso, Latte, Cappuccino, Flat White, Mocha, Americano, Filter, Caffè Crema, Instant). Default: Americano.")
		addFlags.String("t", "americano", "The type of coffee (options: Espresso, Latte, Cappuccino, Flat White, Mocha, Americano, Filter, Caffè Crema, Instant). Default: Americano. ( shorthand )")
		addFlags.String("size", "medium", "The size of the coffee (options: Small, Medium, Large). Default: Medium. Note: Flat White has no size option.")
		addFlags.String("s", "medium", "The size of the coffee (options: Small, Medium, Large). Default: Medium. Note: Flat White has no size option. ( Shorthand )")
		addFlags.String("venue", "home", "Where you are having your coffee (options: Home, Coffee Shop, Office). Default: Home.")
		addFlags.String("v", "home", "Where you are having your coffee (options: Home, Coffee Shop, Office). Default: Home. ( shorthand )")

		addFlags.Parse(flags)
	}

	if cmd == "stats" {
		// list current status
		flags := os.Args[2:]
		day := time.Now().Day()
		fmt.Println("Lets look at the tally... ")
		fmt.Printf("flags contains:  %v", flags)
		statsFlags := flag.NewFlagSet("stats", flag.ExitOnError)
		statsFlags.Duration("day", time.Duration(day), "The day that you want to get stats for")
		// add a week range
		statsFlags.Duration("week", time.Duration(day), "The week you want to get stats for") // need to fix this flag
		statsFlags.Duration("month", time.Duration(time.Now().Month()), "The month that you want to get stats for")
		// add a custom range flag
		statsFlags.Duration("custom", time.Duration(day), "the custom date range you want stats for")
	}
}
