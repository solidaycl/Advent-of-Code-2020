package main

import (
	"testing"
)

func Test1(t *testing.T) {
	seat := getSeat("BFFFBBFRRR")
	if seat.row != 70 {
		t.Errorf("BFFFBBFRRR row, actual: %d, wanted 70", seat.row)
	}
	if seat.column != 7 {
		t.Errorf("BFFFBBFRRR column, actual: %d, wanted 7", seat.column)
	}
	if seat.row*8+seat.column != 567 {
		t.Errorf("BFFFBBFRRR id, actual: %d, wanted 567", seat.row*8+seat.column)
	}
}

func Test2(t *testing.T) {
	seat := getSeat("FFFBBBFRRR")
	if seat.row != 14 {
		t.Errorf("row, actual: %d, wanted 14", seat.row)
	}
	if seat.column != 7 {
		t.Errorf("column, actual: %d, wanted 7", seat.column)
	}
	if seat.row*8+seat.column != 119 {
		t.Errorf("id, actual: %d, wanted 119", seat.row*8+seat.column)
	}
}

func Test3(t *testing.T) {
	seat := getSeat("BBFFBBFRLL")
	if seat.row != 102 {
		t.Errorf("row, actual: %d, wanted 102", seat.row)
	}
	if seat.column != 4 {
		t.Errorf("column, actual: %d, wanted 4", seat.column)
	}
	if seat.row*8+seat.column != 820 {
		t.Errorf("id, actual: %d, wanted 820", seat.row*8+seat.column)
	}
}

func Test4(t *testing.T) {
	seat := getSeat("FFFFFFFLLL")
	if seat.row != 0 {
		t.Errorf("row, actual: %d, wanted 0", seat.row)
	}
	if seat.column != 0 {
		t.Errorf("column, actual: %d, wanted 0", seat.column)
	}
	if seat.row*8+seat.column != 0 {
		t.Errorf("id, actual: %d, wanted 0", seat.row*8+seat.column)
	}
}

func Test5(t *testing.T) {
	seat := getSeat("BBBBBBBRRR")
	if seat.row != 127 {
		t.Errorf("row, actual: %d, wanted 127", seat.row)
	}
	if seat.column != 7 {
		t.Errorf("column, actual: %d, wanted 7", seat.column)
	}
	if seat.row*8+seat.column != 127*8+7 {
		t.Errorf("id, actual: %d, wanted %d", 127*8+7, seat.row*8+seat.column)
	}
}

//BFFFBBFRRR: row 70, column 7, seat ID 567.
//FFFBBBFRRR: row 14, column 7, seat ID 119.
//BBFFBBFRLL: row 102, column 4, seat ID 820.
