package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
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
	"flatWhite",
	"mocha",
	"americano",
	"filter",
	"caffèCrema",
	"instant",
}

var AcceptedCoffeeSize = []string{
	"small",
	"medium",
	"large",
}

var AcceptedVenue = []string{
	"home",
	"coffeeShop",
	"office",
}

type CoffeeLog struct {
	ID         int
	Date       string
	CoffeeType string
	Venue      string
	Size       string
}

var seedData []CoffeeLog
var coffeeLogs []CoffeeLog

var espressoCount,
	latteCount,
	cappuccinoCount,
	flatWhiteCount,
	mochaCount,
	americanoCount,
	filterCount,
	caffèCremaCount,
	instantCount int

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
	lastID := 0
	if len(seedData) == 0 {
		lastID = 0
	}
	for _, CoffeeLog := range seedData {
		if CoffeeLog.ID > lastID {
			lastID = CoffeeLog.ID
		}
	}
	return lastID, nil
}

func addCoffee(CoffeeType, Venue, Size string) (CoffeeLog, error) {
	lastID, err := getLastId(seedData)
	if err != nil {
		log.Fatal(err)
	}

	newCoffeeLog := CoffeeLog{
		ID:         lastID + 1,
		Date:       time.Now().Format("2006-01-02"),
		CoffeeType: CoffeeType,
		Venue:      Venue,
		Size:       Size,
	}

	seedData = append(seedData, newCoffeeLog)

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
		fmt.Println("You're having another coffee?? ")
		addFlags := flag.NewFlagSet("add", flag.ExitOnError)
		passedCoffeeType := addFlags.String("type", "americano", "The type of coffee (options: Espresso, Latte, Cappuccino, Flat White, Mocha, Americano, Filter, Caffè Crema, Instant). Default: Americano.")
		passedCoffeeSize := addFlags.String("size", "medium", "The size of the coffee (options: Small, Medium, Large). Default: Medium. Note: Flat White has no size option.")
		passedVenue := addFlags.String("venue", "home", "Where you are having your coffee (options: Home, Coffee Shop, Office). Default: Home.")

		addFlags.Parse(flags)

		chosenCoffee := strings.ToLower(*passedCoffeeType)
		chosenSize := strings.ToLower(*passedCoffeeSize)
		chosenVenue := strings.ToLower(*passedVenue)

		coffeeIsValid := false
		sizeIsValid := false
		venueIsValid := false

		if slices.Contains(AcceptedCoffeeTypes, chosenCoffee) {
			coffeeIsValid = true
		}

		if slices.Contains(AcceptedCoffeeSize, chosenSize) {
			sizeIsValid = true
		}

		if slices.Contains(AcceptedVenue, chosenVenue) {
			venueIsValid = true
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

		_, err := addCoffee(chosenCoffee, chosenSize, chosenVenue)
		if err != nil {
			log.Fatalf("failed to add your %s to the coffee log. err: %s", chosenCoffee, err)
		}
		fmt.Printf("%v added... Enjoy ☕\n", chosenCoffee)
	}

	if cmd == "today" {
		// list todays coffee consumption stats
		flags := os.Args[2:]
		today := time.Now().Format("2006-01-02")
		fmt.Printf("todays date is: %s\n ", today)
		todayFlags := flag.NewFlagSet("today", flag.ExitOnError)
		todayFlags.String("type", "americano", "filter by type (optional)")
		todayFlags.String("venue", "home", "filter by venue (optional)")
		todayFlags.String("size", "medium", "filter by size (optional)")

		todayFlags.Parse(flags)

		for _, coffeeLog := range seedData {
			if coffeeLog.Date == today {
				coffeeLogs = append(coffeeLogs, coffeeLog)
			}
		}
		fmt.Println("Todays coffeeLogs: ", coffeeLogs)

		for _, coffeeLog := range coffeeLogs {
			switch coffeeLog.CoffeeType {
			case "espresso":
				espressoCount++
			case "latte":
				latteCount++
			case "cappuccino":
				cappuccinoCount++
			case "flatWhite":
				flatWhiteCount++
			case "mocha":
				mochaCount++
			case "americano":
				americanoCount++
			case "filter":
				filterCount++
			case "caffèCrema":
				caffèCremaCount++
			case "instant":
				instantCount++
			}
		}

		if espressoCount > 0 {
			fmt.Printf("youve had %v espresso's today", espressoCount)
		}
		if latteCount > 0 {
			fmt.Printf("you've had %v latte's today", latteCount)
		}
		if cappuccinoCount > 0 {
			fmt.Printf("you've had %v cappuccino's today", cappuccinoCount)
		}
		if flatWhiteCount > 0 {
			fmt.Printf("you've had %v flat whites today", flatWhiteCount)
		}
		if mochaCount > 0 {
			fmt.Printf("you've had %v mocha's today", mochaCount)
		}
		if americanoCount > 0 {
			fmt.Printf("you've had %v americans today", americanoCount)
		}
		if filterCount > 0 {
			fmt.Printf("you've had %v filter coffee's today", filterCount)
		}
		if caffèCremaCount > 0 {
			fmt.Printf("you've had %v caffè Crema's today", caffèCremaCount)
		}
		if instantCount > 0 {
			fmt.Printf("you've had %v instant coffee's today", instantCount)
		}
	}
}
