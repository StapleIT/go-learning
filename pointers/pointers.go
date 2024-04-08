package pointers

// import (
// 	"fmt"
// )

// If the function needs to mutate a variable in the main function
// a pointer to the variable should be given to the function
// this forces the function to change the value at the address
// of the actual variable rather the address of a copy in the function

func Incr(s *int) {
	*s++
}

type S struct {
	Name string
	Age  int
}

// The same applies for complex types such as structs, pass a pointer
// to the struct as a function attribute.  Then it is possible
// to write to the fields in the struct at the address outside of the fcuntion
// indicated by the pointer.

func ChangeAge(t *S, newAge int) { t.Age = newAge }

//func changeAge(t *S, newAge int){(*t).age = newAge} // dereference t before accessing 'age'

// func main() {
// 	fmt.Println("Hello")
// 	a := 0
// 	Incr(&a) // pointer to the variable as attribute of function
// 	fmt.Println(a)
// 	var s = S{name: "Odie", age: 42}
// 	changeAge(&s, 23) // must explicitly include the pointer to 's' (&s)
// 	fmt.Println(s)    // {"Odie" 23}
// }
