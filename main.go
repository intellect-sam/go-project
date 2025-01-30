package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
	medicalForm()
}

func medicalForm() {

	// short variable declaration

	peanutAllergy, firstName, startUpTime := getConfig()

	fmt.Println(firstName + "\n" + strconv.FormatBool(peanutAllergy) + "\n" + startUpTime.Format(time.RFC3339))
}

func getConfig() (bool, string, time.Time) {
	return true, "John", time.Now()
}
