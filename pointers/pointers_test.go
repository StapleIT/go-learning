package pointers

import "testing"

func TestIncr(t *testing.T) {
	a := 3
	Incr(&a)
	got := a
	if got != 4 {
		t.Errorf("got %v, expected 4", got)
	}
}

func TestChangeAge(t *testing.T) {
	structVar := S{Name: "Name", Age: 30}
	ChangeAge(&structVar, 25)
	got := structVar.Age
	if got != 25 {
		t.Errorf("got %v , expected 25", got)
	}
}
