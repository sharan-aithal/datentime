package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Now:", Now())
	fmt.Println("Unix:", Unix())
}

func Now() time.Time {
	return time.Now()
}

func Unix() int64 {
	return time.Now().Unix()
}
