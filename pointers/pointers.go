package pointers

// If the function needs to mutate a variable in the main function
// a pointer to the variable should be given to the function
// this forces the function to change the value at the address
// of the actual variable rather the address of a copy in the function

func Incr(s *int) {
	*s++
}

// The same applies for complex types such as structs, pass a pointer
// to the struct as a function attribute.  Then it is possible
// to write to the fields in the struct at the address outside of the fcuntion
// indicated by the pointer.

type S struct {
	Name string
	Age  int
}

func ChangeAge(t *S, newAge int) { t.Age = newAge }

// Another application of a pointer is in functions acting as methods on variables of a particular type...

type T int

func (t *T) IncrMthd() {
	*t++
}

// 't' is the 'receiver' of the function.  i.e. it is an instance of the type that the method is designed to act upon.
