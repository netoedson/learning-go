package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))

	number := random.Int()

	if isEvenNumber(number) {
		fmt.Printf("%v is even\n", number)
	} else {
		fmt.Printf("%v is odd\n", number)
	}

	switch {
	case number%5 == 0:
		fmt.Println(number, "is multiple of 5.")
		fallthrough
	case number%3 == 0:
		fmt.Println(number, "is multiple of 3.")
	}

}

func isEvenNumber(number int) bool {
	return number%2 == 0
}
