package main

import "strconv"

type seat struct {
	row    int
	column int
	id     int
}

type _range struct {
	min int
	max int
}

func getSeat(s string) seat {

	ver := _range{
		min: 0,
		max: 127,
	}

	hor := _range{
		min: 0,
		max: 7,
	}

	for _, c := range s {
		switch c {
		case 'F':
			ver.max = ver.max - ((ver.max + 1 - ver.min) / 2)
		case 'B':
			ver.min = ver.min + ((ver.max + 1 - ver.min) / 2)
		case 'L':
			hor.max = hor.max - ((hor.max + 1 - hor.min) / 2)
		case 'R':
			hor.min = hor.min + ((hor.max + 1 - hor.min) / 2)
		}
	}

	return seat{
		row:    ver.max,
		column: hor.max,
		id:     ver.max*8 + hor.max,
	}
}

func part1() string {
	max := 0
	for _, line := range getInput() {
		seat := getSeat(line)
		if seat.id > max {
			max = seat.id
		}
	}
	return strconv.Itoa(max)
}
