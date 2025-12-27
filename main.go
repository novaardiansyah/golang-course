package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dbData = []string{"id-1", "id-2", "id-3", "id-4", "id-5"}

func main() {
	t0 := time.Now()
	
	for i := 0; i < len(dbData); i++ {
		dbCall(i)
	}

	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("\nThe result from the database is: %v", dbData[i])
}
