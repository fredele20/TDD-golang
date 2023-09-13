package fizzbuzz

import "strconv"


func Run(nums []int) []string {
	strings := []string{}
	for _, v := range nums {
		if v % 3 == 0 {
			strings = append(strings, "Fizz")
		} else if v % 5 == 0 {
			strings = append(strings, "Buzz")
		} else {
			strings = append(strings, strconv.Itoa(v))
		}
	}
	return strings
}