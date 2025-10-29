package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	filePath     = "/Users/chris/Projects/go/coffee/data.json"
	tempFilePath = "/Users/chris/Projects/go/coffee/tempfile.json"
)

var AcceptedCoffeeTypes = []string{
	"espresso",
	"latte",
	"cappucino",
	"flat White",
	"mocha",
	"americano",
	"filter",
	"caffè crema",
	"instant",
}

var AcceptedCoffeeSize = []string{
	"small",
	"medium",
	"large",
}

var AcceptedVenue = []string{
	"home",
	"coffee shop",
	"office",
}

type CoffeeLog struct {
	ID         int
	Timestamp  string
	CoffeeType string
	Venue      string
	Size       string
}

var seedData []CoffeeLog

// If the read or unmarshalling fails the program should end
func init() {
	bs, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	unmarshalErr := json.Unmarshal(bs, &seedData)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
}

// get the last ID from seedData
func getLastId(seedData []CoffeeLog) (int, error) {
	fmt.Println("Attemping to get the previous coffeeLog ID ")
	lastID := 0
	if len(seedData) == 0 {
		lastID = 0
	}
	for _, CoffeeLog := range seedData {
		if CoffeeLog.ID > lastID {
			lastID = CoffeeLog.ID
		}
	}
	fmt.Println("last ID is: ", lastID)
	return lastID, nil
}

func addCoffee(CoffeeType, Venue, Size string) (CoffeeLog, error) {
	fmt.Println("attempting to add a new log to the coffeeLog")
	lastID, err := getLastId(seedData)
	if err != nil {
		log.Fatal(err)
	}

	newCoffeeLog := CoffeeLog{
		ID:         lastID + 1,
		Timestamp:  time.Now().Format("2006-01-02"),
		CoffeeType: CoffeeType,
		Venue:      Venue,
		Size:       Size,
	}

	fmt.Println("new CoffeeLog: ", newCoffeeLog)

	seedData = append(seedData, newCoffeeLog)

	fmt.Println("new seedData: ", seedData)

	bs, err := json.Marshal(seedData)
	if err != nil {
		log.Fatal(err)
	}
	writeErr := os.WriteFile(tempFilePath, bs, 0644)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
	os.Rename(tempFilePath, filePath)

	return newCoffeeLog, nil
}

func main() {

	cmd := ""

	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	if cmd == "" {
		fmt.Println("help message goes here! ")
	}

	if cmd == "add" {
		// create a new coffee log.
		flags := os.Args[2:]
		fmt.Println("Time to add a coffee ")
		addFlags := flag.NewFlagSet("add", flag.ExitOnError)
		passedCoffeeType := addFlags.String("type", "americano", "The type of coffee (options: Espresso, Latte, Cappuccino, Flat White, Mocha, Americano, Filter, Caffè Crema, Instant). Default: Americano.")
		// passedCoffeeTypeShort := addFlags.String("t", "americano", "The type of coffee (options: Espresso, Latte, Cappuccino, Flat White, Mocha, Americano, Filter, Caffè Crema, Instant). Default: Americano. ( shorthand )")
		passedCoffeeSize := addFlags.String("size", "medium", "The size of the coffee (options: Small, Medium, Large). Default: Medium. Note: Flat White has no size option.")
		// passedCoffeeSizeShort := addFlags.String("s", "medium", "The size of the coffee (options: Small, Medium, Large). Default: Medium. Note: Flat White has no size option. ( Shorthand )")
		passedVenue := addFlags.String("venue", "home", "Where you are having your coffee (options: Home, Coffee Shop, Office). Default: Home.")
		// passedVenueShort := addFlags.String("v", "home", "Where you are having your coffee (options: Home, Coffee Shop, Office). Default: Home. ( shorthand )")

		addFlags.Parse(flags)

		chosenCoffee := strings.ToLower(*passedCoffeeType)
		chosenSize := strings.ToLower(*passedCoffeeSize)
		chosenVenue := strings.ToLower(*passedVenue)

		// if chosenCoffee == "americano" && *passedCoffeeTypeShort != "americano" {
		// 	chosenCoffee = *passedCoffeeTypeShort
		// }
		// if chosenSize == "medium" && *passedCoffeeSizeShort != "medium" {
		// 	chosenSize = *passedCoffeeSizeShort
		// }
		// if chosenVenue == "home" && *passedVenueShort != "home" {
		// 	chosenVenue = *passedVenueShort
		// }

		coffeeIsValid := false
		sizeIsValid := false
		venueIsValid := false

		for _, coffee := range AcceptedCoffeeTypes {
			if chosenCoffee == coffee {
				coffeeIsValid = true
				fmt.Println("recieved an accepted coffee drink")
				break
			}
		}

		for _, size := range AcceptedCoffeeSize {
			if chosenSize == size {
				sizeIsValid = true
				fmt.Println("recieved an accepted coffee size")
				break
			}
		}

		for _, venue := range AcceptedVenue {
			if chosenVenue == venue {
				venueIsValid = true
				fmt.Println("recieved an accepted coffee venue")
				break
			}
		}

		if !coffeeIsValid {
			// Print usage statement
			fmt.Println("type usage message")
		}
		if !sizeIsValid {
			// Print usage statement
			fmt.Println("size usage message")

		}
		if !venueIsValid {
			// Print usage statement
			fmt.Println("venue usage message")

		}

		coffeeLog, err := addCoffee(chosenCoffee, chosenSize, chosenVenue)
		if err != nil {
			log.Fatalf("failed to add your %s to the coffee log. err: %w", chosenCoffee, err)
		}
		fmt.Println("CoffeeLog added: ", coffeeLog)

	}

	if cmd == "today" {
		// list todays coffee consumption stats
		flags := os.Args[2:]
		year, month, day := time.Now().Date()
		fmt.Printf("year: %v, month: %v, day: %v", year, month, day)
		todayFlags := flag.NewFlagSet("today", flag.ExitOnError)
		todayFlags.String("type", "americano", "filter by type (optional)")
		// todayFlags.String("t", "americano", "filter by type (optional)")
		todayFlags.String("venue", "home", "filter by venue (optional)")
		//  todayFlags.String("v", "home", "filter by venue (optional)")
		todayFlags.String("size", "medium", "filter by size (optional)")
		// todayFlags.String("s", "medium", "filter by size (optional)")

		todayFlags.Parse(flags)
	}

	if cmd == "week" {
		// list coffee consumption for this week
		flags := os.Args[2:]
		year, month, day := time.Now().Date()
		fmt.Printf("year: %v, month: %v, day: %v", year, month, day)
		weekFlags := flag.NewFlagSet("today", flag.ExitOnError)
		weekFlags.String("type", "americano", "filter by type (optional)")
		// weekFlags.String("t", "americano", "filter by type (optional)")
		weekFlags.String("venue", "home", "filter by venue (optional)")
		// weekFlags.String("v", "home", "filter by venue (optional)")
		weekFlags.String("size", "medium", "filter by size (optional)")
		// weekFlags.String("s", "medium", "filter by size (optional)")

		weekFlags.Parse(flags)
	}

	if cmd == "month" {
		// list coffee consumption for this month
		flags := os.Args[2:]
		year, month, day := time.Now().Date()
		fmt.Printf("year: %v, month: %v, day: %v", year, month, day)
		monthFlags := flag.NewFlagSet("today", flag.ExitOnError)
		monthFlags.String("type", "americano", "filter by type (optional)")
		// monthFlags.String("t", "americano", "filter by type (optional)")
		monthFlags.String("venue", "home", "filter by venue (optional)")
		// monthFlags.String("v", "home", "filter by venue (optional)")
		monthFlags.String("size", "medium", "filter by size (optional)")
		// monthFlags.String("s", "medium", "filter by size (optional)")

		monthFlags.Parse(flags)
	}

	if cmd == "year" {
		// list coffee consumption for this year
		flags := os.Args[2:]
		year, month, day := time.Now().Date()
		fmt.Printf("year: %v, month: %v, day: %v", year, month, day)
		yearFlags := flag.NewFlagSet("today", flag.ExitOnError)
		yearFlags.String("type", "americano", "filter by type (optional)")
		// yearFlags.String("t", "americano", "filter by type (optional)")
		yearFlags.String("venue", "home", "filter by venue (optional)")
		// yearFlags.String("v", "home", "filter by venue (optional)")
		yearFlags.String("size", "medium", "filter by size (optional)")
		// yearFlags.String("s", "medium", "filter by size (optional)")

		yearFlags.Parse(flags)
	}

	if cmd == "all" {
		// list coffee consumption for this all time.
		flags := os.Args[2:]
		year, month, day := time.Now().Date()
		fmt.Printf("year: %v, month: %v, day: %v", year, month, day)
		allTimeFlags := flag.NewFlagSet("today", flag.ExitOnError)
		allTimeFlags.String("type", "americano", "filter by type (optional)")
		// allTimeFlags.String("t", "americano", "filter by type (optional)")
		allTimeFlags.String("venue", "home", "filter by venue (optional)")
		// allTimeFlags.String("v", "home", "filter by venue (optional)")
		allTimeFlags.String("size", "medium", "filter by size (optional)")
		// allTimeFlags.String("s", "medium", "filter by size (optional)")

		allTimeFlags.Parse(flags)
	}
}
