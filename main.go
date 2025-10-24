package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type CoffeeLog struct {
	ID         int
	Timestamp  string
	CoffeeType string
	Venue      string
	Size       string
}

var file = "data.json"
var seedData []CoffeeLog

// If the read or unmarshalling fails the program should end
func init() {
	bs, err := os.ReadFile("/Users/chris/Projects/go/coffee/data.json")
	if err != nil {
		log.Fatal(err)
	}
	unmarshalErr := json.Unmarshal(bs, &seedData)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
}

// get the len of seedData to find the lastId
func lastId(seedData []CoffeeLog) (int, error) {
	lastId := len(seedData) - 1
	fmt.Printf("lastId is: %v", lastId)
	return lastId, nil
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

	cmd := ""

	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

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

	if cmd == "today" {
		// list todays coffee consumption stats
		flags := os.Args[2:]
		year, month, day := time.Now().Date()
		todayFlags := flag.NewFlagSet("today", flag.ExitOnError)
		todayFlags.String("type", "americano", "filter by type (optional)")
		todayFlags.String("t", "americano", "filter by type (optional)")
		todayFlags.String("venue", "home", "filter by venue (optional)")
		todayFlags.String("v", "home", "filter by venue (optional)")
		todayFlags.String("size", "medium", "filter by size (optional)")
		todayFlags.String("s", "medium", "filter by size (optional)")

		todayFlags.Parse(flags)
	}

	if cmd == "week" {
		// list coffee consumption for this week
		flags := os.Args[2:]
		year, month, day := time.Now().Date()
		weekFlags := flag.NewFlagSet("today", flag.ExitOnError)
		weekFlags.String("type", "americano", "filter by type (optional)")
		weekFlags.String("t", "americano", "filter by type (optional)")
		weekFlags.String("venue", "home", "filter by venue (optional)")
		weekFlags.String("v", "home", "filter by venue (optional)")
		weekFlags.String("size", "medium", "filter by size (optional)")
		weekFlags.String("s", "medium", "filter by size (optional)")

		weekFlags.Parse(flags)
	}

	if cmd == "month" {
		// list coffee consumption for this month
		flags := os.Args[2:]
		year, month, day := time.Now().Date()
		monthFlags := flag.NewFlagSet("today", flag.ExitOnError)
		monthFlags.String("type", "americano", "filter by type (optional)")
		monthFlags.String("t", "americano", "filter by type (optional)")
		monthFlags.String("venue", "home", "filter by venue (optional)")
		monthFlags.String("v", "home", "filter by venue (optional)")
		monthFlags.String("size", "medium", "filter by size (optional)")
		monthFlags.String("s", "medium", "filter by size (optional)")

		monthFlags.Parse(flags)
	}

	if cmd == "year" {
		// list coffee consumption for this year
		flags := os.Args[2:]
		year, month, day := time.Now().Date()
		yearFlags := flag.NewFlagSet("today", flag.ExitOnError)
		yearFlags.String("type", "americano", "filter by type (optional)")
		yearFlags.String("t", "americano", "filter by type (optional)")
		yearFlags.String("venue", "home", "filter by venue (optional)")
		yearFlags.String("v", "home", "filter by venue (optional)")
		yearFlags.String("size", "medium", "filter by size (optional)")
		yearFlags.String("s", "medium", "filter by size (optional)")

		yearFlags.Parse(flags)
	}

	if cmd == "all" {
		// list coffee consumption for this all time.
		flags := os.Args[2:]
		year, month, day := time.Now().Date()
		allTimeFlags := flag.NewFlagSet("today", flag.ExitOnError)
		allTimeFlags.String("type", "americano", "filter by type (optional)")
		allTimeFlags.String("t", "americano", "filter by type (optional)")
		allTimeFlags.String("venue", "home", "filter by venue (optional)")
		allTimeFlags.String("v", "home", "filter by venue (optional)")
		allTimeFlags.String("size", "medium", "filter by size (optional)")
		allTimeFlags.String("s", "medium", "filter by size (optional)")

		allTimeFlags.Parse(flags)
	}
}
