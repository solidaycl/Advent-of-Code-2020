package main

import "strconv"

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func part2() string {

	var seatIds []int
	for _, line := range getInput() {
		seatIds = append(seatIds, getSeat(line).id)
	}

	for i := 0; i <= 1000; i++ {

		if !contains(seatIds, i) && contains(seatIds, i-1) && contains(seatIds, i+1) {
			return strconv.Itoa(i)
		}
	}
	return "not found"
}
