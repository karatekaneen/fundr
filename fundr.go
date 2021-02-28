package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	x := os.Getenv("TEST")
	fmt.Println(x)
}
