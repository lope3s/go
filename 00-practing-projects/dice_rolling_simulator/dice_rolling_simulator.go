package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var DefaultMessage = "You rolled "

func Message(result int) string {
	return fmt.Sprint(DefaultMessage, result)
}

func RollDice() (result int) {
	// You can create a specific rand that follows it's own source:
	// this will always use this source to give always a random value;
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result = r.Intn(6) + 1
	return result
}

func TimesRoll(times int) (results []int) {
	updateResults := func() {
		result := RollDice()
		results = append(results, result)
	}

	if times == 0 {
		updateResults()
		return results
	}

	for i := 0; i < times; i++ {
		updateResults()
	}

	return results
}

func main() {
	// Here we are initializing the DEFAULT source for random values;
	// Without this, calling a rand.Intn will always give the same result
	// sequence.
	// rand.Seed(time.Now().UnixNano())

	args := os.Args

	if len(args) == 1 {
		results := TimesRoll(0)

		message := Message(results[0])
		fmt.Printf("%s\n", message)
		return
	}

	timesIndex := args[1]

	timesToRoll, err := strconv.Atoi(timesIndex)

	if err != nil {
		fmt.Println("Conversion error:", err)
		return
	}

	results := TimesRoll(timesToRoll)

    // range returns: index, value
	for _, result := range results {
		message := Message(result)
		fmt.Printf("%s\n", message)
	}
}
