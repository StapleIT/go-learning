package utilities

import (
	"strings"
	"testing"
)

func TestLetter2Number(t *testing.T) {
	type Test struct {
		letter string
		number int
	}

	ltrs := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	for nmbr, ltr := range ltrs {
		tst := Test{ltr, nmbr + 1}
		t.Run(tst.letter, func(t *testing.T) {
			got := letter2Number(tst.letter)
			if got != tst.number {
				t.Errorf("got %v , expected %v", got, tst.number)
			}
		})
	}

}
