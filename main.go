package main

import (
	"fmt"

	p "github.com/StapleIT/go-learning/pointers"
	u "github.com/StapleIT/go-learning/utils"
)

func main() {
	u.UtilFunc()

	fmt.Println("Hello")
	a := 0
	p.Incr(&a) // pointer to the variable as attribute of function
	fmt.Println(a)
	var s = p.S{Name: "Odie", Age: 42}
	p.ChangeAge(&s, 23) // must explicitly include the pointer to 's' (&s)
	fmt.Println(s)      // {"Odie" 23}
}
