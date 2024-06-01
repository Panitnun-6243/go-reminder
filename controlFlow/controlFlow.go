package controlflow

import "strconv"

func OddEvenNumber(a int) string {
	if a%2 == 0 {
		return "Number " + strconv.Itoa(a) + " is even number"
	}
	return "Number " + strconv.Itoa(a) + " is odd number"
}
