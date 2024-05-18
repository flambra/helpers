package hLog

import (
	"log"
	"time"
)

// Performance calculates and prints the execution time since the given start time.
// It can be used with defer to measure the performance of a function.
//
//	func Function() {
//		defer Performance(time.Now(), "Function")
//		// Code to measure performance
//	}
// 
// Function took: 2.63161ms
func Performance(start time.Time, function string) {
	elapsed := time.Since(start)
	log.Printf("%s took: %s\n", function, elapsed)
}
