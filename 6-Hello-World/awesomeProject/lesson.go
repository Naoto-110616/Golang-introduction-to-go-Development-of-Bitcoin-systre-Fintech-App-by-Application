package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println("Hello, World!", time.Now())
	fmt.Print(user.Current())
}
