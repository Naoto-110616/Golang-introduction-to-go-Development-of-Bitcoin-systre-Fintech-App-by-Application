package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "aksdjfl;jasdf"
}

func main() {
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("mac!")
	case "windows":
		fmt.Println("windows!")
	default:
		fmt.Println("default!", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour() < 17:
		fmt.Println("good afternoon")
	case t.Hour() < 20:
		fmt.Println("good evening")
	case t.Hour() < 24:
		fmt.Println("good night")
	}
}
