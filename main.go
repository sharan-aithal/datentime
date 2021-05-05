package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Now:", Now())
}

func Now() time.Time {
	return time.Now()
}
